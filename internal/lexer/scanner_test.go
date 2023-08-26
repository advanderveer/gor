package lexer

import (
	"unicode"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("positions", func() {
	It("next only", func() {
		lex := NewScanner("a\n\nb")
		for {
			if lex.Next() == EOF {
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
			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(3))
		})

		It("fail", func() {
			lex := NewScanner(`foo`)
			Expect(lex.Keyword("bar")).To(BeFalse())
			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(0))
		})
	})

	Describe("accept/skip", func() {
		It("accept digits", func() {
			lex := NewScanner(`500 `)
			lex.Accept(unicode.IsDigit)

			Expect(lex.start.Offset).To(Equal(0))
			Expect(lex.curr.Offset).To(Equal(3))

			Expect(lex.Value()).To(Equal(`500`))
		})

		It("skip digits", func() {
			lex := NewScanner(`213a `)
			lex.Skip(unicode.IsDigit)

			Expect(lex.start.Offset).To(Equal(3))
			Expect(lex.curr.Offset).To(Equal(3))
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
})
