/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	buildApp "github.com/F1zm0n/enoki/enoki/utils/app"
	getarchitecture "github.com/F1zm0n/enoki/enoki/utils/pkg/get_architecture"
)

const (
	yNLabel = "are you sure you want to do this?"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install: installs the package from the different repositories based on given flag",
	Long: `installs the packages directly to your system from different repositories must specify 
	the flag from where you want to install the package, currently supports only pacman`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		archi := getarchitecture.GetArchitecture()

		pkgName := args[0]

		app := &buildApp.App{
			Arch:    archi,
			PkgName: pkgName,
		}

		if pacman {
			app.PacmanPath = "~/Files/"
			app.PacmanRepos = []string{"extra", "core", "multilib"}

			err := app.LaunchPacman(yNLabel)
			if err != nil {
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
