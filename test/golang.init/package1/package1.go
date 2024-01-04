package package1

import (
	"fmt"

	"github.com/chenhangch/golunzi/test/golang.init/package2"
	"github.com/chenhangch/golunzi/test/golang.init/utils"
)

var V1 = utils.TraceLog("init package1 value", package2.Value1+10)
var V2 = utils.TraceLog("init package1 value", package2.Value2+10)

func init() {
	fmt.Println("init func in package1")
}
