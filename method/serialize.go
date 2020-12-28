package method

import (
	"encoding/json"
	"fmt"
)


/**
 * 序列化操作
 * @Sb:用于接收[]byte进行序列化
 */
type Sb struct {
	A []byte
}
func serialize() string {
	s := Sb{}
	a, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(a)
}


/**
 * 反序列化
 * @Sp:勇于接受反序列化的数据
 */
type Sp struct {
	data []byte
}
func deSerialization() Sp {
	var P Sp
	var d []byte
	json.Unmarshal(d, &P)
	return P
}