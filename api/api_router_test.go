package api_test

import (
	"github.com/noahfriedman-ca/server/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("the API router", func() {
	var (
		ms          *httptest.Server
		getResponse = func(uri string) string {
			if r, e := http.Get(ms.URL + "/api/list-available"); e != nil {
				GinkgoT().Error(e)
			} else {
				if b, e := ioutil.ReadAll(r.Body); e != nil {
					GinkgoT().Error(e)
				} else {
					return string(b)
				}
			}

			return ""
		}
	)
	BeforeEach(func() {
		ms = httptest.NewServer(api.Router())
	})

	It("should successfully call the API function when it exists", func() {
		r := getResponse("/api/list-available")
		Expect(r).NotTo(ContainSubstring("404 page not found"))
		Expect(r).NotTo(ContainSubstring("API function does not exist"))
	})
	It("should fail to call the API function when it does not exist", func() {
		r := getResponse("/api/nonexistent-function")
		Expect(r).NotTo(ContainSubstring("404 page not found"))
		Expect(r).To(ContainSubstring("API function does not exist"))
	})

	AfterEach(func() {
		ms.Close()
	})
})
