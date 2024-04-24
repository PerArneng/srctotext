package main

import (
	"fmt"
	"io/ioutil"

	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type customFormatter struct{}

// Format building the log message string to your specific format.
func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04")
	return []byte(fmt.Sprintf("%s %s %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&customFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	initLogger()

	var rootCmd = &cobra.Command{
		Use:   "srctotext",
		Short: "srctotext is a tool to aggregate source files into a single text file.",
		Run:   run,
	}

	var path, include, output string
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "Specify the root folder to search.")
	rootCmd.Flags().StringVarP(&include, "include", "i", "", "File patterns to include (comma-separated).")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Output file where content will be written.")
	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("include")
	rootCmd.MarkFlagRequired("output")

	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Command execution failed")
	}
}

func run(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
	include, _ := cmd.Flags().GetString("include")
	output, _ := cmd.Flags().GetString("output")

	// Verify the path is a directory
	if stat, err := os.Stat(path); err != nil || !stat.IsDir() {
		log.WithField("path", path).Error("Invalid directory path")
		os.Exit(1)
	}

	// Open output file
	file, err := os.Create(output)
	if err != nil {
		log.WithError(err).Fatal("Failed to create output file")
	}
	defer file.Close()

	// Process files
	patterns := strings.Split(include, ",")
	foundFiles := false
	for _, pattern := range patterns {
		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				log.WithError(err).WithField("file", filePath).Error("Error walking the file path")
				return err
			}
			if !info.IsDir() && matchesPattern(filePath, pattern) {
				if isBinary(filePath) {
					return nil
				}
				if writeToFile(file, path, filePath) {
					foundFiles = true
				}
			}
			return nil
		})
		if err != nil {
			log.WithError(err).WithField("pattern", pattern).Error("Failed to process files")
			os.Exit(1)
		}
	}

	if !foundFiles {
		log.Error("No files matched the specified patterns")
		os.Exit(1)
	}

	log.Info("Files successfully written to output")
}

func matchesPattern(filePath string, pattern string) bool {
	matched, err := filepath.Match(pattern, filepath.Base(filePath))
	if err != nil {
		log.WithError(err).WithField("pattern", pattern).Error("Error matching file pattern")
	}
	return matched
}

func isBinary(filePath string) bool {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.WithError(err).WithField("file", filePath).Error("Failed to read file for binary check")
		return false
	}
	for _, b := range data {
		if b == 0 {
			return true
		}
	}
	return false
}

func writeToFile(file *os.File, baseDir, filePath string) bool {

	// Calculate the relative path from the base directory to the file path
	relativePath, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		log.WithError(err).WithField("file", filePath).Error("Failed to get relative path")
		return false
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.WithError(err).WithField("file", filePath).Error("Failed to read file")
		return false
	}

	_, err = file.WriteString(fmt.Sprintf("# FILE: %s\n\n%s\n\n", relativePath, data))
	if err != nil {
		log.WithError(err).WithField("file", filePath).Error("Failed to write file content to output")
		return false
	}
	return true
}
