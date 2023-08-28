package lexicon

import (
	"fmt"

	"github.com/advanderveer/gor/internal/lexer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("comments", func() {
	DescribeTable("lex comments", func(input string, expErr OmegaMatcher, expOutput string) {
		out, err := lexer.New(input, lexCommentAndThen(nil)).Lex()
		Expect(err).To(expErr)
		Expect(fmt.Sprint(out)).To(Equal(expOutput))
	},
		Entry(`1`, `// foo`, BeNil(), `[2:0.2:COMMENT( foo)]`),
		Entry(`2`, `a`, MatchError(MatchRegexp(`first comment character`)),
			`[1:0.1:ILLEGAL(unexpected input 'a', expected: first comment character)]`),
		Entry(`3`, `/a`, MatchError(MatchRegexp(`second comment character`)),
			`[2:0.2:ILLEGAL(unexpected input 'a', expected: second comment character)]`),
	)
})
