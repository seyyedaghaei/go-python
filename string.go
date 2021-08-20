package python

// #include <Python.h>
//PyObject * BuildStringValue(const char * str, int size) {
//    return Py_BuildValue("s#", str, size);
//}
import "C"

type String C.PyObject

func (str *String) C() CPyObject {
	return str.Object().C()
}

func (str *String) Object() *Object {
	return (*Object)(str)
}

func (str *String) String() string {
	return C.GoString(C.PyUnicode_AsUTF8(str.C()))
}

func (str *String) Interface() interface{} {
	return str.Object().Interface()
}

func PyString(str string) *String {
	return (*String)(C.BuildStringValue(C.CString(str), C.int(len(str))))
}
