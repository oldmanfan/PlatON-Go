// create by platon
package byteutil

import (
	"bytes"
	"encoding/binary"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"math/big"
	"strings"
)

var Command = map[string]interface{}{
	"string":            BytesToString,
	"[]uint8":           OriginBytes,
	"[64]uint8":         BytesTo64Bytes,
	"[32]uint8":         BytesTo32Bytes,
	"int":               BytesToInt,
	"*big.Int":          BytesToBigInt,
	"uint32":            binary.BigEndian.Uint32,
	"uint64":            binary.BigEndian.Uint64,
	"int32":             common.BytesToInt32,
	"int64":             common.BytesToInt64,
	"float32":           common.BytesToFloat32,
	"float64":           common.BytesToFloat64,
	"discover.NodeID":   HexToNodeId,
	"[]discover.NodeID": ArrBytesToNodeId,
	"common.Hash":       common.BytesToHash,
	"[]common.Hash":     ArrBytesToHash,
	"common.Address":    HexToAddress,
}

func ArrBytesToHash(curByte []byte) []common.Hash {
	str := BytesToString(curByte)
	strArr := strings.Split(str, ":")
	var AHash []common.Hash
	for i := 0; i < len(strArr); i++ {
		AHash = append(AHash, common.HexToHash(strArr[i]))
	}
	return AHash
}

func ArrBytesToNodeId(curByte []byte) []discover.NodeID {
	str := BytesToString(curByte)
	strArr := strings.Split(str, ":")
	var ANodeID []discover.NodeID
	for i := 0; i < len(strArr); i++ {
		nodeId, _ := discover.HexID(strArr[i])
		ANodeID = append(ANodeID, nodeId)
	}
	return ANodeID
}

func BytesTo32Bytes(curByte []byte) [32]byte {
	var arr [32]byte
	copy(arr[:], curByte)
	return arr
}

func BytesTo64Bytes(curByte []byte) [64]byte {
	var arr [64]byte
	copy(arr[:], curByte)
	return arr
}

func OriginBytes(curByte []byte) []byte {
	return curByte
}

func BytesToBigInt(curByte []byte) *big.Int {
	return new(big.Int).SetBytes(curByte)
}

func BytesToInt(curByte []byte) int {
	bytesBuffer := bytes.NewBuffer(curByte)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	b := int(x)
	return b
}

func BytesToString(curByte []byte) string {
	return string(curByte)
}

func IntToBytes(curInt int) []byte {
	x := int32(curInt)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &x)
	return bytesBuffer.Bytes()
}

func Uint64ToBytes(val uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	return buf[:]
}

func HexToAddress(b []byte) common.Address {
	return common.HexToAddress(string(b))
}

func HexToNodeId(b []byte) discover.NodeID {
	nodeid, _ := discover.HexID(string(b))
	return nodeid
}