package main

import "fmt"

type ISay interface{
	Say(string)

}


type People struct {
	ISay
	name string
}

func (p *People) Say(s string) {
	fmt.Printf("Hello %s",s)	
}


func main(){
	pp := People{name: "Nguuyen"}

	var d ISay
	d = &pp
	d.Say("Dmm")

	fmt.Printf("Hello %s",pp.name)
}