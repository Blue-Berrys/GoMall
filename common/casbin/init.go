package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MYSQL_ROOT_PASSWORD = "root" // 设置 root 用户的密码
	MYSQL_USER          = "root" // 新用户
	MYSQL_PORT          = "3307"
	MYSQL_ADDR          = "localhost"
	E                   *casbin.Enforcer

	CustomerGroup = "customer"
	MerchantGroup = "merchant"
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", MYSQL_USER, MYSQL_ROOT_PASSWORD, MYSQL_ADDR, MYSQL_PORT)
	a, err := gormadapter.NewAdapter("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	E, err = casbin.NewEnforcer("./model.conf", a)
	InitRolesAndPermissions()
	//sub := "alice"
	//obj := "data1"
	//act := "read"
	//e.AddFunction("my_func", KeyMatchFunc)
	//added, err := e.AddGroupingPolicy("alice", "admin_data1")
	//added2, err2 := e.AddPolicy("admin_data1", "data1", "read")
	//ok, err := e.Enforce(sub, obj, act)

	//fmt.Println(added, err)
	//fmt.Println(added2, err2)
	//if err != nil {
	//	// handle err
	//	fmt.Println(err)
	//}
	//if ok == true {
	//	//permit alice to read data1
	//	fmt.Println("ok")
	//} else {
	//	// deny the request, show an error
	//	fmt.Println("not ok")
	//}
	if err != nil {
		panic(err)
	}
}

func AddPolicy(sub, obj, act string) (bool, error) {
	f, err := E.AddPolicy(sub, obj, act)
	E.LoadPolicy()
	return f, err
}

func CheckPermission(sub, obj, act string) (bool, error) {
	return E.Enforce(sub, obj, act)
}
func AddGroupPolicy(group, object, act string) (bool, error) {
	f, err := E.AddPolicy(group, object, act)
	E.LoadPolicy()
	return f, err
}
func AddUserToGroup(user, group string) (bool, error) {
	f, err := E.AddGroupingPolicy(user, group)
	E.LoadPolicy()
	return f, err
}

func InitRolesAndPermissions() {
	// 为顾客组分配权限
	AddGroupPolicy(CustomerGroup, "order", "read") // 顾客可以查看订单

	// 为商家组分配权限
	AddGroupPolicy(MerchantGroup, "order", "read") // 商家可以查看订单

	// 确保策略加载
	E.LoadPolicy()
}
