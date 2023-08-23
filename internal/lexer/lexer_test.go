package lexer_test

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/advanderveer/gor/internal/lexer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexer(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer")
}

var _ = Describe("lexing", func() {
	It("should emit error state", func() {
		out := lexer.New(`foobar `, func(lc lexer.Control) lexer.State { return lc.Errorf("foo") }).Lex()
		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:ILLEGAL(foo)]`))
	})

	It("should emit regular token", func() {
		out := lexer.New(`foobar `, func(lc lexer.Control) lexer.State {
			lc.Emit(token.IDENT)

			return nil
		}).Lex()

		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:IDENT()]`))
	})
})
