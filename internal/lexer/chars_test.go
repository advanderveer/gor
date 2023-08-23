package lexer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("char assertions", func() {
	Describe("unicode", func() {
		It("newline", func() {
			Expect(isNewline('x')).To(BeFalse())
			Expect(isNewline('\n')).To(BeTrue())
		})

		It("char", func() {
			Expect(isUnicodeChar('\n')).To(BeFalse())
			Expect(isUnicodeChar(' ')).To(BeTrue())
			Expect(isUnicodeChar(' ')).To(BeTrue())
		})

		It("letter", func() {
			Expect(isUnicodeLetter('a')).To(BeTrue())
			Expect(isUnicodeLetter('A')).To(BeTrue())
			Expect(isUnicodeLetter('â')).To(BeTrue())
			Expect(isUnicodeLetter('0')).To(BeFalse())
		})

		It("digit", func() {
			Expect(isUnicodeDigit('a')).To(BeFalse())
			Expect(isUnicodeDigit('0')).To(BeTrue())
		})
	})

	It("letters", func() {
		Expect(isLetter('a')).To(BeTrue())
		Expect(isLetter('_')).To(BeTrue())
		Expect(isLetter('0')).To(BeFalse())
	})

	It("decimal digit", func() {
		Expect(isDecimalDigit('a')).To(BeFalse())
		Expect(isDecimalDigit('0')).To(BeTrue())
		Expect(isDecimalDigit('9')).To(BeTrue())
	})
})
