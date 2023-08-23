package lexer

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lexing", func() {
	It("should emit error state", func() {
		out := New(`foobar `, func(lc Control) StateFunc { return lc.Errorf("foo") }).Lex()
		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:ILLEGAL(foo)]`))
	})

	It("should emit regular token", func() {
		out := New(`foobar `, func(lc Control) StateFunc {
			lc.Emit(IDENT)

			return nil
		}).Lex()

		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:IDENT()]`))
	})
})
