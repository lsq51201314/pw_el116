package elements

import (
	"pw_el116/utils/structEx"
)

//EquipmentAddon 装备附加属性
type EquipmentAddon struct {
	ID        int     `elements:"name:id;type:int;size:4;text:附加属性ID;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name      string  `elements:"name:name;type:unicode;size:64;text:附加属性名称;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	NumParams int     `elements:"name:num_params;type:int;size:4;text:参数个数;" gorm:"column:num_params;type:integer;default:0;not null;" json:"num_params"`
	Param1    float32 `elements:"name:param1;type:int;size:4;text:第1个参数;" gorm:"column:param1;type:float;default:0;not null;" json:"param1"`
	Param2    float32 `elements:"name:param2;type:int;size:4;text:第2个参数;" gorm:"column:param2;type:float;default:0;not null;" json:"param2"`
	Param3    float32 `elements:"name:param3;type:int;size:4;text:第3个参数;" gorm:"column:param3;type:float;default:0;not null;" json:"param3"`
}

func (EquipmentAddon) TableName() string {
	return "el_equipmentaddon"
}

func (e *EquipmentAddon) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *EquipmentAddon) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
