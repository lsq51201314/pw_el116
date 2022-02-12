package elements

import (
	"pw_el116/utils/structEx"
)

type NpcRepairService struct {
	ID       int    `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name     string `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	IdDialog int    `elements:"name:id_dialog;type:int;size:4;text:未知;" gorm:"column:id_dialog;type:integer;default:0;not null;" json:"id_dialog"`
}

func (NpcRepairService) TableName() string {
	return "el_npcrepairservice"
}

func (e *NpcRepairService) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *NpcRepairService) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
