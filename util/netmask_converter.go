package util

import (
	"math"
	"strconv"
	"strings"
)

//convert string to in
//such as "255.255.255.0" -> 24
func Converter(netmask_s string) int {
	var err bool
	var int4, int3, int2, int1 int
	dict := map[string]int{"255": 8, "254": 7, "252": 6, "248": 5, "240": 4, "224": 3, "192": 2, "128": 1, "0": 0}
	idx1 := strings.Index(netmask_s, ".")
	if idx1 < 0 {
		return -1
	}
	idx2 := strings.Index(netmask_s[idx1+1:], ".")
	if idx2 < 0 {
		return -1
	}
	idx3 := strings.Index(netmask_s[idx1+1+idx2+1:], ".")
	if idx3 < 0 {
		return -1
	}
	// fmt.Println(netmask_s,len(netmask_s),idx3,idx2,idx1)
	s1 := netmask_s[0:idx1]
	s2 := netmask_s[idx1+1 : idx2+idx1+1]
	s3 := netmask_s[idx2+idx1+2 : idx3+idx2+idx1+2]
	s4 := netmask_s[idx3+idx2+idx1+3:]
	//fmt.Println(idx1,idx2,idx3)
	//fmt.Println(s1,s2,s3,s4)
	int4, err = dict[s4]
	if !err {
		return -1
	}
	int3, err = dict[s3]
	if !err {
		return -1
	}
	int2, err = dict[s2]
	if !err {
		return -1
	}
	int1, err = dict[s1]
	if !err {
		return -1
	}
	//fmt.Println(int1,int2,int3,int4)
	if int1 == 0 {
		if int2 > 0 || int3 > 0 || int4 > 0 {
			return -1
		} else {
			return 0
		}
	} else if int2 == 0 {
		if int3 > 0 || int4 > 0 {
			return -1
		} else {
			return int1
		}
	} else if int3 == 0 {
		if int4 > 0 {
			return -1
		}
		if int1 != 8 {
			return -1
		}
		return int1 + int2
	} else if int4 == 0 {
		if int1 != 8 || int2 != 8 {
			return -1
		} else {
			return int1 + int2 + int3
		}
	} else {
		if int1 != 8 || int2 != 8 || int3 != 8 {
			return -1
		} else {
			return int1 + int2 + int3 + int4
		}
	}
}

var TenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

/*进制转换的键值查询*/
func findkey(in string) int {
	result := -1
	for k, v := range TenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func anyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findkey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
func NetMaskSysConvert(s string) string { //子网掩码的格式类型转换
	var tempNetMask string
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[:2], 16)) + "."
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[2:4], 16)) + "."
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[4:6], 16)) + "."
	tempNetMask = tempNetMask + strconv.Itoa(anyToDecimal(s[6:], 16))
	return tempNetMask
}
