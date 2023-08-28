// Code generated by "stringer -type ExpectCode -linecomment"; DO NOT EDIT.

package lexerr

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExpectedPackageKeyword-0]
	_ = x[ExpectedLetter-1]
	_ = x[ExpectedWhiteSpace-2]
	_ = x[ExpectedUnicodeLetter-3]
	_ = x[ExpectedComment-4]
	_ = x[FirstCommentCharacter-5]
	_ = x[SecondCommentCharacter-6]
}

const _ExpectCode_name = "'package' keywordletterwhite spaceunicode lettercommentfirst comment charactersecond comment character"

var _ExpectCode_index = [...]uint8{0, 17, 23, 34, 48, 55, 78, 102}

func (i ExpectCode) String() string {
	if i < 0 || i >= ExpectCode(len(_ExpectCode_index)-1) {
		return "ExpectCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExpectCode_name[_ExpectCode_index[i]:_ExpectCode_index[i+1]]
}
