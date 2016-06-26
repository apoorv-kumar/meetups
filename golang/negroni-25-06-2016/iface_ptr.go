package main

import("fmt")

type Item interface{
    GetId() string
}
//------------------

type Application struct {
    id string
}

func (a Application) GetId() string {
    return a.id
}
//-------------------

type Person struct {
    name string
}

func (p *Person) GetId() string {
    return p.name
}
//------------------

func printItemId(i Item){
    fmt.Println(i.GetId())
}

func (a *Application) reset(){
    a.id = "reset value"
}


func main(){
    p := Person{name: "Apoorv"}
    a := Application{id: "VS Code"}

   printItemId(&p)
   //printItemId(p)
   
   printItemId(&a)
   printItemId(a)

   a.reset()
   printItemId(a)
}