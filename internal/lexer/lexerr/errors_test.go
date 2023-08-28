package lexerr_test

import (
	"testing"

	"github.com/advanderveer/gor/internal/lexer/lexerr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexerr(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer/lexerr")
}

var _ = Describe("unexpected", func() {
	It("should format 0 expectations", func() {
		Expect((lexerr.UnexpectedError{}).Error()).To(Equal("unexpected input '\x00'"))
	})

	It("shoul format 1 expectations", func() {
		Expect(lexerr.Unexpected('a', lexerr.ExpectedWhiteSpace)).To(
			MatchError(`unexpected input 'a', expected: white space`))
	})

	It("should format 2 expectations", func() {
		Expect(lexerr.Unexpected('b',
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedPackageKeyword,
		)).To(MatchError(`unexpected input 'b', expected: white space or 'package' keyword`))
	})

	It("should format 3 expectations", func() {
		Expect(lexerr.Unexpected('c',
			lexerr.ExpectedWhiteSpace,
			lexerr.ExpectedImportKeyword,
			lexerr.ExpectedPackageKeyword,
		)).To(MatchError(`unexpected input 'c', expected: white space, 'import' keyword or 'package' keyword`))
	})
})
