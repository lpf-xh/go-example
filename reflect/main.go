package main

import (
	"fmt"
	"reflect"
)

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *point) Print() {
	fmt.Println("point is",p)
}

func main() {
	a := point{
		X: 1,
		Y: 1,
	}

	fmt.Println("original point:",a)

	// 获取 reflect.Value
	av := reflect.ValueOf(&a)
	// 获取 reflect.Type
	at := av.Type()

	// reflect.Value 转 interface
	inf := interface{}(av)
	fmt.Println(inf)

	// 修改字段值
	// 使用Elem获取指针对应的值，才能修改
	av.Elem().FieldByName("Y").SetInt(2)
	fmt.Println("modified point:",a)

	fmt.Println()
	
	// 打印类型
	fmt.Println("av kind:", av.Kind())
	fmt.Println("at kind:", at.Kind())
	fmt.Println("av.Elem kind:", av.Elem().Kind())
	fmt.Println("at.Elem kind:", at.Elem().Kind())
	

	fmt.Println()

	// 打印字段名称和tag
	for i := 0; i < at.Elem().NumField(); i++ {
		f := at.Elem().Field(i)
		fmt.Printf("field %d name: %s\n", i, f.Name)
		fmt.Printf("field %d tag:  %s\n", i, f.Tag.Get("json"))
	}

	fmt.Println()

	// 调用方法
	av.Method(0).Call(nil)
}


