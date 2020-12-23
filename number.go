// Copyright 2020  enmotech.chengdu. All rights reserved.
// Use of this utils to operate something with number

package utils

//GetParamRegionInt  int
//get param when in a GetParamRegion threshold region ,
//when the param is bigger than threshold bigger one or smaller than threshold smaller one
//set param on threshold
//param need check
//defaultValue when param is zero,set param = defaultValue
//max,min : threshold
func GetParamRegionInt(param, defaultValue, max, min int) int {
	if param == 0 {
		return defaultValue
	}
	if param > max {
		return max
	}
	if param < min {
		return min
	}
	return param
}

//GetParamRegionUint  uint
func GetParamRegionUint(param, defaultValue, max, min uint) uint {

	return uint(GetParamRegionInt(int(param), int(defaultValue), int(max), int(min)))
}

//GetParamRegionUint64  uint64
func GetParamRegionUint64(param, defaultValue, max, min uint64) uint64 {

	return uint64(GetParamRegionInt(int(param), int(defaultValue), int(max), int(min)))
}
