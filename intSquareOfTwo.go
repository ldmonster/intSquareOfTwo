package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func SliceOfBytes(intValue int) bool {
	if intValue < 0 {
		intValue = -intValue
	}
	a := int64(intValue)
	if a == 1 {
		return false
	}
	byteSliceRev := *(*[8]byte)(unsafe.Pointer(&a))
	byteSlice := make([]byte, 8)
	firstByteFound := -1
	for i := 0; i < 8; i++ {
		byteSlice[i] = byteSliceRev[7-i]
		if byteSlice[i] > 128 {
			return false
		}
		if byteSlice[i] != 0 {
			if firstByteFound != -1 {
				return false
			}
			firstByteFound = i
		}
	}
	if firstByteFound == -1 {
		return false
	}
	binValue := fmt.Sprintf("%b", byteSlice[firstByteFound])
	firstBitFound := false
	for i := 0; i < len(binValue); i++ {
		if binValue[i:i+1] == "1" {
			if firstBitFound {
				return false
			}
			firstBitFound = true
		}
	}
	return firstBitFound
}

func StrConvToBytes(intValue int) bool {
	if intValue == 1 {
		return false
	}
	binValue := strconv.FormatInt(int64(intValue), 2)
	firstBitFound := false
	for i := 0; i < len(binValue); i++ {
		if binValue[i:i+1] == "1" {
			if firstBitFound {
				return false
			}
			firstBitFound = true
		}
	}
	return firstBitFound
}

func main() {
	//fmt.Println(SliceOfBytes(256))

}
