package arch

import "github.com/F1zm0n/enoki/enoki/utils/entity"

type PacmanApp struct {
	RootPkg     string
	Repos       []string
	Arches      []string
	Path        string
	MirrorLinks []string
	Deps        []string
}

type IPacmanApp interface {
	GetFromPacman(info entity.ArchInfo) ([]byte, error)
	SearchForPkgs() error
	GetMirrorHost(conf map[string]string) error
}

// TODO: сделай нормальный конфиг парсер и миррор линки нормальные

func GetPacmanApp(rootPkg string, repos []string, arches []string) IPacmanApp {
	return &PacmanApp{
		RootPkg: rootPkg,
		Repos:   repos,
		Arches:  arches,
	}
}
