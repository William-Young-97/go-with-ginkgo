package testing_test

import (
	testing2 "github/William-Young-97.com/simple-test/testing"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTesting(t *testing.T) { // This is the entry point for Ginkgo
	RegisterFailHandler(Fail)    // The single line of glue code connecting Ginkgo to Gomega
	RunSpecs(t, "Testing Suite") // Tells Ginkgo to start the test suite
}

var _ = Describe("Person.IsChild()", func() {
	Context("when person is a child", func() {
		It("returns true", func() {
			person := testing2.Person{Age: 10}
			response := person.IsChild()
			Expect(response).To(BeTrue())
		})
	})
})
