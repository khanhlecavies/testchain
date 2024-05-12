package app_test

import (
	fmt "fmt"
	"testing"

	test "github.com/Cosmos-SDK-Learning-101/testchain/x/testchain"
	"github.com/cosmos/cosmos-sdk/codec"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"

	// "github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/tx"
	_ "github.com/cosmos/gogoproto/gogoproto"
	"github.com/cosmos/gogoproto/proto"
)

func TestExampleStruct(t *testing.T) {
	// Initialize InterfaceRegistry and ProtoCodec
	registry := types.NewInterfaceRegistry()

	// Register the struct with the interface registry
	registry.RegisterImplementations((*tx.TxExtensionOptionI)(nil), &test.ExampleStruct{})
	cdc := codec.NewProtoCodec(registry)

	// Create an instance of the struct
	example := &test.ExampleStruct{Id: 1, Name: "Cosmos"}

	// Marshal the struct to bytes using MarshalInterface
	marshalled, err := cdc.MarshalInterface(example)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshalled: %x\n", marshalled)

	registry = types.NewInterfaceRegistry()
	registry.RegisterInterface("Animal", (*proto.Message)(nil))
	registry.RegisterImplementations(
		(*proto.Message)(nil),
		&test.ExampleStruct{},
	)
	cdc = codec.NewProtoCodec(registry)

	// Unmarshal the bytes back to the struct using UnmarshalInterface
	var unmarshalled proto.Message
	err = cdc.UnmarshalInterface(marshalled, &unmarshalled)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshalled: %+v\n", unmarshalled)
	require.Equal(t, example, unmarshalled)

}
