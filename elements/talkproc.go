package elements

import (
	"pw_el116/utils/structEx"
)

type TalkProc struct {
	RAW []byte
}

func (TalkProc) TableName() string {
	return "el_talkproc"
}

func (e *TalkProc) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *TalkProc) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
