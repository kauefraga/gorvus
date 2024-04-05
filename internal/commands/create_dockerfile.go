package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/FelipeMCassiano/gorvus/internal/templates"
	"github.com/FelipeMCassiano/gorvus/internal/utils"
	"github.com/spf13/cobra"
)

func generateDockerfile() *cobra.Command {
	var projectName string
	var language string

	cmd := &cobra.Command{
		Use:   "createDockerfile",
		Short: "Create Dockerfile based on input language and project name",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("listLanguages").Changed {
				utils.ShowSupportedLangs()
				os.Exit(0)
			}

			if len(language) == 0 {
				cmd.Help()
				fmt.Println("\n> You must specify the project name, use --language or -l")
				os.Exit(1)
			}

			if len(projectName) == 0 {
				cmd.Help()
				fmt.Println("\n> You must specify the project name, use --projectName or -p")
				os.Exit(1)
			}

			utils.VerifyIfLangIsSupported(language)

			if strings.ToLower(language) == "go" {
				if err := templates.BuildGoDockerfile(projectName); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}
			}

			if strings.ToLower(language) == "rust" {
				if err := templates.BuildRustDockerfile(projectName); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}
			}

			fmt.Println("Dockerfile created succesfully")
		},
	}

	cmd.Flags().StringVarP(&projectName, "projectName", "p", "", "Define project name")
	cmd.Flags().StringVarP(&language, "language", "l", "", "Define template language")
	cmd.Flags().BoolP("listLanguages", "s", false, "Gives a list with the supported languages")

	return cmd
}
