package utils

import "github.com/sirupsen/logrus"

func TotalSum(nums []int) int {
	const op = "TotalSum()"
	var sum int

	for _, v := range nums {
		sum += v
	}
	logrus.Infof("%s: input - %v, result - %d", op, nums, sum)
	return sum
}
