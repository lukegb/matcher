// Code generated by "stringer -type=charType"; DO NOT EDIT

package main

import "fmt"

const _charType_name = "charNilcharEscapeLiteralcharLiteralcharDotcharBegincharEndcharStarcharConcat"

var _charType_index = [...]uint8{0, 7, 24, 35, 42, 51, 58, 66, 76}

func (i charType) String() string {
	if i < 0 || i >= charType(len(_charType_index)-1) {
		return fmt.Sprintf("charType(%d)", i)
	}
	return _charType_name[_charType_index[i]:_charType_index[i+1]]
}
