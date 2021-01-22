package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
)

var _ = Describe("the router",
	func() {
		DescribeTable("the different routes", func(uri string, localPath string) {
			if _, e := os.Stat(localPath); os.IsNotExist(e) {
				By("creating temporary file '" + localPath + "'")
				if e := os.MkdirAll(path.Dir(localPath), os.ModePerm); e != nil {
					GinkgoT().Error(e)
				} else {
					if f, e := os.Create(localPath); e != nil {
						GinkgoT().Error(e)
					} else {
						if _, e := f.Write([]byte("<html></html>")); e != nil {
							GinkgoT().Error()
						}
						f.Close()
					}
				}

				defer func() {
					if e := os.Remove(localPath); e != nil {
						GinkgoT().Error(e)
					}
					for p := "./" + strings.TrimRight(path.Dir(localPath), "/"); p != "./"; p = "./" + strings.TrimRight(path.Dir(p), "/") {
						if d, e := os.Open(p + "/"); e != nil {
							GinkgoT().Error(e)
						} else {
							if n, e := d.Readdirnames(0); e != nil {
								GinkgoT().Error(e)
							} else {
								if len(n) != 0 {
									break
								} else {
									if e := os.Remove(p); e != nil {
										GinkgoT().Error(e)
									}
								}
							}
						}
					}
				}()
			}

			By("opening '" + localPath + "'")
			f, e := os.Open(localPath)
			if e != nil {
				GinkgoT().Error(e)
			}
			var local string
			if b, e := ioutil.ReadAll(f); e != nil {
				GinkgoT().Error(e)
			} else {
				local = string(b)
			}
			f.Close()

			By("creating a mock server")
			ms := httptest.NewServer(Router())
			defer ms.Close()

			By("requesting the URI '" + uri + "' from the mock server")
			var result string
			if r, e := http.Get(ms.URL + uri); e != nil {
				GinkgoT().Error(e)
			} else {
				if b, e := ioutil.ReadAll(r.Body); e != nil {
					GinkgoT().Error(e)
				} else {
					result = string(b)
				}

			}

			Expect(result).To(Equal(local))
		},
			Entry("the '/' route", "/", "./static/build/index.html"),
			Entry("the 'LICENSE' route", "/LICENSE", "./LICENSE"),
			Entry("the 'sitemap.xml' route", "/sitemap.xml", "./sitemap.xml"),
			Entry("the '/projects/test/' route", "/projects/test/", "./projects/test/build/index.html"),
		)
	})
