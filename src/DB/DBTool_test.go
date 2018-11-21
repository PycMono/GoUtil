package DBTool

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	tempList := make([]string, 0, 10)
	tempList = append(tempList, "1")
	tempList = append(tempList, "2")
	fmt.Println(tempList)
	fmt.Println(fmt.Sprintf("TestConn", GetDB()))
}
