package utils

import (
	"strconv"
	"strings"
)

// SplitAndParseUint64 Split space string and Parse uint64
func SplitAndParseUint64(s string, sep string) uint64 {
	sp := strings.Split(s, sep)
	r, _ := strconv.ParseUint(sp[0], 10, 64)

	return r
}

// SplitAndRangeParseInt64 Split use sep param And ParseInt
func SplitAndRangeParseInt64(s string, sep string) []int64 {
	var r []int64

	sp := strings.Split(s, sep)
	for _, v := range sp {
		vv, _ := strconv.ParseInt(v, 10, 64)
		r = append(r, vv)
	}

	return r
}
