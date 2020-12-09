package utils

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// Parse Int else return zero
func ParseInt(s string) int {
	value := 0
	if s != "" {
		val, err := strconv.ParseInt(s, 10, 32)
		if err == nil {
			value = int(val)
		} else {
			logrus.Errorln(err)
		}
	}
	return value
}

// Parse date else return current timestamp
func ParseInt64(s string) int64 {
	value := int64(0)
	if s != "" {
		val, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			value = val
		} else {
			logrus.Errorln(err)
		}
	}
	return value
}

// Parse date else return current timestamp
func ParseStringArray(s string) []string {
	arr := []string{}
	if s != "" {
		arr = strings.Split(s, ",")
	}
	return arr
}

func ParseInt32Array(s string) ([]int32, error) {
	arr := []int32{}
	if s != "" {
		for _, str := range strings.Split(s, ",") {
			i, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			arr = append(arr, int32(i))
		}
	}
	return arr, nil
}

func ParseInt64Array(s string) ([]int64, error) {
	arr := []int64{}
	if s != "" {
		for _, str := range strings.Split(s, ",") {
			i, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			arr = append(arr, int64(i))
		}
	}
	return arr, nil
}
