package main


interface {
	Method1()
	Method2() err
}

type IncludeSubTask Task

func (t IncludeSubTask) indentedString(Prefix string) string {
	str := Prefix
}