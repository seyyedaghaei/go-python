package python

// #include<Python.h>
import "C"
import (
	"fmt"
	"reflect"
)

type Tuple Object

func NewTuple(size int) *Tuple {
	return (*Tuple)(C.PyTuple_New(C.ssize_t(size)))
}

func (tuple *Tuple) C() CPyObject {
	return tuple.Object().C()
}

func (tuple *Tuple) Object() *Object {
	return (*Object)(tuple)
}

func (tuple *Tuple) SetItem(index int, item interface{}) int {
	return int(C.PyTuple_SetItem(tuple.C(), C.ssize_t(index), toC(item)))
}

func (tuple *Tuple) GetItem(index int) *Object {
	return togo(C.PyTuple_GetItem(tuple.C(), C.ssize_t(index)))
}

func (tuple *Tuple) GetSlice(start int, stop int) *Tuple {
	return (*Tuple)(C.PyTuple_GetSlice(tuple.C(), C.ssize_t(start), C.ssize_t(stop)))
}

func (tuple *Tuple) Size() int {
	return int(C.PyTuple_Size(tuple.C()))
}

func (tuple *Tuple) String() string {
	return tuple.Object().String()
}

func (tuple *Tuple) Interface() interface{} {
	return tuple.Object().Interface()
}

func (tuple *Tuple) AsList() *List {
	size := tuple.Size()
	list := NewList(size)
	for i := 0; i < size; i++ {
		list.SetItem(i, tuple.GetItem(i))
	}
	return list
}

func (tuple *Tuple) Array() interface{} {
	fmt.Println()
	inter := reflect.TypeOf([]interface{}{}).Elem()
	arr := reflect.New(reflect.ArrayOf(tuple.Size(), inter)).Elem()
	for i := 0; i < tuple.Size(); i++ {
		i2 := tuple.GetItem(i).Interface()
		arr.Index(i).Set(reflect.ValueOf(&i2).Elem())
	}
	fmt.Println(arr.Type())
	return arr
}
