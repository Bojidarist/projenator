package projenator

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"projenator/templates"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(electronWebAppCmd)

	electronWebAppCmd.PersistentFlags().Bool("npmi", false, "Run npm install after the template is created")
	electronWebAppCmd.PersistentFlags().Bool("yarni", false, "Run yarn install after the template is created")
	electronWebAppCmd.PersistentFlags().Int("width", 800, "Defines the width of the browser window")
	electronWebAppCmd.PersistentFlags().Int("height", 600, "Defines the height of the browser window")
}

var electronWebAppCmd = &cobra.Command{
	Use:   "electron-web-app name url",
	Short: "Generate electron web app from a url",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		npmInstallFlag, _ := cmd.Flags().GetBool("npmi")
		yarnInstallFlag, _ := cmd.Flags().GetBool("yarni")
		widthFlag, _ := cmd.Flags().GetInt("width")
		heightFlag, _ := cmd.Flags().GetInt("height")

		name := args[0]
		url := args[1]
		dirPath := filepath.Join(".", name)

		if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(dirPath, os.ModePerm)
			if err != nil {
				return err
			}

			tmplObj := map[string]string{
				"ProjectName":  name,
				"WebsiteURL":   url,
				"WindowWidth":  strconv.Itoa(widthFlag),
				"WindowHeight": strconv.Itoa(heightFlag),
			}

			err = writeTemplates(tmplObj)
			if err != nil {
				return err
			}

			if npmInstallFlag {
				fmt.Println("Running npm install")

				cmd := exec.Command("npm", "install")
				cmd.Dir = dirPath
				err = cmd.Run()
				if err != nil {
					return err
				}
			} else if yarnInstallFlag {
				fmt.Println("Running yarn install")

				cmd := exec.Command("yarn", "install")
				cmd.Dir = dirPath
				err = cmd.Run()
				if err != nil {
					return err
				}
			}

		} else {
			cmd.PrintErrln("A file or folder with the name", name, "already exists!")
			return os.ErrExist
		}

		return nil
	},
}

func writeTemplates(tmplObj map[string]string) error {
	tmplFiles, err := fs.ReadDir(templates.TemplatesFS, "electron-web-app")
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(templates.TemplatesFS, path.Join("electron-web-app", tmpl.Name()))
		if err != nil {
			return err
		}

		f, err := os.Create(path.Join(tmplObj["ProjectName"], strings.ReplaceAll(pt.Name(), ".tmpl", "")))
		if err != nil {
			return err
		}

		err = pt.Execute(f, tmplObj)
		if err != nil {
			return err
		}

		f.Close()
	}

	return nil
}
