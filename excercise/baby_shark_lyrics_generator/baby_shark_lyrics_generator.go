package main

import (
	"fmt"
	//. "strings"
)

func main(){
	res := BabySharkLyrics()
	fmt.Print(res)
}

//improved my code little bit
func BabySharkLyrics()(r string){
	for _,v:=range []string{"Baby shark","Mommy shark","Daddy shark","Grandma shark","Grandpa shark","Let's go hunt"}{
		s := v+", doo doo doo doo doo doo\n"
		r+=s+s+s+v+"!\n"
	}
	r+="Run away,…\n"
	return 
}

//my solution that i have submitted
// func BabySharkLyrics()string{
// 	n:=[]string{"Baby shark","Mommy shark","Daddy shark","Grandma shark","Grandpa shark","Let's go hunt"}
// 	d,l,r,x:=", doo doo doo doo doo doo","Run away,…","","\n"
// 	for _,v:=range n{
// 		for range "mah"{
// 			r+=v+d+x
// 		}
// 		r+=v+"!"+x
// 	}
// 	return r+l+""+x
// }

//some interesting solution from others 
//contributor : corebreaker
// func BabySharkLyrics() (r string) {
// 	d := ", doo doo doo doo doo doo\n"
// 	a := func(p string) {r += p+d+p+d+p+d+p+"!\n"}
// 	b := func(p string){a(p+" shark")}
// 	b("Baby");b("Mommy");b("Daddy");b("Grandma");b("Grandpa");a("Let's go hunt")
// 	return r+"Run away,…\n"
// }

//contributor : davidamey
// func BabySharkLyrics() (r string) {
// 	for _, v := range []string{"Baby shark", "Mommy shark", "Daddy shark", "Grandma shark", "Grandpa shark", "Let's go hunt"} {
// 		r += Repeat(v + "," + Repeat(" doo", 6) + "\n", 3) + v + "!\n"
// 	}
// 	return r + "Run away,…\n"
// }

//contributor : Apomelitos
// func BabySharkLyrics() (s string) {
// 	for i,name:=range []string{"Baby","Mommy","Daddy","Grandma","Grandpa","Let's go hunt"} {
// 		gr:=name+" shark"
// 		if i==5 {gr=name}
// 		t:=gr+", doo doo doo doo doo doo\n"
// 		s+=t+t+t+gr+"!\n"
// 	}
// 	return s+"Run away,…\n"
// }



