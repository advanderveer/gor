package scanner_test

import (
	gotoken "go/token"
	"testing"

	"github.com/advanderveer/gor/internal/scanner"
	"github.com/advanderveer/gor/internal/token"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScanner(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/scanner")
}

var _ = Describe("errors", func() {
	It("should report correct line and file", func() {
		var errPos gotoken.Position
		var errMsg string

		fset := gotoken.NewFileSet()
		bytes1 := []byte("0x12\x00aaa")
		file1 := fset.AddFile("file1.gor", -1, len(bytes1))
		scanr := &scanner.Scanner{}
		scanr.Init(file1, bytes1, func(pos gotoken.Position, msg string) {
			errPos, errMsg = pos, msg
		})

		pos, tok, lit := scanr.Scan()
		Expect(pos).To(Equal(gotoken.Pos(0)))
		Expect(tok).To(Equal(token.INT))
		Expect(lit).To(Equal("0x12"))

		Expect(errPos).To(Equal(gotoken.Position{Filename: "file1.gor", Offset: 4, Line: 1, Column: 5}))
		Expect(errMsg).To(MatchRegexp(`illegal character NUL`))
	})
})

var _ = DescribeTable("scan one token", func(src string, expTok1 token.Token, expLit1 string) {
	file := &gotoken.File{}
	scnr := &scanner.Scanner{}
	scnr.Init(file, []byte(src), nil)

	_, tok, lit := scnr.Scan()
	Expect(tok).To(Equal(expTok1))
	Expect(lit).To(Equal(expLit1))
},
	// keywords
	Entry("keyword-1", "break", token.BREAK, `break`),
	Entry("keyword-2", "case", token.CASE, `case`),
	Entry("keyword-3", "chan", token.CHAN, `chan`),
	Entry("keyword-4", "const", token.CONST, `const`),
	Entry("keyword-5", "continue", token.CONTINUE, `continue`),

	Entry("keyword-6", "default", token.DEFAULT, `default`),
	Entry("keyword-7", "defer", token.DEFER, `defer`),
	Entry("keyword-8", "else", token.ELSE, `else`),
	Entry("keyword-9", "fallthrough", token.FALLTHROUGH, `fallthrough`),
	Entry("keyword-10", "for", token.FOR, `for`),

	Entry("keyword-11", "func", token.FUNC, `func`),
	Entry("keyword-12", "go", token.GO, `go`),
	Entry("keyword-13", "goto", token.GOTO, `goto`),
	Entry("keyword-14", "if", token.IF, `if`),
	Entry("keyword-15", "import", token.IMPORT, `import`),

	Entry("keyword-16", "interface", token.INTERFACE, `interface`),
	Entry("keyword-17", "map", token.MAP, `map`),
	Entry("keyword-18", "package", token.PACKAGE, `package`),
	Entry("keyword-19", "range", token.RANGE, `range`),
	Entry("keyword-20", "return", token.RETURN, `return`),

	Entry("keyword-21", "select", token.SELECT, `select`),
	Entry("keyword-22", "struct", token.STRUCT, `struct`),
	Entry("keyword-23", "switch", token.SWITCH, `switch`),
	Entry("keyword-24", "type", token.TYPE, `type`),
	Entry("keyword-25", "var", token.VAR, `var`),

	// number scanning
	Entry("number-26", `1`, token.INT, `1`),
	Entry("number-27", `100`, token.INT, `100`),
	Entry("number-28", `0xff`, token.INT, `0xff`),
	Entry("number-29", `.1`, token.FLOAT, `.1`),
	Entry("number-30", `1.1`, token.FLOAT, `1.1`),

	// specials
	Entry("special-31", `$`, token.ILLEGAL, `$`),
	Entry("special-32", ``, token.EOF, ``),

	// strings
	Entry("strings-1", `"a"`, token.STRING, `"a"`),
	Entry("strings-2", "`a`", token.STRING, "`a`"),

	// operators
	Entry("op-1", `/`, token.QUO, ``),
	Entry("op-2", `/=`, token.QUO_ASSIGN, ``),
	Entry("op-3", `*`, token.MUL, ``),
	Entry("op-4", `*=`, token.MUL_ASSIGN, ``),
	Entry("op-5", `:`, token.COLON, ``),
	Entry("op-6", `:=`, token.DEFINE, ``),
	Entry("op-7", `+`, token.ADD, ``),
	Entry("op-8", `+=`, token.ADD_ASSIGN, ``),
	Entry("op-9", `++`, token.INC, ``),
	Entry("op-10", `-`, token.SUB, ``),
	Entry("op-11", `-=`, token.SUB_ASSIGN, ``),
	Entry("op-12", `--`, token.DEC, ``),
	Entry("op-13", `=`, token.ASSIGN, ``),
	Entry("op-14", `==`, token.EQL, ``),
	Entry("op-15", `!`, token.NOT, ``),
	Entry("op-16", `!=`, token.NEQ, ``),
	Entry("op-17", `<`, token.LSS, ``),
	Entry("op-18", `<-`, token.ARROW, ``),
	Entry("op-19", `<=`, token.LEQ, ``),
	Entry("op-20", `<<`, token.SHL, ``),
	Entry("op-21", `<<=`, token.SHL_ASSIGN, ``),
	Entry("op-22", `>`, token.GTR, ``),
	Entry("op-23", `>=`, token.GEQ, ``),
	Entry("op-24", `>>`, token.SHR, ``),
	Entry("op-25", `>>=`, token.SHR_ASSIGN, ``),
	Entry("op-25", `&^`, token.AND_NOT, ``),
	Entry("op-26", `&^=`, token.AND_NOT_ASSIGN, ``),
	Entry("op-27", `&`, token.AND, ``),
	Entry("op-28", `&=`, token.AND_ASSIGN, ``),
	Entry("op-29", `&&`, token.LAND, ``),
	Entry("op-30", `|`, token.OR, ``),
	Entry("op-31", `|=`, token.OR_ASSIGN, ``),
	Entry("op-32", `||`, token.LOR, ``),
	Entry("op-33", `%`, token.REM, ``),
	Entry("op-34", `%=`, token.REM_ASSIGN, ``),
	Entry("op-35", `^`, token.XOR, ``),
	Entry("op-36", `^=`, token.XOR_ASSIGN, ``),

	// various other tokens
	Entry("other-1", `.`, token.PERIOD, ``),
	Entry("other-2", `...`, token.ELLIPSIS, ``),
	Entry("other-2", `,`, token.COMMA, ``),
	Entry("other-2", `;`, token.SEMICOLON, `;`),

	// parenthesis, brackets and braces
	Entry("pbb-1", `(`, token.LPAREN, ``),
	Entry("pbb-2", `)`, token.RPAREN, ``),
	Entry("pbb-3", `[`, token.LBRACK, ``),
	Entry("pbb-4", `]`, token.RBRACK, ``),
	Entry("pbb-5", `{`, token.LBRACE, ``),
	Entry("pbb-6", `}`, token.RBRACE, ``),

	// comments
	Entry("comment-1", `// foo`, token.COMMENT, `// foo`),
)

var _ = DescribeTable("scan three tokens", func(src string,
	expTok1 token.Token, expLit1 string,
	expTok2 token.Token, expLit2 string,
	expTok3 token.Token, expLit3 string,
) {
	file := &gotoken.File{}
	scnr := &scanner.Scanner{}
	scnr.Init(file, []byte(src), nil)

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
		token.IDENT, `x`,
		token.SEMICOLON, "\n",
		token.EOF, "",
	),

	Entry("semi after ident", `x
	break`,
		token.IDENT, `x`,
		token.SEMICOLON, "\n",
		token.BREAK, `break`,
	),

	Entry("comment with newline", `// foo
	x`,
		token.COMMENT, "// foo",
		token.IDENT, "x",
		token.SEMICOLON, "\n",
	),
)
