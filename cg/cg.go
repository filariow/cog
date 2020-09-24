package cg

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

func Walk(outdir, baseDir string, data map[string]interface{}) error {
	return filepath.Walk(baseDir, func(elPath string, info os.FileInfo, err error) error {
		return step(outdir, elPath, info, data, err)
	})
}

func step(outdir, elPath string, info os.FileInfo, data map[string]interface{}, err error) error {
	fmt.Println(elPath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return createDir(outdir, elPath, info, data)
	}
	return processFile(outdir, elPath, info, data)
}

func createDir(outdir, dirPath string, info os.FileInfo, data map[string]interface{}) error {
	fpath, err := produceOutputPath(outdir, dirPath, data)
	if err != nil {
		return err
	}
	return os.MkdirAll(*fpath, info.Mode().Perm())
}

func processFile(outdir, filePath string, info os.FileInfo, data map[string]interface{}) error {
	fpath, err := produceOutputPath(outdir, filePath, data)
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

func produceOutputPath(outdir, elPath string, data map[string]interface{}) (*string, error) {
	procPath, err := applyTemplateToPath(&elPath, data)
	if err != nil {
		return nil, err
	}
	joinedPath := path.Join(outdir, *procPath)
	return &joinedPath, nil
}

func applyTemplateToPath(path *string, data map[string]interface{}) (*string, error) {
	t := template.New("Path").Funcs(FMap)
	return applyTemplate(t, path, data)
}

func applyTemplateToContent(content *string, data map[string]interface{}) (*string, error) {
	t := template.New("FileContent").Funcs(FMap)
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

func ReadConfig(path string) (map[string]interface{}, error) {
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
