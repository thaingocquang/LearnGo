package main

import renameImport1 "fmt"
import renameImport2 "fmt"
import renameImport3 "fmt"

func main() {
	renameImport1.Println("hello")
	renameImport2.Println("hey")
	renameImport3.Println("hi")
}
