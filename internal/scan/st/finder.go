package st

import (
	"go/ast"

	"github.com/cloud104/pflagstruct/internal/dir"
	"github.com/cloud104/pflagstruct/internal/syntree"
	"github.com/cloud104/pflagstruct/projscan"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

// Finder is a struct that represents a Go struct type finder.
type Finder struct {
	scanner  *syntree.Scanner       // Scanner is a syntree scanner.
	projects projscan.ProjectFinder // ProjectFinder is a project finder.
	packages projscan.PackageFinder // PackageFinder is a package finder.
}

// NewFinder creates a new Finder instance with a given Scanner, ProjectFinder, and PackageFinder.
func NewFinder(scanner *syntree.Scanner, projects projscan.ProjectFinder, packages projscan.PackageFinder) *Finder {
	return &Finder{scanner: scanner, projects: projects, packages: packages}
}

// FindStructByDirectoryAndName searches for a Go struct type in a directory by its name.
func (f *Finder) FindStructByDirectoryAndName(directory, structName string) (*projscan.Struct, error) {
	directory, err := dir.AbsolutePath(directory)
	if err != nil {
		return nil, err
	}

	proj, err := f.projects.FindProjectByDirectory(directory)
	if err != nil {
		return nil, err
	}

	pkg, err := f.packages.FindPackageByDirectory(directory)
	if err != nil {
		return nil, err
	}

	files, err := f.scanner.ScanDirectory(directory)
	if err != nil {
		return nil, err
	}

	result := make([]*projscan.Struct, 0)

	for filename, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			if spec, ok := n.(*ast.TypeSpec); ok && spec.Name.String() == structName {
				st, f, err := f.navigateUntilStructType(proj, filename, file, spec)
				if err != nil {
					slog.Warn("StructType not found", slog.String("StructName", structName), slog.String("File", filename))
					return false
				}
				result = append(result, &projscan.Struct{
					Package: pkg,
					Name:    structName,
					AST:     &projscan.AST{StructType: st, File: f},
				})
				return false
			}
			return true
		})
	}

	if len(result) > 1 {
		return nil, errors.Errorf("%d structs with the same name were found in the same path %q", len(result), directory)
	}

	if len(result) == 0 {
		return nil, errors.Errorf("no structs were found at the path %q", directory)
	}

	return result[0], nil
}

// navigateUntilStructType navigates until a Go struct type is found in a given file.
func (f *Finder) navigateUntilStructType(proj *projscan.Project, filename string, file *ast.File, spec *ast.TypeSpec) (*ast.StructType, *ast.File, error) {
	switch t := spec.Type.(type) {
	case *ast.StructType:
		// it means that the struct type is a pointer
		return t, file, nil
	case *ast.Ident:
		// it means that the struct type is either a built-in type or a struct from the same package
		st, err := f.FindStructByDirectoryAndName(filename, t.Name)
		if err != nil {
			return nil, nil, err
		}

		return st.AST.StructType, st.AST.File, nil
	case *ast.SelectorExpr:
		// it means that the struct type is a struct from another package
		if x, ok := t.X.(*ast.Ident); ok {
			path, err := syntree.WrapFile(file).FindPackagePathByName(x.Name)
			if err != nil {
				return nil, nil, err
			}

			pkg, err := f.packages.FindPackageByPathAndProject(path, proj)
			if err != nil {
				return nil, nil, err
			}

			st, err := f.FindStructByDirectoryAndName(pkg.Directory, t.Sel.Name)
			if err != nil {
				return nil, nil, err
			}

			return st.AST.StructType, st.AST.File, err
		}
	}

	// if the expression is of a different type, the function returns an error
	return nil, nil, errors.New("no struct was found")
}
