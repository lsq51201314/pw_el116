package byteEx

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math"
	"os"
)

func GetFileByte(filepath string, offset, length int64) (data []byte, err error) {
	var file *os.File
	if file, err = os.Open(filepath); err == nil {
		buf := make([]byte, length)
		var n int
		if n, err = file.ReadAt(buf, offset); err == nil {
			data = buf[:n]
		}
		_ = file.Close()
	}
	return
}

func BytesToIntLittle(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	_ = binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int(x)
}

func ByteToFloatLittle(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func WStrByteToStr(wsbyte []byte) string {
	c := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	d, ok := c.Bytes(wsbyte)
	if ok != nil {
		return ""
	}
	d = bytes.TrimRight(d, "\x00")
	return string(d)
}

func ByteToStr(sbyte []byte) string {
	d, ok := ioutil.ReadAll(transform.NewReader(bytes.NewReader(sbyte), simplifiedchinese.GBK.NewDecoder()))
	if ok != nil {
		return ""
	}
	d = bytes.TrimRight(d, "\x00")
	return string(d)
}
