package elements

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"pw_el116/global"
	"pw_el116/utils/byteEx"
	"reflect"
)

var El116Ver = "06000030"

func Init() (err error) {
	var file *os.File
	if file, err = os.Open("./elements.data"); err != nil {
		return
	}
	defer file.Close()
	//判断版本
	ver := make([]byte, 4)
	if _, err = file.Read(ver); err != nil {
		return
	}
	if hex.EncodeToString(ver) != El116Ver {
		err = errors.New("elements.data 文件版本错误。")
		return
	}
	var offset int64
	for i := 0; i < len(ReadMap); i++ {
		offset += 4
		d := make([]byte, 4)
		if _, err = file.Read(d); err != nil {
			return
		}
		if i == EL_TalkProc {
			//暂时跳过
			p := make([]byte, 1917216)
			file.Read(p)
		} else {
			size := byteEx.BytesToIntLittle(d)
			tStruct := ReadMap[i]
			tValue := reflect.ValueOf(tStruct)
			tResult := tValue.Method(0).Call(nil)
			if tResult[1].Interface() != nil {
				err = errors.New(fmt.Sprintf("结构体:%v 获取长度失败。", tValue.Type()))
				return
			}
			tSize := tResult[0].Int()
			tData := make([]byte, tSize)
			params := make([]reflect.Value, 1)
			for j := 0; j < size; j++ {
				offset += tSize
				if _, err = file.Read(tData); err != nil {
					return
				}
				params[0] = reflect.ValueOf(tData)
				tValue.Method(1).Call(params)
				fmt.Println(fmt.Sprintf("%v/%v %v/%v-%v/%v %v %v", i+1, len(ReadMap), j+1, size, offset+4, tSize, tValue.Type(), tStruct))
				global.MDB.Create(tStruct)
			}
		}
	}
	return
}
