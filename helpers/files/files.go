package files

import (
	"io"
	"os"
	"path/filepath"
	"github.com/gopkg.in/yaml"
)

// ReadOneLevelYaml reads YAML file with one level depth and return map 
func ReadOneLevelYaml(relativePathToYamlFile string) (data map[string]string, err error) {
	dataYAML, err := readTextFile(relativePathToYamlFile)
	if err != nil {
		return make(map[string]string), err
	}

	data, err = parseOneLevelYAML(dataYAML)
	if err != nil {
		return make(map[string]string), err
	}

	return
}

// parseOneLevelYAML ...
func parseOneLevelYAML(data string) (parsedData map[string]string, err error) {
	err = yaml.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		return make(map[string]string), err
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
