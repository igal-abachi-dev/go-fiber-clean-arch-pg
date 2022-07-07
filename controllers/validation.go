package controllers

import "strconv"

func isInt(param string) (bool, int32) {
	value, err := strconv.ParseInt(param, 0, 64)
	return err == nil, int32(value)
}
