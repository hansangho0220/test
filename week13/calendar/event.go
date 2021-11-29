package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string

	//Mydate Date
	//익명 구조체
	Date
}

func (t *Event) SetTitle(title string) error {
	//유니코드 방식으로 글자수를 센다. -> 3글자 이상이면 에러 리턴
	if utf8.RuneCountInString(title) > 3 {
		return errors.New("이벤트는 3글자 이내로 작성하셔야 합니다.")
	}

	t.title = title
	return nil
}

func (t *Event) GetTitle() string {
	return t.title
}
