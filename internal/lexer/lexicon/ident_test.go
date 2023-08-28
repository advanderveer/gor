package lexicon

import (
	"fmt"

	"github.com/advanderveer/gor/internal/lexer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("identifiers", func() {
	DescribeTable("lex identifiers", func(input string, expErr OmegaMatcher, expOutput string) {
		out, err := lexer.New(input, lexIdentifierAndThen(nil)).Lex()
		Expect(err).To(expErr)
		Expect(fmt.Sprint(out)).To(Equal(expOutput))
	},
		Entry(`1`, `a`, BeNil(), `[0:0.0:IDENT(a)]`),
		Entry(`2`, `_x9`, BeNil(), `[0:0.0:IDENT(_x9)]`),
		Entry(`3`, `åβ`, BeNil(), `[0:0.0:IDENT(åβ)]`),
		Entry(`4`, `1abc`, MatchError(MatchRegexp(`expected: letter`)),
			`[1:0.1:ILLEGAL(unexpected input '1', expected: letter)]`),
	)
})
