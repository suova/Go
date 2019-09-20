package main

import (
	"fmt"
	"testing"
)

func Test_pushString(t *testing.T){
	arr :=[]string{"a","c","b"}
	elem := "c"
	exp := []string{"a","c","b","c"}
	res:=pushString(arr,elem)
	for i:=0;i<len(exp);i++{
		if res[i]!=exp[i]{
			t.Error("expected", exp, "result",res)
		}
	}
	fmt.Println("Test_pushString -- OK")
}
func Test_popString(t *testing.T){
	exp :=[]string{"a","c","b",}
	elem := "c"
	arr := []string{"a","c","b","c"}
	elemRes, arrRes:=popString(arr)
	for i:=0;i<len(exp);i++{
		if arrRes[i]!=exp[i]{
			t.Error("expected", exp, "result",arrRes)
		}
	}
	if elem != elemRes{
		t.Error("expected", elem, "result",elemRes)
	}
	fmt.Println("Test_popString -- OK")
}

func Test_push(t *testing.T){
	arr :=[]int{1,2,3}
	elem := 4
	exp := []int{1,2,3,4}
	res:=push(arr,elem)
	for i:=0;i<len(exp);i++{
		if res[i]!=exp[i]{
			t.Error("expected", exp, "result",res)
		}
	}
	fmt.Println("Test_push -- OK")
}
func Test_pop(t *testing.T){
	exp :=[]int{1,2,3}
	elem := 4
	arr := []int{1,2,3,4}
	elemRes, arrRes:=pop(arr)
	for i:=0;i<len(exp);i++{
		if arrRes[i]!=exp[i]{
			t.Error("expected", exp, "result",arrRes)
		}
	}
	if elem != elemRes{
		t.Error("expected", elem, "result",elemRes)
	}
	fmt.Println("Test_pop -- OK")
}

func Test_action1(t *testing.T){
	arr := []int{1,2,3,4}
	res, err := action(arr,"+")
	exp := []int{1,2,7}
	for i:=0;i<len(exp);i++{
		if res[i]!=exp[i]{
			t.Error("expected", exp, "result",res)
		}
	}
	if err != nil{
		t.Error("expected", nil, "result",err)
	}
}
func Test_action2(t *testing.T){
	arr := []int{1,2,3,4}
	res, err := action(arr,"-")
	exp := []int{1,2,-1}
	for i:=0;i<len(exp);i++{
		if res[i]!=exp[i]{
			t.Error("expected", exp, "result",res)
		}
	}
	if err != nil{
		t.Error("expected", nil, "result",err)
	}
}
func Test_action3(t *testing.T){
	arr := []int{1,2,3,4}
	res, err := action(arr,"*")
	exp := []int{1,2,12}
	for i:=0;i<len(exp);i++{
		if res[i]!=exp[i]{
			t.Error("expected", exp, "result",res)
		}
	}
	if err != nil{
		t.Error("expected", nil, "result",err)
	}
}
func Test_action4(t *testing.T){
	arr := []int{1,2,4,2}
	res, err := action(arr,"/")
	exp := []int{1,2,2}
	for i:=0;i<len(exp);i++{
		if res[i]!=exp[i]{
			t.Error("expected", exp, "result",res)
		}
	}
	if err != nil{
		t.Error("expected", nil, "result",err)
	}
}

func Test_parseInfix1(t *testing.T) {
	str := []string{"(","8","-","1",")","*","5"}
	res := parseInfix(str)
	exp := "8 1 - 5 *"
	if exp != res{
		t.Error("expected", exp, "result",res)
	}

}
func Test_parseInfix2(t *testing.T) {
	str := []string{"8","-","1","*","5"}
	res := parseInfix(str)
	exp := "8 1 5 * -"
	if exp != res{
		t.Error("expected", exp, "result",res)
	}

}

func Test_calc1(t *testing.T) {
	str := []string{"8","1","-","5","*"}
	res, err := calc(str)
	exp := float64(35)
	if exp != res || err != nil{
		t.Error("expected", exp, "result",res)
	}
}
func Test_calc2(t *testing.T) {
	str := []string{"8","1","5","*","-"}
	res, err := calc(str)
	exp := float64(3)
	if exp != res || err != nil{
		t.Error("expected", exp, "result",res)
	}
}

