package ibc

import (
	"github.com/OCEChain/OCEChain/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(IBCTransferMsg{}, "cosmos-sdk/IBCTransferMsg", nil)
	cdc.RegisterConcrete(IBCReceiveMsg{}, "cosmos-sdk/IBCReceiveMsg", nil)
}
