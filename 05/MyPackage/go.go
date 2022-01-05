package MyPackage

import "fmt"
import "runtime"

func Version() {
	fmt.Println(runtime.Version())
}
