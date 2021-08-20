package python

// #include <Python.h>
import "C"

type Dict C.PyObject

func NewDict() *Dict {
	return (*Dict)(C.PyDict_New())
}

func (dict *Dict) C() CPyObject {
	return dict.Object().C()
}

func (dict *Dict) Object() *Object {
	return (*Object)(dict)
}

func (dict *Dict) SetItem(key interface{}, value interface{}) int {
	return int(C.PyDict_SetItem(dict.C(), toC(key), toC(value)))
}

func (dict *Dict) String() string {
	return dict.Object().String()
}

func (dict *Dict) Interface() interface{} {
	return dict.Object().Interface()
}

func (dict *Dict) Clear() {
	C.PyDict_Clear(dict.C())
}

func (dict *Dict) DelItem(key interface{}) int {
	return int(C.PyDict_DelItem(dict.C(), toC(key)))
}

func (dict *Dict) Copy() *Dict {
	return (*Dict)(C.PyDict_Copy(dict.C()))
}

func (dict *Dict) GetItem(key interface{}) *Object {
	return (*Object)(C.PyDict_GetItem(dict.C(), toC(key)))
}

func (dict *Dict) Size() int {
	return int(C.PyDict_Size(dict.C()))
}

func (dict *Dict) Items() *List {
	return (*List)(C.PyDict_Items(dict.C()))
}

func (dict *Dict) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	items := dict.Items()
	for i := 0; i < items.Size(); i++ {
		item := items.GetItem(i).AsTuple()
		key, value := item.GetItem(0).Interface(), item.GetItem(1).Interface()
		m[key] = value
	}
	return m
}
