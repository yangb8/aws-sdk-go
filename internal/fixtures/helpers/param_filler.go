package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/model/api"
	"github.com/awslabs/aws-sdk-go/internal/util"
)

func typeName(shape *api.Shape) string {
	return shape.GoType()
}

func ParamsStructFromJSON(value interface{}, shape *api.Shape) string {
	return util.GoFmt(paramsStructAny(value, shape))
}

func paramsStructAny(value interface{}, shape *api.Shape) string {
	switch shape.Type {
	case "structure":
		if value != nil {
			vmap := value.(map[string]interface{})
			return paramsStructStruct(vmap, shape)
		}
	case "list":
		vlist := value.([]interface{})
		return paramsStructList(vlist, shape)
	case "map":
		vmap := value.(map[string]interface{})
		return paramsStructMap(vmap, shape)
	case "string", "character":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.String(%#v)", v.Interface())
		}
	case "blob":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("[]byte(%#v)", v.Interface())
		}
	case "boolean":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Boolean(%#v)", v.Interface())
		}
	case "integer":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Integer(%v)", v.Interface())
		}
	case "float":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Float(%v)", v.Interface())
		}
	case "timestamp":
		v := reflect.Indirect(reflect.ValueOf(value))
		if v.IsValid() {
			return fmt.Sprintf("aws.Time(time.Unix(%d, 0))", int(v.Float()))
		}
	default:
		panic("Unhandled type " + shape.Type)
	}
	return ""
}

func paramsStructStruct(value map[string]interface{}, shape *api.Shape) string {
	out := "&" + typeName(shape)[1:] + "{\n"
	for _, n := range shape.MemberNames() {
		ref := shape.MemberRefs[n]
		lowcaseN := strings.ToLower(n[0:1]) + n[1:]
		name := ""
		if _, ok := value[n]; ok {
			name = n
		} else if _, ok = value[lowcaseN]; ok {
			name = lowcaseN
		}

		if val := paramsStructAny(value[name], ref.Shape); val != "" {
			out += fmt.Sprintf("%s: %s,\n", n, val)
		}
	}
	out += "}"
	return out
}

func paramsStructMap(value map[string]interface{}, shape *api.Shape) string {
	out := "&" + typeName(shape)[1:] + "{\n"
	for k, v := range value {
		out += fmt.Sprintf("%q: %s,\n", k, paramsStructAny(v, shape.ValueRef.Shape))
	}
	out += "}"
	return out
}

func paramsStructList(value []interface{}, shape *api.Shape) string {
	out := typeName(shape) + "{\n"
	for _, v := range value {
		out += fmt.Sprintf("%s,\n", paramsStructAny(v, shape.MemberRef.Shape))
	}
	out += "}"
	return out
}