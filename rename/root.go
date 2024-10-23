package rename

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func renameFile(
	filePath string,
	ext string,
	ignore []string,
) error {

	// Get the original file name
	fileName := filepath.Base(filePath)

	// Check if file should be ignored
	ignore = append(ignore, ext)
	for _, suffix := range ignore {
		if strings.HasSuffix(fileName, suffix) {
			fmt.Printf("Skipping:\t%s\n", filePath)
			return nil
		}
	}

	// File should be renamed
	fmt.Printf("Renaming:\t%s\n", filePath)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("rename", filePath, filePath+ext)
	} else {
		cmd = exec.Command("mv", filePath, filePath+ext)
	}

	err := cmd.Run()
	return err

}

// Rename function to walk through the directory and rename files to append the extension
//
// Arguments:
//   - dir: string, absolute directory path to enter
//   - ext: string, extension to append to the files
//   - ignore: []string, ignore files with substrings
func Rename(
	dir string,
	ext string,
	ignore []string,
) {

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			err := renameFile(path, ext, ignore)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}
