package links_test

import (
	"fmt"
	"reflect"
	links "relax/Links"
	"testing"
	"unsafe"
)

type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

func TestInterfaceNil(t *testing.T) {
	var l links.Links

	t.Log("=== 接口变量声明但未赋值 ===")
	p := (*iface)(unsafe.Pointer(&l))
	t.Logf("tab(类型表指针):  %v", p.tab)
	t.Logf("data(数据指针):   %v", p.data)
	t.Logf("%%T:              %T", l)
	t.Logf("%%v:              %v", l)
	t.Logf("reflect.IsValid: %v", reflect.ValueOf(l).IsValid())

	fmt.Println("=== 接口变量声明但未赋值 ===")
	p = (*iface)(unsafe.Pointer(&l))
	fmt.Printf("tab(类型表指针):  %v\n", p.tab)
	fmt.Printf("data(数据指针):   %v\n", p.data)
	fmt.Printf("%%T:              %T\n", l)
	fmt.Printf("%%v:              %v\n", l)
	fmt.Printf("reflect.IsValid: %v\n", reflect.ValueOf(l).IsValid())
}

func TestInterfaceAssigned(t *testing.T) {
	l := links.NewLinkService()

	t.Log("=== 接口变量赋值后 ===")
	p := (*iface)(unsafe.Pointer(&l))
	t.Logf("tab(类型表指针):        %v (非 nil → 指向 itab)", p.tab)
	t.Logf("data(数据指针):         %v (非 nil → 指向 linkService 实例)", p.data)
	t.Logf("%%T:                    %T", l)
	t.Logf("%%v:                    %v", l)

	rv := reflect.ValueOf(l)
	rt := reflect.TypeOf(l)
	t.Logf("reflect.Kind:           %v", rv.Kind())
	t.Logf("reflect.Type:           %v", rt)
	t.Logf("reflect.NumMethod:      %d", rt.NumMethod())
	for i := range rt.NumMethod() {
		m := rt.Method(i)
		t.Logf("  method[%d]: %s %s", i, m.Name, m.Type)
	}

	fmt.Println("=== 接口变量赋值后 ===")
	p = (*iface)(unsafe.Pointer(&l))
	fmt.Printf("tab(类型表指针):   %v\n", p.tab)
	fmt.Printf("data(数据指针):    %v\n", p.data)
	fmt.Printf("%%T:               %T\n", l)
	fmt.Printf("%%v:               %v\n", l)

	rv = reflect.ValueOf(l)
	rt = reflect.TypeOf(l)
	fmt.Printf("reflect.Kind:      %v\n", rv.Kind())
	fmt.Printf("reflect.Type:      %v\n", rt)
	fmt.Printf("reflect.NumMethod: %d\n", rt.NumMethod())
	for i := range rt.NumMethod() {
		m := rt.Method(i)
		fmt.Printf("  method[%d]: %s %s\n", i, m.Name, m.Type)
	}

	_ = l.NewLinkNode
}
