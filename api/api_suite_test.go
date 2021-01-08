package api_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = BeforeSuite(func() {
	if d, e := os.Getwd(); e != nil {
		GinkgoT().Error(e)
	} else {
		if _, f := filepath.Split(strings.TrimSuffix(d, "/")); f == "api" {
			if e := os.Chdir(".."); e != nil {
				GinkgoT().Error(e)
			}
		}
	}
})
