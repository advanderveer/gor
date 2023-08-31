package scanner

import (
	goscanner "go/scanner"
	gotoken "go/token"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("numbers", func() {
	It("should format number kind", func() {
		Expect(numberKind('x')).To(Equal(`hexadecimal literal`))
		Expect(numberKind('o')).To(Equal(`octal literal`))
		Expect(numberKind('b')).To(Equal(`binary literal`))
		Expect(numberKind('g')).To(Equal(`decimal literal`))
	})
})

var _ = Describe("private scaner testing", func() {
	It("should handle ending whitespace correctly", func() {
		src := "func\n"
		errs, fset := goscanner.ErrorList{}, gotoken.NewFileSet()
		file := fset.AddFile("file.gor", -1, len(src))

		scanr := &Scanner{}
		scanr.Init(file, []byte(src), errs.Add)

		scanr.Scan()
		scanr.Scan()

		Expect(scanr.lineOffset).To(Equal(5))
		Expect(file.LineCount()).To(Equal(1))
	})
})
