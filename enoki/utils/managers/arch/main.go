package arch

type PacmanApp struct {
	RootPkg     string
	Repos       []string
	Arches      []string
	Path        string
	MirrorLinks []string
	Deps        []string
}

// IPacmanApp interface provides pacman's methods such as
// GetFromPacman
type IPacmanApp interface {
	// SearchForPkgs method searches for all the of the dependencies of dependencies
	// and so on from pacman repository using 2 function to search for packages:
	//
	// GetPkgSo and GetPkg functions, and after that appends
	// to a PacmanApp struct slice of string with dependencies
	//
	// returns apperror.ErrNoSuchPkg if no package was found
	SearchForPkgs() error
	// GetMirrorHost accepts config map and appends a slice of mirror hosts
	// to a PacmanApp.MirrorLinks return an error
	GetMirrorHost(conf map[string]string) error
}

// TODO: сделай нормальный конфиг парсер и миррор линки нормальные

// GetPacmanApp returns IPacmanApp interface and provides methods of pacman
func GetPacmanApp(rootPkg string, repos []string, arches []string) IPacmanApp {
	return &PacmanApp{
		RootPkg: rootPkg,
		Repos:   repos,
		Arches:  arches,
	}
}
