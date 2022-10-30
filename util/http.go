package util

import (
	"io"
	"io/ioutil"
)

func GetStrRespBody(respBody io.ReadCloser) (strRespBody string, err error) {
	defer respBody.Close()

	tmpRespBodyBites, err := ioutil.ReadAll(respBody)
	if err != nil {
		return "", err
	}

	return string(tmpRespBodyBites), nil
}
