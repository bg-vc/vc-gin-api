package util

import (
	"strconv"
)

func ConvertStringToBool(boolString string, defaultValue bool) bool {
	retBool := defaultValue
	if len(boolString) > 0 {
		b, err := strconv.ParseBool(boolString)
		if err == nil {
			retBool = b
		}
	}
	return retBool
}

func ConvertStringToFloat(val string, defaultValue float64) float64 {
	ret := defaultValue
	if len(val) > 0 {
		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			ret = f
		}
	}
	return ret
}

func ConvertStringToInt(intString string, defaultValue int) int {
	retInt := defaultValue
	if len(intString) > 0 {
		i, err := strconv.Atoi(intString)
		if err == nil {
			retInt = i
		}
	}
	return retInt
}

func ConvertStringToInt64(intString string, defaultValue int64) int64 {
	retInt64 := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseInt(intString, 10, 64)
		if err == nil {
			retInt64 = i64
		}
	}
	return retInt64
}
