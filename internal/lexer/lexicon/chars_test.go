package lexicon

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

	It("whitespace", func() {
		Expect(isWhiteSpace(' ')).To(BeTrue())
		Expect(isWhiteSpace('\t')).To(BeTrue())
		Expect(isWhiteSpace('\r')).To(BeTrue())
		Expect(isWhiteSpace('\n')).To(BeTrue())
		Expect(isWhiteSpace('a')).To(BeFalse())
	})

	It("comment", func() {
		Expect(isCommentCharacter('/')).To(BeTrue())
		Expect(isCommentCharacter('\\')).To(BeFalse())
	})

	It("eof", func() {
		Expect(isEOF(-1)).To(BeTrue())
		Expect(isEOF('0')).To(BeFalse())
	})
})
