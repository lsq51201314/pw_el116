package elements

import (
	"pw_el116/utils/structEx"
)

type MaterialMajorType struct {
	ID   int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
}

func (MaterialMajorType) TableName() string {
	return "el_materialmajortype"
}

func (e *MaterialMajorType) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *MaterialMajorType) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
