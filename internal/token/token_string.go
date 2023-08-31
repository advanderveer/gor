// Code generated by "stringer -type Token -linecomment"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ILLEGAL-0]
	_ = x[EOF-1]
	_ = x[IDENT-2]
	_ = x[COMMENT-3]
	_ = x[BREAK-4]
	_ = x[CASE-5]
	_ = x[CHAN-6]
	_ = x[CONST-7]
	_ = x[CONTINUE-8]
	_ = x[DEFAULT-9]
	_ = x[DEFER-10]
	_ = x[ELSE-11]
	_ = x[FALLTHROUGH-12]
	_ = x[FOR-13]
	_ = x[FUNC-14]
	_ = x[GO-15]
	_ = x[GOTO-16]
	_ = x[IF-17]
	_ = x[IMPORT-18]
	_ = x[INTERFACE-19]
	_ = x[MAP-20]
	_ = x[PACKAGE-21]
	_ = x[RANGE-22]
	_ = x[RETURN-23]
	_ = x[SELECT-24]
	_ = x[STRUCT-25]
	_ = x[SWITCH-26]
	_ = x[TYPE-27]
	_ = x[VAR-28]
	_ = x[literal_beg-29]
	_ = x[INT-30]
	_ = x[FLOAT-31]
	_ = x[STRING-32]
	_ = x[literal_end-33]
	_ = x[SEMICOLON-34]
	_ = x[COLON-35]
	_ = x[PERIOD-36]
	_ = x[ELLIPSIS-37]
	_ = x[COMMA-38]
	_ = x[DEFINE-39]
	_ = x[ARROW-40]
	_ = x[INC-41]
	_ = x[DEC-42]
	_ = x[ASSIGN-43]
	_ = x[EQL-44]
	_ = x[NEQ-45]
	_ = x[NOT-46]
	_ = x[AND_NOT-47]
	_ = x[AND-48]
	_ = x[XOR-49]
	_ = x[OR-50]
	_ = x[LAND-51]
	_ = x[LOR-52]
	_ = x[LPAREN-53]
	_ = x[RPAREN-54]
	_ = x[LBRACK-55]
	_ = x[RBRACK-56]
	_ = x[LBRACE-57]
	_ = x[RBRACE-58]
	_ = x[LSS-59]
	_ = x[LEQ-60]
	_ = x[SHL-61]
	_ = x[GTR-62]
	_ = x[GEQ-63]
	_ = x[SHR-64]
	_ = x[QUO-65]
	_ = x[MUL-66]
	_ = x[ADD-67]
	_ = x[SUB-68]
	_ = x[REM-69]
	_ = x[QUO_ASSIGN-70]
	_ = x[MUL_ASSIGN-71]
	_ = x[ADD_ASSIGN-72]
	_ = x[SUB_ASSIGN-73]
	_ = x[AND_NOT_ASSIGN-74]
	_ = x[AND_ASSIGN-75]
	_ = x[OR_ASSIGN-76]
	_ = x[SHR_ASSIGN-77]
	_ = x[SHL_ASSIGN-78]
	_ = x[REM_ASSIGN-79]
	_ = x[XOR_ASSIGN-80]
}

const _Token_name = "#illegal#eof#ident#commentbreakcasechanconstcontinuedefaultdeferelsefallthroughforfuncgogotoifimportinterfacemappackagerangereturnselectstructswitchtypevarliteral_beg<int><float><string>literal_end;:....,:=<-++--===!=!&^&^|&&||()[]{}<<=<<>>=>>/*+-%/=*=+=-=&^=&=|=>>=<<=%=^="

var _Token_index = [...]uint16{0, 8, 12, 18, 26, 31, 35, 39, 44, 52, 59, 64, 68, 79, 82, 86, 88, 92, 94, 100, 109, 112, 119, 124, 130, 136, 142, 148, 152, 155, 166, 171, 178, 186, 197, 198, 199, 200, 203, 204, 206, 208, 210, 212, 213, 215, 217, 218, 220, 221, 222, 223, 225, 227, 228, 229, 230, 231, 232, 233, 234, 236, 238, 239, 241, 243, 244, 245, 246, 247, 248, 250, 252, 254, 256, 259, 261, 263, 266, 269, 271, 273}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
