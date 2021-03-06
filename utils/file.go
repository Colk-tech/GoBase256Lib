package utils

import (
	"io/ioutil"
)

func ReadFile(fileName string) (result string, err error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return result, err
	}

	result = string(bytes)

	return result, err
}
