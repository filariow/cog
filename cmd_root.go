package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	outDirDefault         = "./out/"
	configFilePathDefault = "./config.yaml"
)

var (
	outdir         string
	tmpldir        string
	configFilePath string

	templateDirs = []string{"./tmpl", "./template"}
)

var rootCmd = &cobra.Command{
	Use:   "gocg",
	Short: "gocg go-code-generation",
	Args:  rootCmdParseArgs,
	Run:   rootCmdRun,
}

func rootCmdParseArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 && !existsDefaultTemplateDir() {
		return fmt.Errorf("template directory not found")
	}

	tmpldir = args[0]
	if !existsDir(tmpldir) {
		return fmt.Errorf("template directory not found: %s", tmpldir)
	}
	return nil
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	if err := processTemplates(configFilePath, tmpldir); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

func processTemplates(config, templates string) error {
	data, err := readConfig(config)
	if err != nil {
		return fmt.Errorf("error reading configuration %s", config)
	}

	if err := walk(templates, data); err != nil {
		return fmt.Errorf("error walking into template dir %s", templates)
	}
	return nil
}

func existsDefaultTemplateDir() bool {
	return existsOneDir(templateDirs...)
}

func existsOneDir(paths ...string) bool {
	for _, path := range paths {
		if d := existsDir(path); d {
			return true
		}
	}
	return false
}

func existsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("can't read path \"%s\": %s", path, err)
		return false
	}
	return info.IsDir()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outdir, "outdir", "o", outDirDefault, "The path where to store the generated code")
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", configFilePath, "The config file to use")
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
