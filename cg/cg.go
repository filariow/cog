package cg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Walk ...
func Walk(outdir, baseDir string, data map[string]interface{}, extension string) error {
	return filepath.Walk(baseDir, func(elPath string, info os.FileInfo, err error) error {
		return step(outdir, elPath, baseDir, info, data, extension, err)
	})
}

func step(outdir, elPath, baseDir string, info os.FileInfo, data map[string]interface{}, extension string, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		elPath = strings.TrimPrefix(elPath, baseDir)
		return createDir(outdir, elPath, info, data)
	}
	return processFile(outdir, elPath, baseDir, info, data, extension)
}

func createDir(outdir, dirPath string, info os.FileInfo, data map[string]interface{}) error {
	fpath, err := produceOutputPath(outdir, dirPath, data)
	if err != nil {
		return err
	}
	return os.MkdirAll(*fpath, info.Mode().Perm())
}

func processFile(outdir, filePath, baseDir string, info os.FileInfo, data map[string]interface{}, extension string) error {
	fp2 := strings.TrimPrefix(filePath, baseDir)
	fp2 = strings.TrimLeft(fp2, "/")
	fpath, err := produceOutputPath(outdir, fp2, data)
	if err != nil {
		return err
	}

	readFile, err := os.Open(filePath)
	if err != nil {
		return err
	}

	fcontent, err := ioutil.ReadAll(readFile)
	if err != nil {
		return err
	}

	fcontentStr := string(fcontent)
	ts := fmt.Sprintf(".%s", extension)
	if !strings.HasSuffix(*fpath, ts) {
		return ioutil.WriteFile(*fpath, []byte(fcontentStr), info.Mode().Perm())
	}

	tfContent, err := applyTemplateToContent(&fcontentStr, data)
	if err != nil {
		return err
	}

	tfpath := strings.TrimSuffix(*fpath, ts)
	return ioutil.WriteFile(tfpath, []byte(*tfContent), info.Mode().Perm())
}

func produceOutputPath(outdir, dirPath string, data map[string]interface{}) (*string, error) {
	procPath, err := applyTemplateToPath(&dirPath, data)
	if err != nil {
		return nil, err
	}
	joinedPath := path.Join(outdir, *procPath)
	return &joinedPath, nil
}

func applyTemplateToPath(path *string, data map[string]interface{}) (*string, error) {
	t := template.New("Path").Funcs(fmap)
	return applyTemplate(t, path, data)
}

func applyTemplateToContent(content *string, data map[string]interface{}) (*string, error) {
	t := template.New("FileContent").Funcs(fmap)
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

// ReadConfig reads the config file and returns a map
func ReadConfig(path string) (map[string]interface{}, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var configMap map[string]interface{}
	if err := yaml.Unmarshal(configData, &configMap); err != nil {
		return nil, err
	}
	return configMap, nil
}
