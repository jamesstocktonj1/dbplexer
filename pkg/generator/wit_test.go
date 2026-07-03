package generator

import (
	"testing"

	"github.com/jamesstocktonj1/wit"
	"github.com/stretchr/testify/assert"
)

func Test_generateWitFromDatabaseNames(t *testing.T) {
	pkgName := wit.NewPackage("test", "package", "foo").WithVersion("0.2.0")
	dbNames := []string{
		"primary",
		"secondary",
	}

	exp := wit.NewWorld(DefaultWorldName).
		WithExports(wit.NewPackage("test", "package", "foo").WithVersion("0.2.0")).
		WithImports(
			wit.NewPackage("test", "package", "foo").WithVersion("0.2.0").WithName("primary"),
			wit.NewPackage("test", "package", "foo").WithVersion("0.2.0").WithName("secondary"),
		)

	res := generateWitFromDatabaseNames(pkgName, dbNames)
	t.Log(res)

	assert.Equal(t, DefaultNamespace, res.Package.Namespace)
	assert.Equal(t, DefaultPackage, res.Package.Package)
	if assert.NotEmpty(t, res.Worlds) {
		assert.Equal(t, exp, res.Worlds[0])
	}
}
