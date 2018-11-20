package reflectUtil

import (
	//"fmt"
	"moqikaka.com/Test/src/Model"
	//"reflect"
	"fmt"
	"testing"
)

func TestReflect(context *testing.T) {
	// 反射调用方法
	//tempFunc := func(i int) int {
	//	return i
	//}
	//reflectResult := reflect.ValueOf(tempFunc)
	//fmt.Println("fv is reflect.Func ?", value.Kind() == reflect.Func)
	//params := make([]reflect.Value, 1)
	//params[0] = reflect.ValueOf(20)
	//result := reflectResult.Call(params)
	//fmt.Println(result[0])

	//newReflect := NewReflectUtil("张", 20)
	//value := reflect.ValueOf(newReflect)
	//params := make([]reflect.Value, 1)
	//params[0] = reflect.ValueOf(50)
	//fmt.Println(value.Method(2).Call(nil))
	//fmt.Println(value.Method(0).Call(nil)[0])
	//fmt.Println(value.Method(1).Call(params))
	//if value.Kind() == reflect.Ptr {
	//	value = value.Elem()
	//}
	//
	//value.FieldByName("Age").Set(reflect.ValueOf(2000))
	//fmt.Println()
	//typeOfT := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	f := value.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}

	//fmt.Println(value.Method(2).Call(nil)[0])
	//result := value.MethodByName("GetAge")
	//fmt.Println(result.Call(nil))
	//fmt.Println(value.Method(1).Call(nil)[0])

	newReflect := NewReflectUtil()
	tmepList, err := newReflect.CreateInstance(Model.NewHaremInfo(), "p_harem_info")
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range tmepList {
		haremInfo := value.(Model.HaremInfo)
		fmt.Println(haremInfo.PlayerID)
	}
}
