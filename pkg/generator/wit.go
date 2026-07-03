package generator

import "github.com/jamesstocktonj1/wit"

const (
	DefaultWorldName = "shim"
	DefaultNamespace = "dbplexer"
	DefaultPackage   = "shim"
)

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
