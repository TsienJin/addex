package cmd

import (
	"addex/rename"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Run() {

	var dir string
	var ext string
	var ignore []string

	// Disregard
	var force bool

	rootCmd := &cobra.Command{
		Use:   "addex",
		Short: "CLI tool to add a post-fix extension to files.",
		PreRun: func(cmd *cobra.Command, args []string) {

			// Ensure required args
			if dir == "" && ext == "" {
				log.Fatalln("Required args not present!")
			}

			// Ensure that ext does not contain illegal chars
			// and ensures that it starts with "."
			pattern := `^\.[a-zA-Z0-9]+$`
			re := regexp.MustCompile(pattern)
			if !re.MatchString(ext) {
				log.Fatalln("Extension must follow this `^\\.[a-zA-Z0-9]+$` regex pattern!")
			}

			// Get absolute path of dir
			_, absDirErr := filepath.Abs(dir)
			if absDirErr != nil {
				log.Fatalln("Invalid directory path!")
			}

		},
		Run: func(cmd *cobra.Command, args []string) {

			if !force {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("You are about to add the extension '%s' to all files in '%s'. Do you want to continue? (y/N): ", ext, dir)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.ToLower(input) != "y" {
					fmt.Println("Operation canceled.")
					os.Exit(0)
				}
			}

			// Entry into main running logic
			absDir, _ := filepath.Abs(dir)
			rename.Rename(absDir, ext, ignore)
		},
	}

	rootCmd.Flags().StringVarP(&dir, "dir", "d", "", "Directory to add files to")
	rootCmd.Flags().StringVarP(&ext, "ext", "e", "", "Extension to post-fix")
	rootCmd.Flags().StringSliceVarP(&ignore, "ignore", "i", []string{}, "Ignore files with substrings as suffix")

	rootCmd.Flags().BoolP("yes", "y", false, "Skip confirmation step")

	_ = rootCmd.MarkFlagDirname("dir")
	_ = rootCmd.MarkFlagRequired("dir")
	_ = rootCmd.MarkFlagRequired("ext")

	// Execute the command to process flags and run logic
	if exErr := rootCmd.Execute(); exErr != nil {
		panic(exErr)
	}

}
