package utils

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	GetStatusError = errors.New("cannot send get request")
)

func GetStatusResponse(url string) (string, error) {
	const op = "GetStatusResponse()"
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("%s: %v", op, err)
		return "", GetStatusError
	}
	logrus.Infof("%s to %s with response %s", op, url, resp.Status)
	return resp.Status, nil
}
