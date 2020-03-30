package example

import (
	"encoding/json"
	"log"
	"testing"
)

type customString string

func TestDefaults(t *testing.T) {

	type SomeStruct struct {
		Inta  int   `example:"10"`
		Int8  int8  `example:"11"`
		Int16 int16 `example:"12"`
		Int32 int32 `example:"13"`
		Int64 int64 `example:"14"`

		UInta  uint   `example:"15"`
		UInt8  uint8  `example:"16"`
		UInt16 uint16 `example:"17"`
		UInt32 uint32 `example:"18"`
		UInt64 uint64 `example:"19"`

		Float32 float32 `example:"20.1"`
		Float64 float64 `example:"20.2"`

		String       string       `example:"string"`
		CustomString customString `example:"custom string"`

		Bool bool `example:"true"`

		IntaPtr  *int   `example:"10"`
		Int8Ptr  *int8  `example:"11"`
		Int16Ptr *int16 `example:"12"`
		Int32Ptr *int32 `example:"13"`
		Int64Ptr *int64 `example:"14"`

		UIntaPtr  *uint   `example:"15"`
		UInt8Ptr  *uint8  `example:"16"`
		UInt16Ptr *uint16 `example:"17"`
		UInt32Ptr *uint32 `example:"18"`
		UInt64Ptr *uint64 `example:"19"`

		Float32Ptr *float32 `example:"20.1"`
		Float64Ptr *float64 `example:"20.2"`

		StringPtr       *string       `example:"string"`
		CustomStringPtr *customString `example:"custom string"`

		BoolPtr *bool `example:"true"`

		StringPtrNoVal *string `example:"wow"`
	}
	a := SomeStruct{}

	Process(&a)
	prettyPrint(a)
}

func TestDefaultsWithSliceStructs(t *testing.T) {
	type AnotherStruct struct {
		Name string `example:"this should be the value"`
	}
	type SomeStruct struct {
		AnotherStruct []AnotherStruct `example:"len=2"`
	}
	a := SomeStruct{}

	Process(&a)
	prettyPrint(a)
}

func TestDefaultsWithStructs(t *testing.T) {
	type AnotherStruct struct {
		Name string `example:"name name"`
	}
	type SomeStruct struct {
		AnotherStruct AnotherStruct
	}
	a := SomeStruct{}

	Process(&a)
	prettyPrint(a)
}

func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	log.Println(string(s))
}

func TestGetLen(t *testing.T) {
	log.Printf(`getLen("len=1"): %#+v\n`, getLen("len=1"))
}
