package buildApp

import (
	"errors"
	"fmt"

	"github.com/F1zm0n/enoki/enoki/utils/arch"
	"github.com/F1zm0n/enoki/enoki/utils/pkg/cli"
	pkgremover "github.com/F1zm0n/enoki/enoki/utils/pkg/pkgRemover"
)

const configPath = "~/.config/enoki/enoki.conf"

func (a *App) LaunchPacman(yNLabel string) error {
	pkgs, err := arch.SearchForPkgs(a.PkgName)
	if err != nil {
		fmt.Println("error getting package info")
		return err
	}

	fmt.Println(pkgs)

	// pkgs, err = pkgremover.GetPkgToInst(a.PacmanPath, pkgs)
	// if err != nil {
	// 	fmt.Println("error finding packages in installed directory")
	// 	return err
	// }

	yN := cli.YesNoPrompt(yNLabel, false)
	if !yN {
		return nil
	}

	fmt.Println("starting download")
	// conf, err := configs.ReadConfig(configPath)
	// if err != nil {
	// 	return err
	// }
	// host, err := arch.GetMirrorHost(conf)
	// if err != nil {
	// 	return err
	// }
	host := "mirror.hoster.kz"

	for _, pkg := range pkgs {
		fmt.Println(pkg)
		info, err := arch.GetPkg(pkg)
		if err != nil {
			fmt.Println("error")
			info, err = arch.GetPkgSo(pkg)
			if err != nil {
				if errors.Is(err, arch.ErrNoSuchPackage) {
					continue
				}
				fmt.Println(err)
				return err
			}
		}
		paths, err := arch.UnpakAndInstPacman(info, host, a.PacmanPath)
		if err != nil {
			fmt.Println("error occured installing packages ", err)
			fmt.Println("please wait before exiting")
			fmt.Println("deleting all packages safely")
			if err = pkgremover.RemovePkgs(paths, a.PacmanPath); err != nil {
				fmt.Println("error deleteing packages")
				return err
			}
			return err
		}
	}

	return nil
}
