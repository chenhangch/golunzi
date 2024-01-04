package package2

import (
	"fmt"

	"github.com/chenhangch/golunzi/test/golang.init/utils"
)

var Value1 = utils.TraceLog("init package2 value1", 20)
var Value2 = utils.TraceLog("init package2 value2", 30)

func init() {
	fmt.Println("init func1 in package2")
}

func init() {
	fmt.Println("init func2 in package2")
}