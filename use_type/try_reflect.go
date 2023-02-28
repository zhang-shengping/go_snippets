package main

import (
	"fmt"
	"reflect"
)

type MyInt int

type Tester interface {
	Test(t string) string
}

func (MyInt) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (MyInt) Test(t string) string {
	return t
}

func try_reflect() {
	var r interface{} = MyInt(1)
	var test Tester = Tester(MyInt(2))
	t := fmt.Sprintf("%T", r)
	z := fmt.Sprintf("%T", test)
	fmt.Println(t)
	fmt.Println(z)
	fmt.Println("type: ", reflect.TypeOf(test))
	fmt.Println("value: ", reflect.ValueOf(test).String())
	fmt.Println("value: ", reflect.ValueOf(test))
	v := reflect.ValueOf(r)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is ", v.Kind())
	fmt.Println("kind is ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Int())
	// use value return to interface, reflect back to interface
	w := v.Interface()
	i := w.(Tester)
	fmt.Println("the fmt Printf auto unpack interface v as the reflect.Value : ", v)
	fmt.Printf("use value return to interface %T \n", w)
	fmt.Printf("assert interface is Tester type %T\n ", i)

	// settability
	// settability is determined by whether the reflection object holds the original item.
	// var x float64 = 3.4
	// NOTE: this is can not be set, we pass the copy of x to the reflect.ValueOf.
	// y := reflect.ValueOf(x)
	// Just keep in mind that reflection Values need the address of something in order to modify what they represent
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	u := p.Elem()
	// can not use this to get value of pointer e := *p
	fmt.Printf("u is %T \n", u)
	fmt.Println("settability of v:", u.CanSet())
	u.SetFloat(123.456)
	fmt.Println(u)
	fmt.Println(x)
}

// package main

// import (
//    "fmt"
//    "reflect"
// )

// func main() {
//    j := 42

// 	p := &j
// 	fmt.Println(p)
// 	fmt.Println(*p)
// 	fmt.Println(reflect.ValueOf(p).Elem())
// 	fmt.Println(reflect.ValueOf(p).Type())
// 	fmt.Println(reflect.ValueOf(p).Kind())
// }
