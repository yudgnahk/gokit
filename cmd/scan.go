package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yudgnahk/gokit/constants"
	"gorm.io/gorm/utils"
)

type File struct {
	Name      string
	Extension string
}

type Dir struct {
	Name  string
	Path  string
	Dirs  []Dir
	Files []File
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the structure of the project",
	Run: func(cmd *cobra.Command, args []string) {
		// get source path
		path, _ := cmd.Flags().GetString("source")
		rootDir := Scan(path)

		// print the structure
		printDir(rootDir, 0)
	},
}

// Scan I want to scan the structure of the project
func Scan(path string) Dir {
	// use recursive to get all the folders
	if len(path) == 0 {
		path, _ = os.Getwd()
	}

	rootDir := getDir(path)
	return rootDir
}

func getDir(path string) Dir {
	ignoreComponents := []string{".DS_Store", ".git", ".gitignore", ".idea"}

	res := Dir{
		Name:  filepath.Base(path),
		Path:  path,
		Dirs:  []Dir{},
		Files: []File{},
	}

	// get all children folders
	children, _ := os.ReadDir(path)
	for _, child := range children {
		if child.IsDir() {
			if !utils.Contains(ignoreComponents, child.Name()) {
				res.Dirs = append(res.Dirs, Dir{
					Name: child.Name(),
					Path: filepath.Join(path, child.Name()),
				})
			}
		} else {
			if !utils.Contains(ignoreComponents, child.Name()) {
				//if filepath.Ext(child.Name()) == ".go" {
				res.Files = append(res.Files, File{
					Name:      child.Name(),
					Extension: child.Type().String(),
				})
				//}
			}
		}
	}

	for i := range res.Dirs {
		res.Dirs[i] = getDir(res.Dirs[i].Path)
	}

	return res
}

func getTab(level int) string {
	res := ""
	for i := 0; i < level; i++ {
		res += "───"
	}

	return res
}

func printDir(dir Dir, level int) {
	fmt.Println(constants.ColorRed, "├──"+getTab(level), dir.Name)
	level++
	for _, child := range dir.Files {
		fmt.Println(constants.ColorGreen, "├──"+getTab(level), child.Name)
	}

	for _, child := range dir.Dirs {
		printDir(child, level)
	}
}

func init() {
	RootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringP("source", "s", "", "Source path")
}
