package buildApp

import (
	"fmt"

	"github.com/F1zm0n/enoki/enoki/utils/arch"
	"github.com/F1zm0n/enoki/enoki/utils/pkg/cli"
	pkgremover "github.com/F1zm0n/enoki/enoki/utils/pkg/pkgRemover"
)

func (a *App) LaunchPacman(yNLabel string) error {
	pkgs, err := arch.SearchForPkgs(a.PkgName)
	if err != nil {
		fmt.Println("error getting package info")
		return err
	}
	cli.YesNoPrompt(yNLabel, false)

	pkgs, err = pkgremover.GetPkgToInst(a.PacmanPath, pkgs)
	if err != nil {
		fmt.Println("error finding packages in installed directory")
		return err
	}

	paths, err := arch.UnpakAndInstPacman(a.PacmanRepos, a.Arch, a.PkgName, a.PacmanPath)
	if err != nil {
		fmt.Println("error occured installing packages")
		fmt.Println("please wait before exiting")
		fmt.Println("deleting all packages safely")
		if err = pkgremover.RemovePkgs(paths, a.PacmanPath); err != nil {
			fmt.Println("error deleteing packages")
			return err
		}
		return err
	}

	return nil
}
