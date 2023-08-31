package parser_test

import (
	gotoken "go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/advanderveer/gor/internal/ast"
	"github.com/advanderveer/gor/internal/parser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/parser")
}

var _ = DescribeTable("testdata",
	func(file string, expErr OmegaMatcher, f func(*ast.File)) {
		src, err := os.ReadFile(file)
		Expect(err).ToNot(HaveOccurred())

		fset := gotoken.NewFileSet()
		root, err := parser.ParseFile(fset, filepath.Base(file), src)
		Expect(err).To(expErr)

		if f != nil {
			f(root)
		}
	},
	Entry("1", "testdata/imports/import1.src", BeNil(), func(f *ast.File) {
		Expect(f.Name.Name).To(Equal("import1"))
		Expect(f.Decls).To(HaveLen(2))

		decl1, decl2 := f.Decls[0].(*ast.GenDecl), f.Decls[1].(*ast.GenDecl)
		Expect(decl1.Specs).To(HaveLen(1))
		Expect(decl2.Specs).To(HaveLen(1))

		spec1 := decl1.Specs[0].(*ast.ImportSpec)
		Expect(spec1.Name).To(BeNil())
		Expect(spec1.Path.Value).To(Equal(`"foo"`))
	}),
	Entry("2", "testdata/imports/import2.src", BeNil(), func(f *ast.File) {
		Expect(f.Name.Name).To(Equal("import2"))
		Expect(f.Decls).To(HaveLen(1))

		decl1 := f.Decls[0].(*ast.GenDecl)
		Expect(decl1.Specs).To(HaveLen(4))

		spec1, spec2, spec3, spec4 := decl1.Specs[0].(*ast.ImportSpec),
			decl1.Specs[1].(*ast.ImportSpec),
			decl1.Specs[2].(*ast.ImportSpec),
			decl1.Specs[3].(*ast.ImportSpec)

		Expect(spec1.Name).To(BeNil())
		Expect(spec1.Path.Value).To(Equal(`"baar"`))
		Expect(spec2.Name.Name).To(Equal(`foo`))
		Expect(spec2.Path.Value).To(Equal(`"dar"`))
		Expect(spec3.Name.Name).To(Equal(`_`))
		Expect(spec3.Path.Value).To(Equal(`"deeed"`))
		Expect(spec4.Name.Name).To(Equal(`.`))
		Expect(spec4.Path.Value).To(Equal(`"xez"`))
	}),

	Entry("3", "testdata/imports/import3.src", MatchError(MatchRegexp(`import path must be a string`)), nil),
	Entry("4", "testdata/imports/import4.src", MatchError(MatchRegexp(`missing import path`)), nil),
)

var _ = DescribeTable("basic parsing",
	func(src string, expErr OmegaMatcher, expFile *ast.File) {
		fset := gotoken.NewFileSet()
		file, err := parser.ParseFile(fset, "file.gor", []byte(src))

		Expect(err).To(expErr)
		Expect(file).To(Equal(expFile))
	},
	Entry("1", `111`, MatchError(MatchRegexp(`expected 'package', found 111`)),
		&ast.File{
			Package: 1,
			Name:    &ast.Ident{Name: "_", NamePos: gotoken.Pos(4)},
		}),
	Entry("1", `"a"`, MatchError(MatchRegexp(`expected 'package', found "a"`)),
		&ast.File{
			Package: 1,
			Name:    &ast.Ident{Name: "_", NamePos: gotoken.Pos(4)},
		}),
	Entry("1", `break`, MatchError(MatchRegexp(`expected 'package', found 'break'`)),
		&ast.File{
			Package: 1,
			Name:    &ast.Ident{Name: "_", NamePos: gotoken.Pos(6)},
		}),
	Entry("1", ` package foo`, BeNil(), &ast.File{
		Package: 2,
		Name:    &ast.Ident{Name: "foo", NamePos: gotoken.Pos(10)},
	}),
)
