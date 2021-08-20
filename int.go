package python

// #include <Python.h>
//PyObject * BuildIntValue(int num) {
//    return Py_BuildValue("i", num);
//}
import "C"

type Int C.PyObject

func (i *Int) C() CPyObject {
	return i.Object().C()
}

func (i *Int) Object() *Object {
	return (*Object)(i)
}

func (i *Int) Int() int {
	return int(C.PyFloat_AsDouble(i.C()))
}

func (i *Int) String() string {
	return i.Object().String()
}

func (i *Int) Interface() interface{} {
	return i.Object().Interface()
}

func PyInt(num int) *Int {
	return (*Int)(C.BuildIntValue(C.int(num)))
}
