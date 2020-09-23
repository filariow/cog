package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	outdir         = "./out/"
	configFilePath = "./config.yaml"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func run() error {
	data, err := readConfig(configFilePath)
	if err != nil {
		return err
	}

	return walk("./tmpl/", data)
}

func walk(baseDir string, data map[string]interface{}) error {
	return filepath.Walk(baseDir, func(elPath string, info os.FileInfo, err error) error {
		return step(elPath, info, data, err)
	})
}

func step(elPath string, info os.FileInfo, data map[string]interface{}, err error) error {
	fmt.Println(elPath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return createDir(elPath, info, data)
	}
	return processFile(elPath, info, data)
}

func createDir(dirPath string, info os.FileInfo, data map[string]interface{}) error {
	fpath, err := finalPath(dirPath, data)
	if err != nil {
		return err
	}
	return os.MkdirAll(*fpath, info.Mode().Perm())
}

func processFile(filePath string, info os.FileInfo, data map[string]interface{}) error {
	fpath, err := finalPath(filePath, data)
	if err != nil {
		return err
	}

	readFile, err := os.Open(filePath)
	if err != nil {
		return err
	}

	fcontent, err := ioutil.ReadAll(readFile)
	fcontentStr := string(fcontent)
	tfContent, err := applyTemplateToContent(&fcontentStr, data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(*fpath, []byte(*tfContent), info.Mode().Perm())
}

func finalPath(elPath string, data map[string]interface{}) (*string, error) {
	procPath, err := applyTemplateToPath(&elPath, data)
	if err != nil {
		return nil, err
	}
	joinedPath := path.Join(outdir, *procPath)
	return &joinedPath, nil
}

func applyTemplateToPath(path *string, data map[string]interface{}) (*string, error) {
	t := template.New("Path")
	return applyTemplate(t, path, data)
}

func applyTemplateToContent(content *string, data map[string]interface{}) (*string, error) {
	t := template.New("FileContent")
	return applyTemplate(t, content, data)
}

func applyTemplate(t *template.Template, value *string, data map[string]interface{}) (*string, error) {
	t, err := t.Parse(*value)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return nil, err
	}
	content := buf.String()
	return &content, nil
}

func readConfig(path string) (map[string]interface{}, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var configDS map[string]interface{}
	if err := yaml.Unmarshal(configData, &configDS); err != nil {
		return nil, err
	}
	return configDS, nil
}
