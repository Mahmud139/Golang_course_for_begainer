package main

func main(){
	one()
}

func one(){
	two()
}

func two(){
	three()
}

func three(){
	panic("this call stack's too deep for me!")
}