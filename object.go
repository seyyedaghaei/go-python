package python

// #include <Python.h>
import "C"

type Object C.PyObject

type PyObject interface {
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
	return togo(C.PyObject_Call(obj.C(), toObj(args).AsList().AsTuple().C(), toC(kwargs)))
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

func (obj *Object) AsFloat() *Float {
	return (*Float)(obj)
}

func (obj *Object) AsTuple() *Tuple {
	return (*Tuple)(obj)
}

func (obj *Object) AsList() *List {
	return (*List)(obj)
}

func (obj *Object) AsDict() *Dict {
	return (*Dict)(obj)
}

func (obj *Object) AsString() *String {
	return (*String)(obj)
}

func (obj *Object) Interface() interface{} {
	if obj != nil {
		switch (*obj.C()).ob_type {
		case &C.PyUnicode_Type:
			return obj.AsString().String()
		case &C.PyFloat_Type:
			return obj.AsFloat().Float64()
		case &C.PyDict_Type:
			return obj.AsDict().Map()
		case &C.PyList_Type:
			return obj.AsList().Slice()
		case &C.PyTuple_Type:
			return obj.AsTuple().Array()
		case &C.PyLong_Type:
			return obj.AsInt().Int()
		}
	}
	return nil
}

func (obj *Object) Type() *Object {
	return (*Object)(C.PyObject_Type(obj.C()))
}
