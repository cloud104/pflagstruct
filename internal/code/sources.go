package code

import (
	"fmt"
	"os"

	"github.com/cloud104/pflagstruct/projscan"
	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type FlagSource struct {
	Struct  *projscan.Struct
	Package *projscan.Package
	Blocks  []Block

	variables []string
	imports   map[string]string
}

func (f *FlagSource) ImportName(path, name string) {
	if f.imports == nil {
		f.imports = map[string]string{
			"github.com/spf13/cobra": "cobra",
			"github.com/spf13/pflag": "pflag",
		}
	}

	f.imports[path] = name
	f.variables = lo.Uniq(append(f.variables, name))
}

func (f *FlagSource) File() *jen.File {
	file := jen.NewFilePathName(f.Package.Path, f.Package.Name)
	file.HeaderComment("Code generated by pflagstruct. DO NOT EDIT.")
	file.HeaderComment(fmt.Sprintf("//go:generate pflagstruct --package %s --struct-name %s", f.Struct.Package.Path, f.Struct.Name))
	file.ImportNames(f.imports)

	for _, block := range f.Blocks {
		file.Add(block.Statement())
	}

	return file
}

func (f *FlagSource) Bytes() []byte {
	file := f.File()
	return []byte(fmt.Sprintf("%#v", file))
}

func (f *FlagSource) Print() {
	bytes := f.Bytes()
	fmt.Println(string(bytes))
}

func (f *FlagSource) WriteFile(filepath string) error {
	bytes := f.Bytes()
	if err := os.WriteFile(filepath, bytes, 0o644); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
