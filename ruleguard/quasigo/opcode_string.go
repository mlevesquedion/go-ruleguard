// Code generated by "stringer -type=opcode -trimprefix=op"; DO NOT EDIT.

package quasigo

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opInvalid-0]
	_ = x[opPop-1]
	_ = x[opDup-2]
	_ = x[opPushParam-3]
	_ = x[opPushLocal-4]
	_ = x[opPushFalse-5]
	_ = x[opPushTrue-6]
	_ = x[opPushConst-7]
	_ = x[opSetLocal-8]
	_ = x[opReturnTop-9]
	_ = x[opReturnFalse-10]
	_ = x[opReturnTrue-11]
	_ = x[opJump-12]
	_ = x[opJumpFalse-13]
	_ = x[opJumpTrue-14]
	_ = x[opCallBuiltin-15]
	_ = x[opIsNil-16]
	_ = x[opIsNotNil-17]
	_ = x[opNot-18]
	_ = x[opEqInt-19]
	_ = x[opNotEqInt-20]
	_ = x[opGtInt-21]
	_ = x[opGtEqInt-22]
	_ = x[opLtInt-23]
	_ = x[opLtEqInt-24]
	_ = x[opEqString-25]
	_ = x[opNotEqString-26]
	_ = x[opConcat-27]
	_ = x[opAdd-28]
	_ = x[opSub-29]
	_ = x[opStringSlice-30]
	_ = x[opStringSliceFrom-31]
	_ = x[opStringSliceTo-32]
	_ = x[opStringLen-33]
}

const _opcode_name = "InvalidPopDupPushParamPushLocalPushFalsePushTruePushConstSetLocalReturnTopReturnFalseReturnTrueJumpJumpFalseJumpTrueCallBuiltinIsNilIsNotNilNotEqIntNotEqIntGtIntGtEqIntLtIntLtEqIntEqStringNotEqStringConcatAddSubStringSliceStringSliceFromStringSliceToStringLen"

var _opcode_index = [...]uint16{0, 7, 10, 13, 22, 31, 40, 48, 57, 65, 74, 85, 95, 99, 108, 116, 127, 132, 140, 143, 148, 156, 161, 168, 173, 180, 188, 199, 205, 208, 211, 222, 237, 250, 259}

func (i opcode) String() string {
	if i >= opcode(len(_opcode_index)-1) {
		return "opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _opcode_name[_opcode_index[i]:_opcode_index[i+1]]
}