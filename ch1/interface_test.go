package ch1

import (
	"fmt"
	"testing"
	"unsafe"
)

type GoProgram interface {

}

type Employee struct {
	Id int
	Name string
}

func TestInter(t *testing.T){
	e:= Employee{1,"bon"}
	e1 := Employee{Id:2,Name:"wang"}
	t.Log(e.Id)
	t.Log(e1.Name)
	e2:= new(Employee)//返回指针
	t.Logf("e is %T",&e)
	t.Logf("e2 is %T",e2)

	e3 := Employee{2,"dddd"}

	t.Log(e3.Strings())
	fmt.Printf("String is %x",unsafe.Pointer(&e3.Name))

	//t.Log(e3.Strings())
}

func(e *Employee) String() string {//todo 引用传值
	fmt.Printf("String is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%d----Name:%s",e.Id,e.Name)
}

func(e Employee) Strings() string {//todo 会有数据拷贝
	fmt.Printf("Strings is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%d-Name:%s",e.Id,e.Name)
}

type testInner func(op int)int
func Get(inner testInner) func(op int) int  {
	return func(op int) int {
		fmt.Print(op,"-----")
		return op*op +1
	}
}

func Show(op int) int {
	//return op*op*op
	return 0
}

func TestGet(t *testing.T)  {
	a := Get(Show)
	t.Log(a(2))
}
