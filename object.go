package python

// #include <Python.h>
import "C"

type Object C.PyObject

type PyObject interface {
	C() CPyObject
	Object() *Object
}

func (obj *Object) C() CPyObject {
	return (CPyObject)(obj)
}

func (obj *Object) Object() *Object {
	return obj
}

func (obj *Object) GetAttribute(name string) *Object {
	return togo(C.PyObject_GetAttrString(obj.C(), C.CString(name)))
}

func (obj *Object) Call(args ...interface{}) *Object {
	return obj.CallFull(args, nil)
}

func (obj *Object) Func() func(args ...interface{}) *Object {
	return obj.Call
}

func (obj *Object) CallFull(args []interface{}, kwargs map[string]interface{}) *Object {
	tuple := NewTuple(len(args))
	for i, arg := range args {
		tuple.SetItem(i, arg)
	}
	kw := NewDict()
	for key, value := range kwargs {
		kw.SetItem(key, value)
	}
	return togo(C.PyObject_Call(obj.C(), tuple.C(), kw.C()))
}

func (obj *Object) CallOneArg(arg *Object) *Object {
	return togo(C.PyObject_CallOneArg(obj.C(), arg.C()))
}

func (obj *Object) CallNoArgs() *Object {
	return togo(C.PyObject_CallNoArgs(obj.C()))
}

func (obj *Object) String() string {
	return C.GoString(C.PyUnicode_AsUTF8(C.PyObject_Repr(obj.C())))
}

func (obj *Object) AsInt() *Int {
	return (*Int)(obj)
}
