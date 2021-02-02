package main
import "fmt"

var global string = " hey global scope " // global scope
func main(){
	var local int = 3
	fmt.Println("Global scope variable:",global)
	fmt.Println("Local variable:", 	local)
	sayhello()
}
func sayhello(){
	//here i can refer to global, but not to local
	// i can say local is something else
	var local bool = true;
	fmt.Println("Here local is something else:", local)
	fmt.Println("But global is still:", global)
 	return 0
}