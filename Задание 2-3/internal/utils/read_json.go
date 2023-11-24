package utils

import (
	"encoding/json"
	"errors"
	"os"
	"task2/internal/entity"

	"github.com/sirupsen/logrus"
)

var (
	ErrReadJSON  = errors.New("cannot read json file")
	ErrUnmarshal = errors.New("cannot unmarshal")
)

func GetArrFromJson(filePath string) ([]int, error) {
	const op = "GetArrFromJson()"
	var numsArr []int
	var blocks entity.NumberBlocks
	f, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Errorf("%s: %v", op, err)
		return nil, ErrReadJSON
	}

	err = json.Unmarshal(f, &blocks)
	if err != nil {
		logrus.Errorf("%s: %v", op, err)
		return nil, ErrUnmarshal
	}
	for _, v := range blocks.Blocks {
		numsArr = append(numsArr, v.Num)
	}
	logrus.Infof("%s: data - %v", op, numsArr)
	return numsArr, nil

}
