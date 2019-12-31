package utils

import (
	"fmt"
	"github.com/astaxie/beego/validation"

	"Go-blog-server/pkg/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Println(err.Message)
		logging.Info(err.Key, err.Message)
	}

	return
}
// func QueryFilter(model interface{}, filterkeys []string) string {
// 	var res string
// 	fields := reflect.TypeOf(model)

// 	num := fields.NumField()

// 	for i := 0; i < num; i++ {
// 		field := fields.Field(i)
// 		if()
// 		index := arrays.ContainsString(filterkeys, field.Name)
// 		if(index == -1){
// 			fmt.Println(field.Name)
// 			res += field.Name
// 			res += ","
// 		}
// 	}

// 	return res
// }

// func Display(name string, x interface{}) {
//     display(name, reflect.ValueOf(x))
// }

// func display(path string, v reflect.Value) {
//     switch v.Kind() {
//     case reflect.Invalid:
//         fmt.Printf("%s = invalid\n", path)
//     case reflect.Slice, reflect.Array:
//         for i := 0; i < v.Len(); i++ {
//             display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
//         }
//     case reflect.Struct:
//         for i := 0; i < v.NumField(); i++ {
//             fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
//             display(fieldPath, v.Field(i))
//         }
//     case reflect.Map:
//         for _, key := range v.MapKeys() {
//             display(fmt.Sprintf("%s[%s]", path,
//                 formatAtom(key)), v.MapIndex(key))
//         }
//     case reflect.Ptr:
//         if v.IsNil() {
//             fmt.Printf("%s = nil\n", path)
//         } else {
//             display(fmt.Sprintf("(*%s)", path), v.Elem())
//         }
//     case reflect.Interface:
//         if v.IsNil() {
//             fmt.Printf("%s = nil\n", path)
//         } else {
//             fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
//             display(path+".value", v.Elem())
//         }
//     default: // basic types, channels, funcs
//         fmt.Printf(v.Type().Field().Name)
//     }
// }

// func formatAtom(v reflect.Value) string {
//     switch v.Kind() {
//     case reflect.Invalid:
//         return "invalid"
//     case reflect.Int, reflect.Int8, reflect.Int16,
//         reflect.Int32, reflect.Int64:
//         return strconv.FormatInt(v.Int(), 10)
//     case reflect.Uint, reflect.Uint8, reflect.Uint16,
//         reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//         return strconv.FormatUint(v.Uint(), 10)
//     // ...floating-point and complex cases omitted for brevity...
//     case reflect.Bool:
//         return strconv.FormatBool(v.Bool())
//     case reflect.String:
//         return strconv.Quote(v.String())
//     case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
//         return v.Type().String() + " 0x" +
//             strconv.FormatUint(uint64(v.Pointer()), 16)
//     default: // reflect.Array, reflect.Struct, reflect.Interface
//         return v.Type().String() + " value"
//     }
// }

// func fieldSet(fields ...string) map[string]bool {
// 	set := make(map[string]bool, len(fields))
// 	for _, s := range fields {
// 		set[s] = true
// 	}
// 	return set
// }

// func (model interface{}) SelectFields(fields ...string) map[string]interface{} {
// 	fs := fieldSet(fields...)
// 	rt, rv := reflect.TypeOf(*s), reflect.ValueOf(*s)
// 	out := make(map[string]interface{}, rt.NumField())
// 	for i := 0; i < rt.NumField(); i++ {
// 		field := rt.Field(i)
// 		jsonKey := field.Tag.Get("json")
// 		if fs[jsonKey] {
// 			out[jsonKey] = rv.Field(i).Interface()
// 		}
// 	}
// 	return out
// }

// func (model interface{}) UnSelectFields(fields ...string) map[string]interface{} {
// 	fs := fieldSet(fields...)
// 	rt, rv := reflect.TypeOf(*s), reflect.ValueOf(*s)
// 	out := make(map[string]interface{}, rt.NumField())
// 	for i := 0; i < rt.NumField(); i++ {
// 		field := rt.Field(i)
// 		jsonKey := field.Tag.Get("json")
// 		if !fs[jsonKey] {
// 			out[jsonKey] = rv.Field(i).Interface()
// 		}
// 	}
// 	return out
// }