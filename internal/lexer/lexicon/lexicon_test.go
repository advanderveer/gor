package lexicon

import (
	"fmt"
	"testing"

	"github.com/advanderveer/gor/internal/lexer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexicon(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer/lexicon")
}

var _ = Describe("file", func() {
	DescribeTable("just package", func(inp string, expOut string) {
		out := lexer.New(inp, LexPackage).Lex()
		Expect(fmt.Sprint(out)).To(Equal(expOut))
	},
		Entry("1", ` package foo`, `[1:0.1:PACKAGE(package) 9:0.9:IDENT(foo)]`),
		Entry("2", " \n "+`package foo`, `[3:1.1:PACKAGE(package) 11:1.9:IDENT(foo)]`),
	)
})
