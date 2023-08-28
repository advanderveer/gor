package lexer

import (
	"github.com/advanderveer/gor/internal/lexer/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("matcher", func() {
	It("shoul format messages", func() {
		Expect(TokenValuesToBeEqual([]Item{}).FailureMessage("a")).To(MatchRegexp(`to be equivalent to`))
		Expect(TokenValuesToBeEqual([]Item{}).NegatedFailureMessage("a")).To(MatchRegexp(`to not be equivalent to`))
	})

	It("should error on wrong type", func() {
		succeed, err := TokenValuesToBeEqual([]Item{}).Match("a")
		Expect(err).To(MatchError(MatchRegexp(`actual must be slice of items, got: string`)))
		Expect(succeed).To(BeFalse())
	})

	It("should error on wrong length", func() {
		succeed, err := TokenValuesToBeEqual([]Item{}).Match([]Item{{}})
		Expect(err).To(MatchError(MatchRegexp(`length not equal`)))
		Expect(succeed).To(BeFalse())
	})

	It("should error on wrong item token", func() {
		succeed, err := TokenValuesToBeEqual([]Item{{Tok: token.COMMENT}}).Match([]Item{{Tok: token.IDENT}})
		Expect(err).To(MatchError(MatchRegexp(`got token IDENT, expected: COMMENT`)))
		Expect(succeed).To(BeFalse())
	})

	It("should error on wrong item value", func() {
		succeed, err := TokenValuesToBeEqual(
			[]Item{{Tok: token.COMMENT, Val: `abc`}}).Match([]Item{{Tok: token.COMMENT, Val: `bar`}})
		Expect(err).To(MatchError(MatchRegexp(`got token value 'bar', expected: 'abc'`)))
		Expect(succeed).To(BeFalse())
	})

	It("should match", func() {
		succeed, err := TokenValuesToBeEqual(
			[]Item{{Tok: token.COMMENT, Val: `abc`}}).Match([]Item{{Tok: token.COMMENT, Val: `abc`}})
		Expect(err).To(Not(HaveOccurred()))
		Expect(succeed).To(BeTrue())
	})
})
