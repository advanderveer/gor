package lexer

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

type matcherError struct {
	act, exp []Item
	msg      string
}

func (e matcherError) Error() string {
	return fmt.Sprintf("%s: actual: %v, expected: %v", e.msg, e.act, e.exp)
}

// tokenMatcher matches token equality.
type tokenMatcher struct {
	exp []Item
}

func (m tokenMatcher) Match(act any) (bool, error) {
	got, ok := act.([]Item)
	if !ok {
		return false, matcherError{act: nil, exp: m.exp, msg: fmt.Sprintf("actual must be slice of items, got: %T", act)}
	}

	if len(got) != len(m.exp) {
		return false, matcherError{act: got, exp: m.exp, msg: "length not equal"}
	}

	for idx, gotItem := range got {
		expItem := m.exp[idx]

		if expItem.Tok != gotItem.Tok {
			return false, matcherError{
				act: got, exp: m.exp,
				msg: fmt.Sprintf("@%d: got token %s, expected: %s", idx, gotItem.Tok, expItem.Tok),
			}
		}

		if expItem.Val != gotItem.Val {
			return false, matcherError{
				act: got, exp: m.exp,
				msg: fmt.Sprintf("@%d: got token value '%s', expected: '%s'", idx, gotItem.Val, expItem.Val),
			}
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
func TokenValuesToBeEqual(exp []Item) gomega.OmegaMatcher {
	return tokenMatcher{exp: exp}
}
