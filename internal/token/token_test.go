package token_test

import (
	"fmt"
	"testing"

	"github.com/advanderveer/gor/internal/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestToken(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/token")
}

var _ = Describe("token", func() {
	It("should print expected tokens", func() {
		Expect(fmt.Sprint(token.BREAK)).To(Equal("break"))
	})

	It("should print unexpected tokens", func() {
		Expect(fmt.Sprint(token.Token(11111))).To(Equal("Token(11111)"))
	})

	It("should correctly do IsLiteral", func() {
		Expect(token.STRING.IsLiteral()).To(BeTrue())
		Expect(token.BREAK.IsLiteral()).To(BeFalse())
	})

	It("should return correct presedence", func() {
		Expect(token.RETURN.Precedence()).To(Equal(token.LowestPrec))
		Expect(token.LOR.Precedence()).To(Equal(1))
		Expect(token.LAND.Precedence()).To(Equal(2))
		Expect(token.EQL.Precedence()).To(Equal(3))
		Expect(token.ADD.Precedence()).To(Equal(4))
		Expect(token.MUL.Precedence()).To(Equal(5))
	})

	It("should lookup", func() {
		Expect(token.Lookup("break")).To(Equal(token.BREAK))
	})
})
