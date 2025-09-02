package jen

import (
	"io"
	"path"
)

func ImportPackages(paths []string) *Statement {
	return newStatement().ImportPackages(paths)
}

func (s *Statement) ImportPackage(name string) *Statement {
	t := pkgToken{
		path:  name,
		alias: path.Base(name), // 默认的别名会去掉下划线
	}
	*s = append(*s, t)
	return s
}

func (s *Statement) ImportPackages(paths []string) *Statement {
	for _, path := range paths {
		s.ImportPackage(path)
	}
	return s
}

var Raw = Op

type pkgToken struct {
	path  string
	alias string
}

func (t pkgToken) render(f *File, w io.Writer, s *Statement) error {
	f.ImportName(t.path, t.alias)
	f.register(t.path)
	return nil
}

func (t pkgToken) isNull(f *File) bool {
	return false
}
