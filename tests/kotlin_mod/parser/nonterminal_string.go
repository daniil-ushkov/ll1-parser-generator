// Code generated by "stringer -type=NonTerminal"; DO NOT EDIT.

package parser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ret-0]
	_ = x[Sig-1]
	_ = x[Params-2]
	_ = x[Param-3]
	_ = x[Params1-4]
	_ = x[Params2-5]
}

const _NonTerminal_name = "RetSigParamsParamParams1Params2"

var _NonTerminal_index = [...]uint8{0, 3, 6, 12, 17, 24, 31}

func (i NonTerminal) String() string {
	if i < 0 || i >= NonTerminal(len(_NonTerminal_index)-1) {
		return "NonTerminal(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NonTerminal_name[_NonTerminal_index[i]:_NonTerminal_index[i+1]]
}
