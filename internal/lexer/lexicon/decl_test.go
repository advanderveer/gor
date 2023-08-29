package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("declarations", func() {
	DescribeTable("var decl", func(input string, expErr OmegaMatcher, expOut ...lexer.Item) {
		out, err := lexer.New(input, lexDecls).Lex()
		Expect(err).To(expErr)
		if err == nil {
			Expect(out).To(lexer.TokenValuesToBeEqual(expOut))
		}
	},
		Entry(`1`, `var a,b string = "bar", "dar"`, BeNil(),
			T(token.VAR, "var"),
			T(token.IDENT, "a"),
			T(token.COMMA, ","),
			T(token.IDENT, "b"),
			T(token.IDENT, "string"),
			T(token.ASSIGN, "="),
			T(token.STRING, "bar"),
			T(token.COMMA, ","),
			T(token.STRING, "dar"),
		),
	)
})
