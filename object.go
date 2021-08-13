package pyhton

// #include <Python.h>
//const char * PyObjectAsString(PyObject* object) {
//    PyObject* repr = PyObject_Repr(object);
//    PyObject* str = PyUnicode_AsEncodedString(repr, "utf-8", "~E~");
//    return PyBytes_AS_STRING(str);
//}
import "C"

type Object C.PyObject

func (obj *Object) GetAttribute(name string) *Object {
	return togo(C.PyObject_GetAttrString(toc(obj), C.CString(name)))
}

func (obj *Object) Call(args ...interface{}) *Object {
	return obj.CallFull(args, nil)
}

func (obj *Object) Func() func(args ...interface{}) *Object {
	return obj.Call
}

func toObj(obj interface{}) *Object {
	switch o := obj.(type) {
	case *Object:
		return o
	case int:
		return Int(o)
	case string:
		return String(o)
	}
	switch o := obj.(type) {
	case *Dict:
		return (*Object)(o)
	}
	switch o := obj.(type) {
	case *Tuple:
		return (*Object)(o)
	}
	return nil
}

func (obj *Object) CallFull(args []interface{}, kwargs map[string]interface{}) *Object {
	tuple := NewTuple(len(args))
	for i, arg := range args {
		tuple.SetItem(i, toObj(arg))
	}
	kw := NewDict()
	for key, value := range kwargs {
		kw.SetItem(key, toObj(value))
	}
	return togo(C.PyObject_Call(toc(obj), toc((*Object)(tuple)), toc((*Object)(kw))))
}

func (obj *Object) CallOneArg(arg *Object) *Object {
	return togo(C.PyObject_CallOneArg(toc(obj), toc(arg)))
}

func (obj *Object) CallNoArgs() *Object {
	return togo(C.PyObject_CallNoArgs(toc(obj)))
}

func (obj *Object) String() string {
	return C.GoString(C.PyObjectAsString(toc(obj)))
}
