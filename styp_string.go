// Code generated by "stringer -type=styp"; DO NOT EDIT

package main

import "fmt"

const _styp_name = "matchsplitsingle"

var _styp_index = [...]uint8{0, 5, 10, 16}

func (i styp) String() string {
	if i < 0 || i >= styp(len(_styp_index)-1) {
		return fmt.Sprintf("styp(%d)", i)
	}
	return _styp_name[_styp_index[i]:_styp_index[i+1]]
}
