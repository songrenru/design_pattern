package singleton

import (
	"fmt"
	"sync"
)

// go实现单例要点
/**
 1. struct名称小写
 2. 方法名，对外暴露的，大写
 */

// 懒汉式: sync.Once()
// 饿汉式：直接声明

// 本例采用懒汉式
type singleton struct {}

func (s singleton) Do() {
	fmt.Println("singleton Do()")
}

var instance *singleton // 要指针，不然会赋值副本
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

