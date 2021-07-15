package pkg1

import (
	"errors"
	"log"
	"plugin"
)

func init() {
	log.Println("pkg1 init")
}

type MyInterface interface {
	M1()
}

func LoadAndInvokeSomethingFromPlugin(pluginPath string) error {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}

	// 导出整型变量
	v, err := p.Lookup("V")
	if err != nil {
		return err
	}
	*v.(*int) = 15

	// 导出函数变量
	f, err := p.Lookup("F")
	if err != nil {
		return err
	}
	f.(func())()

	// 导出自定义类型变量
	f1, err := p.Lookup("Foo")
	if err != nil {
		return err
	}
	i, ok := f1.(MyInterface)
	if !ok {
		return errors.New("f1 does not implement MyInterface")
	}
	i.M1()

	return nil
}
