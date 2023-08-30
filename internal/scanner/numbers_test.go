package scanner

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("numbers", func() {
	It("should format number kind", func() {
		Expect(numberKind('x')).To(Equal(`hexadecimal literal`))
		Expect(numberKind('o')).To(Equal(`octal literal`))
		Expect(numberKind('b')).To(Equal(`binary literal`))
		Expect(numberKind('g')).To(Equal(`decimal literal`))
	})
})
