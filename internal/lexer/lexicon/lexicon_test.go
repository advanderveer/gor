package lexicon

import (
	"fmt"
	"testing"

	"github.com/advanderveer/gor/internal/lexer"
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func TestLexicon(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer/lexicon")
}

var _ = Describe("file", func() {
	DescribeTable("just package", func(inp string, expErr OmegaMatcher, expOut ...lexer.Item) {
		out, err := lexer.New(inp, LexPackage).Lex()
		Expect(err).To(expErr)
		if err == nil {
			Expect(out).To(TokenValuesToBeEqual(expOut))
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
})

// T test utility for creating a token to compare against.
func T(tok token.Token, val string) lexer.Item {
	return lexer.Item{Tok: tok, Val: val}
}

// tokenMatcher matches token equality.
type tokenMatcher struct {
	exp []lexer.Item
}

func (m tokenMatcher) Match(act any) (bool, error) {
	got, ok := act.([]lexer.Item)
	if !ok {
		return false, fmt.Errorf("actual value must be slice of lexer.Item")
	}

	if len(got) != len(m.exp) {
		return false, fmt.Errorf("length %d not equal to expected length: %d", len(got), len(m.exp))
	}

	for idx, gotItem := range got {
		expItem := m.exp[idx]

		if expItem.Tok != gotItem.Tok {
			return false, fmt.Errorf("%v@%d: got token %s, expected: %s", got, idx, gotItem.Tok, expItem.Tok)
		}

		if expItem.Val != gotItem.Val {
			return false, fmt.Errorf("%v@%d: got token value '%s', expected: '%s'", got, idx, gotItem.Val, expItem.Val)
		}
	}

	return true, nil
}

func (m tokenMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to be equivalent to", m.exp)
}

func (m tokenMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "to not be equivalent to", m.exp)
}

// TokenValuesToBeEqual returns a matcher for just looking at token and value of the lexer items.
func TokenValuesToBeEqual(exp []lexer.Item) OmegaMatcher {
	return tokenMatcher{exp: exp}
}
