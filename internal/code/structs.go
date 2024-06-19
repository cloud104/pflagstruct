package code

import (
	"github.com/cloud104/pflagstruct/projscan"
	"github.com/dave/jennifer/jen"
)

type FlagsBuilderStruct struct {
	Name   string
	Struct *projscan.Struct
}

func (cfs *FlagsBuilderStruct) Statement() *jen.Statement {
	fields := []jen.Code{
		jen.Id("flags").Op("*").Qual("github.com/spf13/pflag", "FlagSet"),
	}

	return jen.Type().Id(cfs.Name).Struct(fields...)
}
