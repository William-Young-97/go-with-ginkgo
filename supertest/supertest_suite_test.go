package supertest_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSupertest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Supertest Suite")
}
