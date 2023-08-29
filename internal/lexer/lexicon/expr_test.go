package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("expressions", func() {
	DescribeTable("lex expressions", func(input string, expErr OmegaMatcher, expOut ...lexer.Item) {
		out, err := lexer.New(input, lexExpressionAndThen(nil)).Lex()
		Expect(err).To(expErr)
		if err == nil {
			Expect(out).To(lexer.TokenValuesToBeEqual(expOut))
		}
	},
		Entry("1", `"a"`, BeNil(), T(token.STRING, `a`)),
	)
})
