package pkg1

import (
	"fmt"

	_ "github.com/init_order/pkg2"
)

const Pkaname string = "pkg1"

var PkgName string = getPkgName()

func getPkgName() string {
	fmt.Println("Initializing pkg1")
	return Pkaname
}
func init() {
	fmt.Println("pkg1 init method invoked")
}
