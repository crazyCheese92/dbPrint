package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type tFirst struct {
	id interface{int}
	name interface{string}
	isHere interface{bool}
	t interdace{null}
	arr inerface{[]}
	str interface{string, string}
}

type tSecond struct {
	calendar, autoContinue, buttons, status interface{bool}
	delay interface{int}
	list interface{[]}
}

func main(){
	file, _ := ioutil.ReadFile("test.json")
	data := tFirst{}
	_ = json.Unmarshal([]byte(file), &data)
	data1 := tSecond{}
	_ = json.Unmarshal([]byte(file), &data1)
	
	var first tFirst
	first.id, first.name, first.isHere, first.t, first.arr,  first.str = data.id, data.name, data.isHere, data.t, data.arr, data.str
	
	var second tSecond
	second.calendar, second.autoContinue, second.buttons, second.status, second.delay, second.list = data1.calendar, data1.autoContinue, data1.buttons, data1.status, data1.delay, data1.list
	
	fmt.Println(first)
	fmt.Println(second)
}
