package obc

import (
	"encoding/hex"
	"os"

	"github.com/cleey/can-go/pkg/common"
	"github.com/cleey/can-go/pkg/descriptor"
	"github.com/cleey/can-go/pkg/generate"
	// "obc_core/pkg/dbc"
)

type ObcCore struct {
	messages map[uint32]*descriptor.Message
}

type ObcParam struct {
	Name string
	Val  float64
	Unit string
}

func NewObcByDBC(dbcFile string) (*ObcCore, error) {
	obc := &ObcCore{}
	err := obc.Init(dbcFile)
	return obc, err
}

func (o *ObcCore) Init(dbcFile string) error {
	// Load the DBC file
	data, err := os.ReadFile(dbcFile)
	if err != nil {
		return err
	}
	compileRet, err := generate.Compile(dbcFile, data)
	if err != nil {
		return err
	}

	var msgs = map[uint32]*descriptor.Message{}
	for _, v := range compileRet.Database.Messages {
		msgs[v.ID] = v
	}
	o.messages = msgs
	return nil
}

func (o *ObcCore) ParseCanHexStr(canID uint32, hexStr string) []ObcParam {
	hexByte, _ := hex.DecodeString(hexStr)

	params := []ObcParam{}
	dbc_bo, ok := o.messages[canID]
	if !ok {
		return params
	}

	datanew := common.Data{}
	copy(datanew[:], hexByte)

	for _, s := range dbc_bo.Signals {
		params = append(params, ObcParam{
			Name: s.Name,
			Val:  s.UnmarshalPhysical(datanew),
			Unit: s.Unit,
		})
	}

	return params
}
