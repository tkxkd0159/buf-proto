package main

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/anypb"

	gtype "github.com/tkxkd0159/buf-proto/buf-tutorial/pb/google/type"
	jstype "github.com/tkxkd0159/buf-proto/buf-tutorial/pb/js/v1"
)

func main() {
	tz := gtype.DateTime_TimeZone{TimeZone: &gtype.TimeZone{Id: "Asia/Seoul"}}
	currentTime := time.Now()
	ts := gtype.DateTime{
		Year:       int32(currentTime.Year()),
		Month:      int32(currentTime.Month()),
		Day:        int32(currentTime.Day()),
		Hours:      int32(currentTime.Hour()),
		Minutes:    int32(currentTime.Minute()),
		Seconds:    int32(currentTime.Second()),
		Nanos:      int32(currentTime.Nanosecond()),
		TimeOffset: &tz,
	}

	anyTarget := &jstype.AnyTarget{Id: 1, Name: "ljs"}
	anyv, err := anypb.New(anyTarget)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Any: %s -> %v\n", anyv.TypeUrl, anyv.Value)

	extendedAny := new(jstype.AnyTargetExtend)
	if err := anyv.UnmarshalTo(extendedAny); err != nil {
		fmt.Println(err) // mismatched message type: got "js.v1.AnyTargetExtend", want "js.v1.AnyTarget"
	}
	origin := new(jstype.AnyTarget)
	if err := anyv.UnmarshalTo(origin); err != nil {
		panic(err)
	}

	ptype := &jstype.AllType{
		Bool:      true,
		ByteSlice: []byte("hello"),
		String_:   []string{"hello", "world"},
		Map:       map[string]string{"hello": "world"},
		Enum:      jstype.EnumType_ENUM_TYPE_STARTED,
		Any:       []*jstype.AnyType{{Any: anyv}},
		CreatedAt: &ts,
		Float32:   1.1,
		Float64:   -2.23,
		Int32:     1,
		Int64:     -2,
		Int32_2:   5, // sint32
		Int64_2:   6,
		Int32_3:   7, // sfixed32
		Int64_3:   8,
		Uint32:    3,
		Uint64:    4,
		Uint32_2:  5, // fixed32
		Uint64_2:  6,
	}

	if ptype.CreatedAt != nil {
		fmt.Printf("%s\nvalues:\n\t%s", ptype.ProtoReflect().Descriptor().FullName(), ptype)
	}
}
