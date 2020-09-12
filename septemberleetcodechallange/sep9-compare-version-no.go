package septemberleetcodechallange

import (
	"strconv"
	"strings"
)

func strToNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func compareVersion(version1 string, version2 string) int {
	numsInV1 := strings.Split(version1, ".")
	numsInV2 := strings.Split(version2, ".")

	if len(numsInV1) >= len(numsInV2) {
		for i, numStr := range numsInV1 {
			if i < len(numsInV2) {
				if strToNum(numStr) > strToNum(numsInV2[i]) {
					return 1
				} else if strToNum(numStr) < strToNum(numsInV2[i]) {
					return -1
				}
			} else {
				if strToNum(numStr) > 0 {
					return 1
				}
			}
		}
	} else if len(numsInV2) > len(numsInV1) {
		for i, numStr := range numsInV2 {
			if i < len(numsInV1) {
				if strToNum(numStr) > strToNum(numsInV1[i]) {
					return -1
				} else if strToNum(numStr) < strToNum(numsInV1[i]) {
					return 1
				}
			} else {
				if strToNum(numStr) > 0 {
					return -1
				}
			}
		}
	}
	return 0
}
