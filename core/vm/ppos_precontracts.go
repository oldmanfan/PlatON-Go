package vm

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/byteutil"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"reflect"
)

// Ppos pre-compiled contract address
var PrecompiledContractsPpos = map[common.Address]PrecompiledContract{
	common.CandidatePoolAddr: &CandidateContract{},
	common.TicketPoolAddr:    &TicketContract{},
}

var (
	ErrParamsRlpDecode = errors.New("Rlp decode fail")
	ErrParamsBaselen   = errors.New("Params Base length does not match")
	ErrParamsLen       = errors.New("Params length does not match")
	ErrUndefFunction   = errors.New("Undefined function")
	ErrCallRecode      = errors.New("Call recode error, panic...")
)

// execute decode input data and call the function.
func execute(input []byte, command map[string]interface{}) ([]byte, error) {
	log.Info("execute==> ", "input: ", hex.EncodeToString(input))
	defer func() {
		if err := recover(); nil != err {
			// catch call panic
			log.Error("execute==> ", "ErrCallRecode: ", ErrCallRecode.Error())
		}
	}()
	var source [][]byte
	if err := rlp.Decode(bytes.NewReader(input), &source); nil != err {
		log.Error("execute==> ", err.Error())
		return nil, ErrParamsRlpDecode
	}
	if len(source) < 2 {
		log.Error("execute==> ", "ErrParamsBaselen: ", ErrParamsBaselen.Error())
		return nil, ErrParamsBaselen
	}
	// get func and param list
	if _, ok := command[byteutil.BytesToString(source[1])]; !ok {
		log.Error("execute==> ", "ErrUndefFunction: ", ErrUndefFunction.Error())
		return nil, ErrUndefFunction
	}
	funcValue := command[byteutil.BytesToString(source[1])]
	paramList := reflect.TypeOf(funcValue)
	paramNum := paramList.NumIn()
	params := make([]reflect.Value, paramNum)
	if paramNum != len(source)-2 {
		log.Error("execute==> ", "ErrParamsLen: ", ErrParamsLen.Error())
		return nil, ErrParamsLen
	}
	for i := 0; i < paramNum; i++ {
		targetType := paramList.In(i).String()
		originByte := []reflect.Value{reflect.ValueOf(source[i+2])}
		params[i] = reflect.ValueOf(byteutil.Command[targetType]).Call(originByte)[0]
	}
	result := reflect.ValueOf(funcValue).Call(params)
	log.Info("execute==> ", "result[0]: ", result[0].Bytes())
	if _, err := result[1].Interface().(error); !err {
		return result[0].Bytes(), nil
	}
	log.Info(result[1].Interface().(error).Error())
	return result[0].Bytes(), result[1].Interface().(error)
}

// ResultCommon is the struct of transaction event.
type ResultCommon struct {
	Ret    bool
	Data   string
	ErrMsg string
}

// DecodeResultStr format the string of value to bytes.
func DecodeResultStr(result string) []byte {
	resultBytes := []byte(result)
	strHash := common.BytesToHash(common.Int32ToBytes(32))
	sizeHash := common.BytesToHash(common.Int64ToBytes(int64((len(resultBytes)))))
	var dataRealSize = len(resultBytes)
	if (dataRealSize % 32) != 0 {
		dataRealSize = dataRealSize + (32 - (dataRealSize % 32))
	}
	dataByt := make([]byte, dataRealSize)
	copy(dataByt[0:], resultBytes)

	finalData := make([]byte, 0)
	finalData = append(finalData, strHash.Bytes()...)
	finalData = append(finalData, sizeHash.Bytes()...)
	finalData = append(finalData, dataByt...)
	return finalData
}