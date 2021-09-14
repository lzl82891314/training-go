package base

type Instance interface {
	DoSomething()
}

var instanceMap = make(map[string]func() Instance, 10)

func Register(name string, creatingFunc func() Instance) {
	instanceMap[name] = creatingFunc
}

func Create(name string) Instance {
	if f, ok := instanceMap[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
