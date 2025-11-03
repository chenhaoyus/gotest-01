package pkg2

import (
	"fmt"
)

const Pkgname3 string = "pkg3"

var PkgNameVar3 string = getPkgName3()

func getPkgName3() string {
	fmt.Println("pkg3.PkgNameVar has been initialized")
	return Pkgname3
}

func init() {
	fmt.Println("pkg3 init method invoked")
}
