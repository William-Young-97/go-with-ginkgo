package supertest_test

import (
	"github/William-Young-97.com/simple-test/supertest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Person.IsChild()", func() {
	Context("when person is a child", func() {
		It("returns true", func() {
			person := supertest.Person{Age: 10}
			response := person.IsChild()
			Expect(response).To(BeTrue())
		})
	})
})
