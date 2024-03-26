package tool

import (
	"bytes"
	"fmt"
	"reflect"
)

func ToString(arg any) string {
	buf := bytes.NewBufferString("")
	_, err := fmt.Fprint(buf, arg)
	if err != nil {
		fmt.Errorf("ToString err %s %+v", arg, err)
		return ""
	}
	return buf.String()
}

func ToStringCompatiblePtr(arg any) string {
	if arg == nil {
		return ""
	}

	if reflect.TypeOf(arg).Kind() != reflect.Ptr {
		return ToString(arg)
	}

	return ToStringCompatiblePtr(reflect.Indirect(reflect.ValueOf(arg)))
}
