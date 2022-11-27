package yml

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

type OneLevelMap map[string]string

func ParseOneLevelYAML(data string) (parsedData map[string]string, err error) {
	err = yaml.Unmarshal([]byte(data), &parsedData)
	if checkAndReturnIfError(err) {
		return OneLevelMap{}, err
	}

	return
}

type TwoLevelMap map[string]map[string]string

func ParseTwoLevelYAML(data string) (parsedData map[string]map[string]string, err error) {
	err = yaml.Unmarshal([]byte(data), &parsedData)
	if checkAndReturnIfError(err) {
		return TwoLevelMap{}, err
	}

	return
}

type ThreeLevelMap map[string]map[string]map[string]string

func ParseThreeLevelYAML(data string) (parsedData map[string]map[string]map[string]string, err error) {
	err = yaml.Unmarshal([]byte(data), &parsedData)
	if checkAndReturnIfError(err) {
		return ThreeLevelMap{}, err
	}

	return
}

type FourLevelMap map[string]map[string]map[string]map[string]string

func ParseFourLevelYAML(data string) (parsedData FourLevelMap, err error) {
	err = yaml.Unmarshal([]byte(data), &parsedData)
	if checkAndReturnIfError(err) {
		return FourLevelMap{}, err
	}

	return
}

func checkAndReturnIfError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}

	return false
}
