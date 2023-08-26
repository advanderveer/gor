package lexicon

// import (
// 	"fmt"

// 	"github.com/advanderveer/gor/internal/lexer"
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("identifiers", func() {
// 	DescribeTable("acceptIdentifier", func(input string, expErr OmegaMatcher, expOutput string) {
// 		lex := lexer.New(input, nil)
// 		err := acceptIdentifier(lex)
// 		Expect(err).To(expErr)

// 		Expect(fmt.Sprint(lex.Lex())).To(Equal(expOutput))
// 	},
// 		Entry(`1`, `a`, BeNil(), `[0:0.0:IDENT(a)]`),
// 		Entry(`2`, `_x9`, BeNil(), `[0:0.0:IDENT(_x9)]`),
// 		Entry(`3`, `åβ`, BeNil(), `[0:0.0:IDENT(åβ)]`),
// 		Entry(`4`, `1abc`, MatchError(`invalid input, got: '1' expected: 'letter'`), `[]`),
// 	)
// })
