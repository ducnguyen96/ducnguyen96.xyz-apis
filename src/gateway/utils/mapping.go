package utils

import "strconv"

func Uint64ToString(number uint64) string {
	s := strconv.FormatUint(number, 10)
	return s
}

func Float32ToFloat64(number float32) float64 {
	return float64(number)
}

func Int32ToInt(number int32) int {
	return int(number)
}

func OptionalInt32ToOptionalInt(number *int32) *int {
	if number == nil {
		return nil
	}
	num := int(*number)
	return &num

}
