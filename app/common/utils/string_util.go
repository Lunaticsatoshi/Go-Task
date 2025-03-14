package utils

import "strconv"

func ConvertIntToString(value int) string {
	return strconv.Itoa(value)
}

func ConvertStringToUInt(value string) uint {
	uintValue, _ := strconv.ParseUint(value, 10, 32)
	return uint(uintValue)
}
