package generator

type BuildOptions struct {
	PackageName   string
	DatabaseNames []string
	Output        string
}

func DefaultBuildOptions() BuildOptions {
	return BuildOptions{
		PackageName:   "",
		DatabaseNames: []string{},
		Output:        "output",
	}
}
