package lexer_test

import (
	"fmt"
	"testing"

	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/lexerr"
	"github.com/advanderveer/gor/internal/lexer/token"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexer(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer")
}

var _ = Describe("lexing", func() {
	It("should emit and return errors", func() {
		out, err := lexer.New(`foobar `, func(lc lexer.Control) lexer.State {
			return lc.Unexpected('x', lexerr.ExpectedComment)
		}).Lex()
		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:ILLEGAL(unexpected input 'x', expected: comment)]`))
		Expect(err).To(MatchError(MatchRegexp(`unexpected input`)))
	})

	It("should emit regular token", func() {
		out, err := lexer.New(`foobar `, func(lc lexer.Control) lexer.State {
			lc.Emit(token.IDENT)

			return nil
		}).Lex()

		Expect(err).ToNot(HaveOccurred())
		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:IDENT()]`))
	})
})
