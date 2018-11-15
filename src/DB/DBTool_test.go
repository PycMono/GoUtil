package DBTool

import (
	"testing"
	"fmt"
)

func TestConn(t *testing.T) {
	fmt.Println(fmt.Sprintf("TestConn",GetDB()))
}