package lexer

import (
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
		lex := NewScanner("a\n\nb")
		for {
			if lex.Next() == eof {
				break
			}
		}

		Expect(lex.Pos().Offset).To(Equal(4))
		Expect(lex.Pos().Column).To(Equal(1)) // zero-based
		Expect(lex.Pos().Line).To(Equal(2))   // zero-based
	})

	It("next and backup", func() {
		lex := NewScanner("a\n\nb")
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
			lex := NewScanner(`foo`)
			Expect(lex.Keyword("foo")).To(BeTrue())
			Expect(lex.start.Offset).To(Equal(3))
			Expect(lex.curr.Offset).To(Equal(3))
		})

		It("fail", func() {
			lex := NewScanner(`foo`)
			Expect(lex.Keyword("bar")).To(BeFalse())
			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(0))
		})
	})

	Describe("accept", func() {
		It("digits", func() {
			lex := NewScanner(`500 `)
			lex.Accept(isDecimalDigit)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(4))
		})

		It("unicode letters", func() {
			lex := NewScanner(`ūβįβ1`)
			lex.Accept(isUnicodeLetter)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(9))
		})

		It("latin letters", func() {
			lex := NewScanner(`abdc1`)
			lex.Accept(isUnicodeLetter)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(5))
		})
	})

	It("should peek correctly", func() {
		lex := NewScanner(`foobar `)
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

	// It("should error state correctly", func() {
	// 	out := New(`foobar `, func(lc LexControl) StateFunc { return lc.Errorf("foo") }).Run()
	// 	Expect(fmt.Sprint(out)).To(Equal(`[0:0.0:ILLEGAL(foo)]`))
	// })
})
