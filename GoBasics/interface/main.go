package main

type Usber interface {
	strat()
	stop()
}

// 如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type phone struct {
	Name string
}

// 手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p phone) strat() {
	println(p.Name, "start")
}
func (p phone) stop() {
	println(p.Name, "stop")
}
func main() {
	p := phone{
		Name: "华为手机",
	}
	//p.strat()

	var p1 Usber
	p1 = p
	p1.strat()
}
