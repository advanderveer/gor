package scanner_test

import (
	"testing"

	"github.com/advanderveer/gor/internal/scanner"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScanner(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/scanner")
}

var _ = DescribeTable("scan one token", func(src string, expTok1 scanner.Token, expLit1 string) {
	file := &scanner.File{}
	scnr := scanner.New(file, []byte(src))

	_, tok, lit := scnr.Scan()
	Expect(tok).To(Equal(expTok1))
	Expect(lit).To(Equal(expLit1))
},
	// keywords
	Entry("number-1", `break`, scanner.BREAK, `break`),

	// number scanning
	Entry("number-1", `1`, scanner.INT, `1`),
	Entry("number-2", `100`, scanner.INT, `100`),
	Entry("number-3", `0xff`, scanner.INT, `0xff`),
	Entry("number-4", `.1`, scanner.FLOAT, `.1`),
	Entry("number-5", `1.1`, scanner.FLOAT, `1.1`),

	// specials
	Entry("special-1", `$`, scanner.ILLEGAL, `$`),
	Entry("special-2", ``, scanner.EOF, ``),

	// strings
	Entry("strings-1", `"a"`, scanner.STRING, `"a"`),
	Entry("strings-2", "`a`", scanner.STRING, "`a`"),

	// operators
	Entry("op-1", `/`, scanner.QUO, ``),
	Entry("op-2", `/=`, scanner.QUO_ASSIGN, ``),
	Entry("op-3", `*`, scanner.MUL, ``),
	Entry("op-4", `*=`, scanner.MUL_ASSIGN, ``),
	Entry("op-5", `:`, scanner.COLON, ``),
	Entry("op-6", `:=`, scanner.DEFINE, ``),
	Entry("op-7", `+`, scanner.ADD, ``),
	Entry("op-8", `+=`, scanner.ADD_ASSIGN, ``),
	Entry("op-9", `++`, scanner.INC, ``),
	Entry("op-10", `-`, scanner.SUB, ``),
	Entry("op-11", `-=`, scanner.SUB_ASSIGN, ``),
	Entry("op-12", `--`, scanner.DEC, ``),
	Entry("op-13", `=`, scanner.ASSIGN, ``),
	Entry("op-14", `==`, scanner.EQL, ``),
	Entry("op-15", `!`, scanner.NOT, ``),
	Entry("op-16", `!=`, scanner.NEQ, ``),
	Entry("op-17", `<`, scanner.LSS, ``),
	Entry("op-18", `<-`, scanner.ARROW, ``),
	Entry("op-19", `<=`, scanner.LEQ, ``),
	Entry("op-20", `<<`, scanner.SHL, ``),
	Entry("op-21", `<<=`, scanner.SHL_ASSIGN, ``),
	Entry("op-22", `>`, scanner.GTR, ``),
	Entry("op-23", `>=`, scanner.GEQ, ``),
	Entry("op-24", `>>`, scanner.SHR, ``),
	Entry("op-25", `>>=`, scanner.SHR_ASSIGN, ``),
	Entry("op-25", `&^`, scanner.AND_NOT, ``),
	Entry("op-26", `&^=`, scanner.AND_NOT_ASSIGN, ``),
	Entry("op-27", `&`, scanner.AND, ``),
	Entry("op-28", `&=`, scanner.AND_ASSIGN, ``),
	Entry("op-29", `&&`, scanner.LAND, ``),
	Entry("op-30", `|`, scanner.OR, ``),
	Entry("op-31", `|=`, scanner.OR_ASSIGN, ``),
	Entry("op-32", `||`, scanner.LOR, ``),
	Entry("op-33", `%`, scanner.REM, ``),
	Entry("op-34", `%=`, scanner.REM_ASSIGN, ``),
	Entry("op-35", `^`, scanner.XOR, ``),
	Entry("op-36", `^=`, scanner.XOR_ASSIGN, ``),

	// various other tokens
	Entry("other-1", `.`, scanner.PERIOD, ``),
	Entry("other-2", `...`, scanner.ELLIPSIS, ``),
	Entry("other-2", `,`, scanner.COMMA, ``),
	Entry("other-2", `;`, scanner.SEMICOLON, `;`),

	// parenthesis, brackets and braces
	Entry("pbb-1", `(`, scanner.LPAREN, ``),
	Entry("pbb-2", `)`, scanner.RPAREN, ``),
	Entry("pbb-3", `[`, scanner.LBRACK, ``),
	Entry("pbb-4", `]`, scanner.RBRACK, ``),
	Entry("pbb-5", `{`, scanner.LBRACE, ``),
	Entry("pbb-6", `}`, scanner.RBRACE, ``),

	// comments
	Entry("comment-1", `// foo`, scanner.COMMENT, `// foo`),
)

var _ = DescribeTable("scan three tokens", func(src string,
	expTok1 scanner.Token, expLit1 string,
	expTok2 scanner.Token, expLit2 string,
	expTok3 scanner.Token, expLit3 string,
) {
	file := &scanner.File{}
	scnr := scanner.New(file, []byte(src))

	_, tok, lit := scnr.Scan()
	Expect(tok).To(Equal(expTok1), "token 1")
	Expect(lit).To(Equal(expLit1), "token 1")

	_, tok, lit = scnr.Scan()
	Expect(tok).To(Equal(expTok2), "token 2")
	Expect(lit).To(Equal(expLit2), "token 2")

	_, tok, lit = scnr.Scan()
	Expect(tok).To(Equal(expTok3), "token 3")
	Expect(lit).To(Equal(expLit3), "token 3")
},
	Entry("semi before eof", `x`,
		scanner.IDENT, `x`,
		scanner.SEMICOLON, "\n",
		scanner.EOF, "",
	),

	Entry("semi after ident", `x
	break`,
		scanner.IDENT, `x`,
		scanner.SEMICOLON, "\n",
		scanner.BREAK, `break`,
	),

	Entry("comment with newline", `// foo
	x`,
		scanner.COMMENT, "// foo",
		scanner.IDENT, "x",
		scanner.SEMICOLON, "\n",
	),
)
