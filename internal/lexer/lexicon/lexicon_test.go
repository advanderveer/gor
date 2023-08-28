package lexicon

import (
	"testing"

	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexicon(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer/lexicon")
}

var _ = DescribeTable("package", func(inp string, expErr OmegaMatcher, expOut ...lexer.Item) {
	out, err := lexer.New(inp, LexPackage).Lex()
	Expect(err).To(expErr)
	if err == nil {
		Expect(out).To(lexer.TokenValuesToBeEqual(expOut))
	}
},
	Entry("1", ` package foo`, BeNil(),
		T(token.PACKAGE, `package`),
		T(token.IDENT, `foo`)),
	Entry("2", " \n "+`package foo`, BeNil(),
		T(token.PACKAGE, `package`),
		T(token.IDENT, `foo`)),
	Entry("3", " \n "+`x`,
		MatchError(MatchRegexp(`expected: 'package' keyword`))),
	Entry("4", `
			// comment 1

			// comment 2
			package foo
		`, BeNil(),
		T(token.COMMENT, ` comment 1`),
		T(token.COMMENT, ` comment 2`),
		T(token.PACKAGE, `package`),
		T(token.IDENT, `foo`),
	),
	Entry("5", ``,
		MatchError(MatchRegexp(`white space, comment or unicode letter`)), nil),
)

// T test utility for creating a token to compare against.
func T(tok token.Token, val string) lexer.Item {
	return lexer.Item{Tok: tok, Val: val}
}
