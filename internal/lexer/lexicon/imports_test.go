package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("imports", func(inp string, expErr OmegaMatcher, expOut ...lexer.Item) {
	out, err := lexer.New(inp, lexImports).Lex()
	Expect(err).To(expErr)
	if err == nil {
		Expect(out).To(lexer.TokenValuesToBeEqual(expOut))
	}
},
	Entry("1", ` import ()`, BeNil(),
		T(token.IMPORT, `import`),
		T(token.LPAREN, `(`),
		T(token.RPAREN, `)`),
	),
	Entry("2", ` import("foo")`, BeNil(),
		T(token.IMPORT, `import`),
		T(token.LPAREN, `(`),
		T(token.STRING, `foo`),
		T(token.RPAREN, `)`),
	),
	Entry("3", `// some comment
	import (    "foo" 
	"bar")`, BeNil(),
		T(token.COMMENT, ` some comment`),
		T(token.IMPORT, `import`),
		T(token.LPAREN, `(`),
		T(token.STRING, `foo`),
		T(token.STRING, `bar`),
		T(token.RPAREN, `)`),
	),
	Entry("4", ` asdf`, MatchError(
		MatchRegexp(`expected: 'import' keyword`))),
	Entry("4", `importa`, MatchError(
		MatchRegexp(`expected: left parenthesis`))),
	Entry("6", ` import (  1)`, MatchError(
		MatchRegexp(`expected: white space, letter or string literal`))),
	Entry("7", `!`, MatchError(
		MatchRegexp(`expected: white space, comment or unicode letter`))),

	Entry("8", ` import(foo"foo""bar")`, BeNil(),
		T(token.IMPORT, `import`),
		T(token.LPAREN, `(`),
		T(token.IDENT, `foo`),
		T(token.STRING, `foo`),
		T(token.STRING, `bar`),
		T(token.RPAREN, `)`),
	),
)
