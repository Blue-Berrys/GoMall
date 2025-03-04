package casbin

import (
	"fmt"
	"testing"
)

func TestInit_Run(t *testing.T) {
	Init()
	ok, err := AddUserToGroup("alice", CustomerGroup)
	fmt.Println(ok, err)
	// 确保策略加载
	E.LoadPolicy()
	ok, err = CheckPermission("alice", "order", "read")
	fmt.Println(ok, err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)
}
