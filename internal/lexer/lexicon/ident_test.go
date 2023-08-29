package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("identifiers", func() {
	DescribeTable("lex identifiers", func(input string, expErr OmegaMatcher, expOut ...lexer.Item) {
		out, err := lexer.New(input, lexIdentifierAndThen(nil)).Lex()
		Expect(err).To(expErr)
		if err == nil {
			Expect(out).To(lexer.TokenValuesToBeEqual(expOut))
		}
	},
		Entry("1", `a`, BeNil(), T(token.IDENT, `a`)),
		Entry("2", `_x9`, BeNil(), T(token.IDENT, `_x9`)),
		Entry("3", `åβ`, BeNil(), T(token.IDENT, `åβ`)),
		Entry("4", `1abc`, MatchError(MatchRegexp(`expected: letter`)),
			T(token.IDENT, `åβ`)),
	)

	DescribeTable("identifier lists", func(input string, expErr OmegaMatcher, expOut ...lexer.Item) {
		out, err := lexer.New(input, lexIdentListAndThen(nil, isEOF)).Lex()
		Expect(err).To(expErr)
		if err == nil {
			Expect(out).To(lexer.TokenValuesToBeEqual(expOut))
		}
	},
		Entry("1", `a, b,`+"\n"+`_`, BeNil(),
			T(token.IDENT, `a`),
			T(token.COMMA, `,`),
			T(token.IDENT, `b`),
			T(token.COMMA, `,`),
			T(token.IDENT, `_`),
		),
	)
})
