package Exception

import "errors"

type myError error

var (
	noError                myError = nil
	emptyCarNameError      myError = errors.New("차 이름은 적어도 1자 이상이어야 합니다")
	notAlphabetLetterError myError = errors.New("차 이름으로 알파벳을 입력해주세요")
)
