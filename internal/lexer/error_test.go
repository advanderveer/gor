package lexer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("errors", func() {
	It("UnexpectedInputError", func() {
		e1 := unexpectedInput('x', Pos{10, 2, 8}, "digit")
		Expect(e1).To(MatchError(`10:2.8: invalid input, got: 'x' expected: 'digit'`))
	})
})
