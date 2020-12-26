package method

import (
	"encoding/json"
	"fmt"
)

type Sb struct {
	A []byte
}
/**
 * 序列化操作
 */
func serialize() string {
	s := Sb{}
	a, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(a)
}

type Sp struct {
	data []byte
}
/**
 * 反序列化
 */
func deSerialization() Sp {
	var P Sp
	var d []byte
	json.Unmarshal(d, &P)
	return P
}