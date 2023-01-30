package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEntry(t *testing.T) {
	var test1 = Describe("CreatePhoneNumber()", func() {
		It("basic test", func() {
			Expect(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})).To(Equal("(123) 456-7890"))
		})
	})

	t.Log(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	t.Logf("%T %v", test1, test1)
}
