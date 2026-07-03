package generator

import (
	"path/filepath"

	"github.com/jamesstocktonj1/wit"
)

func NewGenerator() *generator {
	return &generator{}
}

type generator struct {
}

func (g *generator) Run(opts BuildOptions) error {
	pkg, err := wit.ParsePackage(opts.PackageName)
	if err != nil {
		return err
	}

	err = GenerateWit(
		pkg,
		opts.DatabaseNames,
		filepath.Join(opts.Output, "wit", "world.wit"),
	)
	if err != nil {
		return err
	}
	return nil
}
