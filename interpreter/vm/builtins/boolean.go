package builtins

type trueClass struct {
	valueStub
}

func NewTrueClass() Value {
	o := &trueClass{}
	o.initialize()
	o.class = NewClassValue().(Class)
	return o
}

func (obj *trueClass) String() string {
	return "True"
}

type true struct {
	valueStub
}

func (obj *trueClass) New() Value {
	o := &true{}
	o.initialize()
	o.class = obj

	return o
}

type falseClass struct {
	valueStub
}

func NewFalseClass() Value {
	o := &falseClass{}
	o.initialize()
	o.class = NewClassValue().(Class)
	return o
}

func (obj *falseClass) String() string {
	return "False"
}

type false struct {
	valueStub
}

func (obj *falseClass) New() Value {
	o := &false{}
	o.initialize()
	o.class = obj

	return o
}
