package files

import (
	"io"
	"os"
	"path/filepath"

	"github.com/DmitryLogunov/go-proxy-server/internal/helpers/yml"
)

// ReadOneLevelYaml reads YAML file with one level depth and return map
func ReadOneLevelYaml(relativePathToYamlFile string) (data yml.OneLevelMap, err error) {
	dataYAML, err := readTextFile(relativePathToYamlFile)
	if err != nil {
		return make(yml.OneLevelMap), err
	}

	data, err = yml.ParseOneLevelYAML(dataYAML)
	if err != nil {
		return make(yml.OneLevelMap), err
	}

	return
}

// ReadTwoLevelYaml reads YAML file with two level depth and return map
func ReadTwoLevelYaml(relativePathToYamlFile string) (data yml.TwoLevelMap, err error) {
	dataYAML, err := readTextFile(relativePathToYamlFile)
	if err != nil {
		return make(yml.TwoLevelMap), err
	}

	data, err = yml.ParseTwoLevelYAML(dataYAML)
	if err != nil {
		return make(yml.TwoLevelMap), err
	}

	return
}

// absolutePath ...
func absolutePath(relativePath string) (absolutePath string, err error) {
	dir, err := filepath.Abs(relativePath)
	if err != nil {
		return
	}

	return dir, nil
}

// readTextFile ...
func readTextFile(path string) (resultString string, err error) {
	resultString = ""

	filepath, err := absolutePath(path)
	if err != nil {
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		resultString += string(data[:n])
	}

	err = nil
	return
}
