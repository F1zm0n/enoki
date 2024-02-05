/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"

	"github.com/F1zm0n/enoki/enoki/utils/arch"
	"github.com/F1zm0n/enoki/enoki/utils/pkg/cli"
	getarchitecture "github.com/F1zm0n/enoki/enoki/utils/pkg/get_architecture"
	pkgremover "github.com/F1zm0n/enoki/enoki/utils/pkg/pkgRemover"
)

const (
	yNLabel = "are you sure you want to do this?"
)

var Repositories = []string{"extra", "core", "multilib"}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install: installs the package from the different repositories based on given flag",
	Long: `installs the packages directly to your system from different repositories must specify 
	the flag from where you want to install the package, currently supports only pacman`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		installedDeps := make([]string, 0)

		archi := runtime.GOARCH

		archi = getarchitecture.GetArchitecture(archi)

		// timeout, err := time.ParseDuration(os.Getenv("CTX_TIMEOUT")) // сделать адекватно
		// if err != nil {
		// fmt.Println("uknown error timeout not specified")
		// return
		// }
		timeout := 20 * time.Second // убрать эту хуйню

		pkgName := args[0]
		if pacman {
			info, err := arch.GetPkgInfoBoth(timeout, Repositories, archi, pkgName)
			if err != nil {
				fmt.Println("no package found or another error ", err)
				return
			}

			// packages, err = arch.GetPackageList(info.Depends)

			YOrN := cli.YesNoPrompt(yNLabel, false)
			if !YOrN {
				return
			}

			pacmanPath := os.Getenv("PACMAN_PATH")

			pacmanPath = "~/Files/" // Поменять это убрать

			installedDeps, err = arch.InstDepsPacman(
				timeout,
				Repositories,
				info,
				archi,
				pacmanPath,
				installedDeps,
			)
			if err != nil {
				fmt.Println("error occured while trying to install packages ", err)
				fmt.Println("please wait")
				fmt.Println("exiting with safe mode")
				err = pkgremover.RemovePkgs(installedDeps, pacmanPath)
				if err != nil {
					fmt.Println("couldn't remove all packages ", err)
					return
				}
				return
			}

			installedDeps, err = arch.UnpakAndInstPacman(
				timeout,
				installedDeps,
				Repositories,
				info.Arch,
				pkgName,
				pacmanPath,
			)
			if err != nil {
				fmt.Println("error occured while trying to install packages")
				fmt.Println("please wait")
				fmt.Println("exiting with safe mode")
				err = pkgremover.RemovePkgs(installedDeps, pacmanPath)
				if err != nil {
					fmt.Println("couldn't remove all packages ", err)
					return
				}
				return
			}
		}
		fmt.Println("installed everything")
	},
}
var pacman bool

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")
	installCmd.Flags().BoolVarP(&pacman, "pacman", "P", false, "install package from pacman")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	installCmd.MarkFlagsOneRequired("pacman")
}
