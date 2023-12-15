package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	gtype "github.com/tkxkd0159/buf-proto/buf-tutorial/pb/google/type"
	jstype "github.com/tkxkd0159/buf-proto/buf-tutorial/pb/js/v1"
)

var interfaceRegistry = make(map[string]reflect.Type)

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

	fmt.Println(strings.Repeat("=", 10), "<Marshal/Unmarshal protobuf type>", strings.Repeat("=", 10))
	if ptype.CreatedAt != nil {
		fmt.Printf("Proto TypeURL: %s\n", ptype.ProtoReflect().Descriptor().FullName())
	}
	// 1. Normal Marshal/Unmarshal
	seriealized, err := proto.Marshal(ptype)
	if err != nil {
		panic(err)
	}
	ptype2 := new(jstype.AllType)
	err = proto.Unmarshal(seriealized, ptype2)
	if err != nil {
		panic(err)
	}

	// 2. Can I unmarshal to extended type?
	originProto := &jstype.AnyTarget{Id: 2, Name: "ljs"}
	b, err := proto.Marshal(originProto)
	if err != nil {
		panic(err)
	}
	extendedProto := new(jstype.AnyTargetExtend)
	if err := proto.Unmarshal(b, extendedProto); err != nil {
		panic(err)
	}
	fmt.Printf("Success!!!\n")

	// 3. How can I process Any type?
	interfaceRegistry[string(ptype.ProtoReflect().Descriptor().FullName())] = reflect.TypeOf(new(jstype.AllType))
	ptypeAny, err := anypb.New(ptype)
	if err != nil {
		panic(err)
	}
	msg := reflectAny(ptypeAny, interfaceRegistry)

	if res := processMsg(msg); res != nil {
		fmt.Println(res)
	}
}

func reflectAny(any *anypb.Any, ir map[string]reflect.Type) proto.Message {
	fmt.Println(strings.Repeat("=", 10), "<reflect protobuf Any type>", strings.Repeat("=", 10))
	typeURL := strings.Split(any.GetTypeUrl(), "/")[1]
	t := ir[typeURL]
	v := reflect.New(t.Elem())
	fmt.Println(t, t.Kind(), t.Elem()) // *jsv1.AllType / ptr / jsv1.AllType
	if v.Kind() == reflect.Ptr {
		fmt.Println(v.Pointer(), v.Type(), v.Elem().Type()) // <uintptr> / *jsv1.AllType / jsv1.AllType
	}

	msg, ok := v.Interface().(proto.Message)
	if !ok {
		panic("not proto.Message")
	}

	err := proto.Unmarshal(any.GetValue(), msg)
	if err != nil {
		panic(err)
	}
	return msg
}

func processMsg(m proto.Message) *jstype.AllType {
	fmt.Println(strings.Repeat("=", 10), "<process interface type with reflection>", strings.Repeat("=", 10))
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	fmt.Println(t, t.Kind(), t.Elem()) // *jsv1.AllType / ptr / jsv1.AllType
	if v.Kind() == reflect.Ptr {
		fmt.Println(v.Pointer(), v.Type(), v.Elem().Type()) // <uintptr> / *jsv1.AllType / jsv1.AllType
	}
	if t.Elem().String() == "jsv1.AllType" {
		realValue, ok := m.(*jstype.AllType)
		if ok {
			return realValue
		}
	}

	return nil
}
