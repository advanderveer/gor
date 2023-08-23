package lexicon

import (
	"github.com/advanderveer/gor/internal/lexer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("errors", func() {
	It("UnexpectedInputError", func() {
		e1 := unexpectedInput('x', lexer.Pos{Offset: 10, Line: 2, Column: 8}, "digit")
		Expect(e1).To(MatchError(`10:2.8: invalid input, got: 'x' expected: 'digit'`))
	})
})
