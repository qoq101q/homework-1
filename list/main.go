package main

import (
	"fmt"
	"list/storages/list"
	"list/storages/model"
	"list/storages/slice"
)

func main() {
	l := list.NewList()
	model.Add(l, 15)
	model.Add(l, 9)
	model.Add(l, -2)
	model.Add(l, 35)
	model.Add(l, 7)
	model.Add(l, 1)
	//model.Delete(l, 3)
	//model.Print(l)
	fmt.Println("Sort List")
	model.SortInc(l)
	//l.SortIncLink()
	//model.SortDec(l)
	//l.SortDecLink()
	//model.Print(l)

	s := slice.Slice{}
	model.Add(&s, 15)
	model.Add(&s, 9)
	model.Add(&s, -2)
	model.Add(&s, 35)
	model.Add(&s, 7)
	model.Add(&s, 1)
	model.Print(&s)
	fmt.Println("Sort List")
	//model.SortInc(&s)
	model.SortDec(&s)
	s.Print()
}
