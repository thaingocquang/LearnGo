package MyPackage

import "fmt"
import "runtime"

func Version() {
	fmt.Println("Check")
	fmt.Println(runtime.Version())
}
