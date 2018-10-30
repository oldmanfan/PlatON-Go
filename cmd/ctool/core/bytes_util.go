package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"strconv"
	"strings"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func Int32ToBytes(n int32) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int32(tmp)
}

func Int64ToBytes(n int64) []byte {
	tmp := int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int64(tmp)
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
func BoolToBytes(b bool) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, b)
	return buf.Bytes()
}

/**
  []string 转[]byte
*/
func convertStringsToBytes(content []string) []byte {
	byteContent := "\x00" + strings.Join(content, "\x02\x00")
	return []byte(byteContent)
}

func BytesConverter(source []byte, t string) interface{} {
	var dest interface{}
	switch t {
	case "int32":
		dest = BytesToInt32(source)
	case "int64":
		dest = BytesToInt64(source)
	case "float32":
		dest = BytesToFloat32(source)
	case "float64":
		dest = BytesToFloat64(source)
	case "string":
		dest = string(source[:])
	default:
		dest = source
	}
	return dest
}

func StringConverter(source string, t string) ([]byte, error) {
	switch t {
	case "int32", "uint32", "uint", "int":
		dest, err := strconv.Atoi(source)
		return Int32ToBytes(int32(dest)), err
	case "int64", "uint64":
		dest, err := strconv.ParseInt(source, 10, 64)
		return Int64ToBytes(dest), err
	case "float32":
		dest, err := strconv.ParseFloat(source, 32)
		return Float32ToBytes(float32(dest)), err
	case "float64":
		dest, err := strconv.ParseFloat(source, 64)
		return Float64ToBytes(dest), err
	case "bool":
		if "true" == source {
			return BoolToBytes(true), nil
		} else if "false" == source {
			return BoolToBytes(false), nil
		} else {
			return []byte{}, errors.New("invalid boolean param")
		}
	default:
		return []byte(source), nil
	}

}
