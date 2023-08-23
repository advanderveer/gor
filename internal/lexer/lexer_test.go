package lexer

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLexer(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/lexer")
}

var _ = Describe("positions", func() {
	It("next only", func() {
		lex := New("a\n\nb", nil)
		for {
			if lex.Next() == eof {
				break
			}
		}

		Expect(lex.curr.Offset).To(Equal(4))
		Expect(lex.curr.Column).To(Equal(1)) // zero-based
		Expect(lex.curr.Line).To(Equal(2))   // zero-based
	})

	It("next and backup", func() {
		lex := New("a\n\nb", nil)
		lex.Next()
		lex.Next()

		Expect(lex.curr.Offset).To(Equal(2))
		Expect(lex.curr.Line).To(Equal(1))

		lex.Backup()
		Expect(lex.curr.Offset).To(Equal(1))
		Expect(lex.curr.Line).To(Equal(0))

		lex.Next()
		lex.Next()
		lex.Next()

		Expect(lex.curr.Offset).To(Equal(4))
		Expect(lex.curr.Line).To(Equal(2))
		Expect(lex.curr.Column).To(Equal(1))

		lex.Backup()

		Expect(lex.curr.Offset).To(Equal(3))
		Expect(lex.curr.Line).To(Equal(2))
		Expect(lex.curr.Column).To(Equal(0))
	})
})

var _ = Describe("scanning controls", func() {
	Describe("keywords", func() {
		It("succeed", func() {
			lex := New(`foo`, nil)
			Expect(lex.Keyword("foo")).To(BeTrue())
			Expect(lex.start.Offset).To(Equal(3))
			Expect(lex.curr.Offset).To(Equal(3))
		})

		It("fail", func() {
			lex := New(`foo`, nil)
			Expect(lex.Keyword("bar")).To(BeFalse())
			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(0))
		})
	})

	Describe("accept", func() {
		It("digits", func() {
			lex := New(`500 `, nil)
			lex.Accept(isDecimalDigit)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(4))
		})

		It("unicode letters", func() {
			lex := New(`ūβįβ1`, nil)
			lex.Accept(isUnicodeLetter)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(9))
		})

		It("latin letters", func() {
			lex := New(`abdc1`, nil)
			lex.Accept(isUnicodeLetter)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(5))
		})
	})

	It("should peek correctly", func() {
		lex := New(`foobar `, nil)
		Expect(lex.Peek()).To(Equal('f'))
		Expect(lex.start.Offset).To(Equal(0))
		Expect(lex.curr.Offset).To(Equal(0))
		Expect(lex.Peek()).To(Equal('f'))

		Expect(lex.Next()).To(Equal('f'))
		Expect(lex.start.Offset).To(Equal(0))
		Expect(lex.curr.Offset).To(Equal(1))

		Expect(lex.Peek()).To(Equal('o'))
		Expect(lex.Peek()).To(Equal('o'))
		Expect(lex.start.Offset).To(Equal(0))
		Expect(lex.curr.Offset).To(Equal(1))
	})

	It("should error state correctly", func() {
		out := New(`foobar `, func(lc LexControl) StateFunc { return lc.Errorf("foo") }).Run()
		Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:ILLEGAL(foo)]`))
	})
})
