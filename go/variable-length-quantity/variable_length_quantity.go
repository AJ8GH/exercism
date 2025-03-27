package variablelengthquantity

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func EncodeVarint(input []uint32) []byte {
	// out := make([]byte, 4)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, input)
	// for _, v := range input {
	// 	binary.Write(buf, binary.LittleEndian, v)
	// 	// binary.LittleEndian.PutUint32(out, v)
	// }
	return buf.Bytes()
}

func Encode(input int) []byte {
	return []byte(strconv.Itoa(input))
}

func DecodeVarint(input []byte) ([]uint32, error) {
	panic("Please implement the DecodeVarint function")
}
