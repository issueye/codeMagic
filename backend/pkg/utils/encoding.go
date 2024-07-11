package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

func Base64Encode(value []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(value)))
	base64.StdEncoding.Encode(dst, value)
	return dst
}

func Base64EncodeString(value string) string {
	data := String2Bytes(value)
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data []byte) (value []byte, err error) {
	src := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(src, data)
	return src[:n], err
}

func Base64DecodeString(data string) (value string, err error) {
	var dst []byte
	dst, err = base64.StdEncoding.DecodeString(data)
	if err == nil {
		value = Bytes2String(dst)
	}
	return
}

func Serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func UnSerialize(valueBytes []byte, registers ...interface{}) (value interface{}, err error) {
	for _, v := range registers {
		gob.Register(v)
	}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&value)
	return
}
