package generator

import (
	"os"

	"github.com/jamesstocktonj1/wit"
)

const (
	DefaultWorldName = "shim"
	DefaultNamespace = "dbplexer"
	DefaultPackage   = "shim"
)

func GenerateWit(pkg wit.Package, names []string, witPath string) error {
	witGen := generateWitFromDatabaseNames(pkg, names)

	f, err := os.Create(witPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return wit.NewEncoder(f).Encode(witGen)
}

func generateWitFromDatabaseNames(pkg wit.Package, names []string) wit.Wit {
	imports := []wit.Importable{}
	for _, name := range names {
		imports = append(imports, pkg.WithName(name))
	}

	return wit.NewWit().
		WithPackage(wit.NewPackage(DefaultNamespace, DefaultPackage)).
		WithWorld(
			wit.NewWorld(DefaultWorldName).
				WithExports(pkg).
				WithImports(imports...),
		)
}
