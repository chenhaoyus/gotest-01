package pkg2

import (
	"fmt"
)

const Pkgname string = "pkg2"

var PkgNameVar string = getPkgName()

func getPkgName() string {
	fmt.Println("pkg2.PkgNameVar has been initialized")
	return Pkgname
}

func init() {
	fmt.Println("pkg2 init method invoked")
}
