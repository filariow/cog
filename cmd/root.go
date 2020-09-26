package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/FrancescoIlario/gocg/cg"
	"github.com/spf13/cobra"
)

const (
	configFileDefault     = "gocg.yaml"
	contextDefault        = "."
	outputFolderDefault   = "./out/"
	templateFolderDefault = "templates"
)

var (
	context            string
	outputFolder       string
	configFilePath     string
	templateFolderName string
	templateContext    string
)

var rootCmd = &cobra.Command{
	Use:   "gocg [flags] PATH",
	Short: "gocg go-code-generation",
	Args:  cobra.ExactValidArgs(1),
	RunE:  rootCmdRunE,
}

func parseInputs(cmd *cobra.Command, args []string) error {
	var err error
	if context, err = parseContext(args); err != nil {
		return err
	}

	if templateContext, err = parseTemplateContext(context); err != nil {
		return err
	}

	if configFilePath, err = parseConfigFilePath(context); err != nil {
		return err
	}
	return nil
}

func parseContext(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no context provided")
	}

	ctx := args[0]
	if !existsDir(ctx) {
		return "", fmt.Errorf("context folder not found: %s", ctx)
	}
	return ctx, nil
}

func parseTemplateContext(basePath string) (string, error) {
	tmpl := path.Join(basePath, templateFolderName)
	if !existsDir(tmpl) {
		return "", fmt.Errorf("template folder not found: %s", tmpl)
	}
	return tmpl, nil
}

func parseConfigFilePath(basePath string) (string, error) {
	cfp := configFilePath
	if cfp == "" {
		cfp = path.Join(basePath, configFileDefault)
	}

	info, err := os.Stat(cfp)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", fmt.Errorf("expected a file for config file, obtained a directory: %s", configFilePath)
	}
	return cfp, err
}

func rootCmdRunE(cmd *cobra.Command, args []string) error {
	if err := parseInputs(cmd, args); err != nil {
		return err
	}
	outputFolder = strings.TrimPrefix(outputFolder, "./")
	templateContext = strings.TrimPrefix(templateContext, "./")
	return processTemplates(outputFolder, configFilePath, templateContext)
}

func processTemplates(outdir, config, templates string) error {
	configData, err := cg.ReadConfig(config)
	if err != nil {
		return fmt.Errorf("error reading configuration %s", config)
	}

	if err := cg.Walk(outdir, templates, configData); err != nil {
		return fmt.Errorf("error walking into template dir %s: %v", templates, err)
	}
	return nil
}

func existsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFolder, "outdir", "o", outputFolderDefault, "The path where to store the generated code")
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "", "The path to the config file to use (default gocg.yaml in context folder)")
	rootCmd.PersistentFlags().StringVarP(&templateFolderName, "template", "t", templateFolderDefault, "The template folder in the context one")

	rootCmd.MarkFlagRequired("outdir")
}

// Execute ...
func Execute() error {
	return rootCmd.Execute()
}
