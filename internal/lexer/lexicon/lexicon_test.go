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
	DescribeTable("just package", func(inp string, expErr OmegaMatcher, expOut string) {
		out, err := lexer.New(inp, LexPackage).Lex()
		Expect(err).To(expErr)
		Expect(fmt.Sprint(out)).To(Equal(expOut))
	},
		Entry("1", ` package foo`, BeNil(), `[1:0.1:PACKAGE(package) 9:0.9:IDENT(foo)]`),
		Entry("2", " \n "+`package foo`, BeNil(), `[3:1.1:PACKAGE(package) 11:1.9:IDENT(foo)]`),
		Entry("3", " \n "+`x`, MatchError(MatchRegexp(`'package' keyword`)),
			`[3:1.1:ILLEGAL(unexpected input 'x', expected: 'package' keyword)]`),
		Entry("4", " \n "+`x`, MatchError(MatchRegexp(`'package' keyword`)),
			`[3:1.1:ILLEGAL(unexpected input 'x', expected: 'package' keyword)]`),
		Entry("5", `
			// comment 1

			// comment 2
			package foo
		`, BeNil(), `[6:1.5:COMMENT( comment 1) 23:3.5:COMMENT( comment 2) 37:4.3:PACKAGE(package) 45:4.11:IDENT(foo)]`),
	)
})
