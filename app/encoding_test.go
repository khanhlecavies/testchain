package app_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/stretchr/testify/require"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("Animal", (*testdata.Animal)(nil))
	registry.RegisterImplementations(
		(*testdata.Animal)(nil),
		&testdata.Dog{},
		&testdata.Cat{},
	)
	return registry
}

func TestMarshalAny(t *testing.T) {
	registry := types.NewInterfaceRegistry()
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&testdata.Cat{},
	)
	cdc := codec.NewProtoCodec(registry)

	kitty := &testdata.Cat{Moniker: "Kitty"}
	bz, err := cdc.MarshalInterface(kitty)
	require.NoError(t, err)

	var animal testdata.Animal

	// empty registry should fail
	err = cdc.UnmarshalInterface(bz, &animal)
	require.Error(t, err)

	// wrong type registration should fail
	registry.RegisterImplementations((*testdata.Animal)(nil), &testdata.Dog{})
	err = cdc.UnmarshalInterface(bz, &animal)
	require.Error(t, err)

	// should pass
	registry = NewTestInterfaceRegistry()
	cdc = codec.NewProtoCodec(registry)
	err = cdc.UnmarshalInterface(bz, &animal)
	require.NoError(t, err)
	require.Equal(t, kitty, animal)

	// nil should fail
	_ = NewTestInterfaceRegistry()
	err = cdc.UnmarshalInterface(bz, nil)
	require.Error(t, err)
}

func TestUnpackAny(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)

	privKey := ed25519.GenPrivKey()
	pk := privKey.PubKey()

	pkAny, err := types.NewAnyWithValue(pk)
	require.NoError(t, err)

	fmt.Println("pkAny:", pkAny)
	bz, err := cdc.Marshal(pkAny)
	require.NoError(t, err)

	var pkAny2 types.Any
	err = cdc.UnmarshalJSON(bz, &pkAny2)
	require.NoError(t, err)

	fmt.Println("pkAny :", pkAny2)

	var pkI cryptotypes.PubKey
	err = interfaceRegistry.UnpackAny(&pkAny2, &pkI)
	require.NoError(t, err)
	pk2 := pkAny2.GetCachedValue().(cryptotypes.PubKey)
	require.True(t, pk2.Equals(pk))
}
