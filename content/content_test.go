package confirmation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	confirmation "github/William-Young-97.com/simple-test/content"
)

var _ = Describe("Person.IsChild()", func() {
	Context("when person is a child", func() {
		It("returns true", func() {
			person := confirmation.Person{Age: 10}
			response := person.IsChild()
			Expect(response).To(BeTrue())
		})
	})
})
