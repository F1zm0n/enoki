package entity

import "time"

type ArchInfo struct {
	PkgName        string     `json:"pkgname"`
	PkgBase        string     `json:"pkgbase"`
	Repo           string     `json:"repo"`
	Arch           string     `json:"arch"`
	PkgVer         string     `json:"pkgver"`
	PkgRel         string     `json:"pkgrel"`
	Epoch          int        `json:"epoch"`
	PkgDesc        string     `json:"pkgdesc"`
	URL            string     `json:"url"`
	Filename       string     `json:"filename"`
	CompressedSize int        `json:"compressed_size"`
	InstalledSize  int        `json:"installed_size"`
	BuildDate      *time.Time `json:"build_date"`
	LastUpdate     *time.Time `json:"last_update"`
	FlagDate       *time.Time `json:"flag_date"`
	Maintainers    []string   `json:"maintainers"`
	Packager       string     `json:"packager"`
	Groups         []string   `json:"groups"`
	Licenses       []string   `json:"licenses"`
	Conflicts      []string   `json:"conflicts"`
	Provides       []string   `json:"provides"`
	Replaces       []string   `json:"replaces"`
	Depends        []string   `json:"depends"`
	OptDepends     []string   `json:"optdepends"`
	MakeDepends    []string   `json:"makedepends"`
	CheckDepends   []string   `json:"checkdepends"`
}

type InstallInfo struct {
	PackageNames []string
}
