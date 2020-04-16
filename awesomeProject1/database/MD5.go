package database

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"reflect"
	"unsafe"
)

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}


func Encode(data string) string {
	BASE64Table := "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}


func Decode(data string) string {
	BASE64Table := "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
	coder := base64.NewEncoding(BASE64Table)
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}