package python

// #include <Python.h>
import "C"

type Object C.PyObject

type PyObject interface {
	Object() *Object
}

func (obj *Object) Object() *Object {
	return obj
}

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
	case PyObject:
		return o.Object()
	case int:
		return PyInt(o).Object()
	case string:
		return PyString(o)
	case float64:
		return PyFloat(o).Object()
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
	return C.GoString(C.PyUnicode_AsUTF8(C.PyObject_Repr(toc(obj))))
}
