package main

import (
	"fmt"
	"os"
	"task2/config"
	"task2/internal/utils"
	"task2/pkg/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()
	logger.SetLogrus(cfg.Log)

	arguments := os.Args
	if len(arguments) < 4 {
		fmt.Println("incorrect arguments")
		os.Exit(1)
	}

	logsOutput := arguments[1]

	switch logsOutput {
	case "file":
		f, _ := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		logrus.SetOutput(f)
	case "stdout":
		logrus.SetOutput(os.Stdout)
	default:
		fmt.Println("incorrect logger output")
		os.Exit(1)
	}

	dataFrom := arguments[2]

	var nums []int
	var sum int

	switch dataFrom {
	case "file":
		var err error
		nums, err = utils.GetArrFromJson(cfg.Storage)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum = utils.TotalSum(nums)
	case "stdin":
		var len int

		fmt.Println("Enter array length")
		fmt.Fscan(os.Stdin, &len)
		for i := 0; i < len; i++ {
			var num int
			fmt.Printf("Enter %d number: \n", i+1)
			fmt.Fscan(os.Stdin, &num)
			nums = append(nums, num)
		}
		sum = utils.TotalSum(nums)
	default:
		fmt.Println("incorrect data input")
		os.Exit(1)
	}

	status, err := utils.GetStatusResponse(cfg.URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataTo := arguments[3]

	switch dataTo {
	case "file":
		f, _ := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		defer f.Close()
		f.WriteString(fmt.Sprintf("Array - %v, Sum - %d, Response Status to %s - %s", nums, sum, cfg.URL, status))
	case "stdout":
		fmt.Println(fmt.Sprintf("Array - %v, Sum - %d, Response Status to %s - %s", nums, sum, cfg.URL, status))
	default:
		fmt.Println("incorrect data output")
		os.Exit(1)
	}

}
