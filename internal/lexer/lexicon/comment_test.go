package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("comments", func() {
	DescribeTable("lex comments", func(input string, expErr OmegaMatcher, expOut ...lexer.Item) {
		out, err := lexer.New(input, lexCommentAndThen(nil)).Lex()
		Expect(err).To(expErr)
		if err == nil {
			Expect(out).To(TokenValuesToBeEqual(expOut))
		}
	},
		Entry(`1`, `// foo`, BeNil(), T(token.COMMENT, " foo")),
		Entry(`2`, `a`, MatchError(MatchRegexp(`first comment character`))),
		Entry(`3`, `/a`, MatchError(MatchRegexp(`second comment character`))),
	)
})
