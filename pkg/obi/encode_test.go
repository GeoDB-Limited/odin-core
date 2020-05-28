package obi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ID uint8

type Inner struct {
	A ID    `obi:"a"`
	B uint8 `obi:"b"`
}

type ExampleData struct {
	Symbol string  `obi:"symbol"`
	Px     uint64  `obi:"px"`
	In     Inner   `obi:"in"`
	Arr    []int16 `obi:"arr"`
}

func TestEncodeBytes(t *testing.T) {
	require.Equal(t, MustEncode(ExampleData{
		Symbol: "BTC",
		Px:     9000,
		In: Inner{
			A: 1,
			B: 2,
		},
		Arr: []int16{10, 11},
	}), []byte{0x0, 0x0, 0x0, 0x3, 0x42, 0x54, 0x43, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x23, 0x28, 0x1, 0x2, 0x0, 0x0, 0x0, 0x2, 0x0, 0xa, 0x0, 0xb})
}

func TestDecode1(t *testing.T) {
	var n ExampleData
	err := Decode([]byte{0x0, 0x0, 0x0, 0x3, 0x42, 0x54, 0x43, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x23, 0x28, 0x1, 0x2, 0x0, 0x0, 0x0, 0x2, 0x0, 0xa, 0x0, 0xb}, &n)
	require.Nil(t, err)
	require.Equal(t, n, ExampleData{
		Symbol: "BTC",
		Px:     9000,
		In: Inner{
			A: 1,
			B: 2,
		},
		Arr: []int16{10, 11},
	})
}

// Uint8
func TestEncodeBytesUint8(t *testing.T) {
	num := uint8(123)
	require.Equal(t, []byte{num}, MustEncode(num))
}

func TestEncodeBytesAliasUint8(t *testing.T) {
	type ID uint8
	num := uint8(123)
	id := ID(num)
	require.Equal(t, []byte{num}, MustEncode(id))
}

// Uint16
func TestEncodeBytesUint16(t *testing.T) {
	num := uint16(123)
	require.Equal(t, []byte{0x7b, 0x0}, MustEncode(num))
}

func TestEncodeBytesAliasUint16(t *testing.T) {
	type ID uint16
	num := uint16(123)
	id := ID(num)
	require.Equal(t, []byte{0x7b, 0x0}, MustEncode(id))
}

// Uint32
func TestEncodeBytesUint32(t *testing.T) {
	num := uint32(123)
	require.Equal(t, []byte{0x7b, 0x0, 0x0, 0x0}, MustEncode(num))
}

func TestEncodeBytesAliasUint32(t *testing.T) {
	type ID uint32
	num := uint32(123)
	id := ID(num)
	require.Equal(t, []byte{0x7b, 0x0, 0x0, 0x0}, MustEncode(id))
}

// Uint64
func TestEncodeBytesUint64(t *testing.T) {
	num := uint64(123)
	require.Equal(t, []byte{0x7b, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, MustEncode(num))
}

func TestEncodeBytesAliasUint64(t *testing.T) {
	type ID uint64
	num := uint64(123)
	id := ID(num)
	require.Equal(t, []byte{0x7b, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, MustEncode(id))
}

// Int8
func TestEncodeBytesInt8(t *testing.T) {
	num := int8(-123)
	require.Equal(t, []byte{0x85}, MustEncode(num))
}

func TestEncodeBytesAliasInt8(t *testing.T) {
	type ID int8
	num := int8(-123)
	id := ID(num)
	require.Equal(t, []byte{0x85}, MustEncode(id))
}

// Int16
func TestEncodeBytesInt16(t *testing.T) {
	num := int16(-123)
	require.Equal(t, []byte{0x85, 0x00}, MustEncode(num))
}

func TestEncodeBytesAliasInt16(t *testing.T) {
	type ID int16
	num := int16(-123)
	id := ID(num)
	require.Equal(t, []byte{0x85, 0x00}, MustEncode(id))
}

// Int32
func TestEncodeBytesInt32(t *testing.T) {
	num := int32(-123)
	require.Equal(t, []byte{0x85, 0x00, 0x00, 0x00}, MustEncode(num))
}

func TestEncodeBytesAliasInt32(t *testing.T) {
	type ID int32
	num := int32(-123)
	id := ID(num)
	require.Equal(t, []byte{0x85, 0x00, 0x00, 0x00}, MustEncode(id))
}

// Int64
func TestEncodeBytesInt64(t *testing.T) {
	num := int64(-123)
	require.Equal(t, []byte{0x85, 0x00, 0x00, 0x00}, MustEncode(num))
}

func TestEncodeBytesAliasInt64(t *testing.T) {
	type ID int64
	num := int32(-123)
	id := ID(num)
	require.Equal(t, []byte{0x85, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, MustEncode(id))
}
