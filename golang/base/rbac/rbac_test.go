package rbac

import "fmt"

// 定义角色类型
type Role string

const (
	Admin  Role = "Admin"
	Editor Role = "Editor"
	Viewer Role = "Viewer"
)

// 用户结构体
type User struct {
	Username string
	Roles    []Role
}

// RBAC系统结构体
type RBAC struct {
	rolePermissions map[Role][]string
}

// 初始化RBAC系统
func NewRBAC() *RBAC {
	rbac := &RBAC{
		rolePermissions: make(map[Role][]string),
	}
	return rbac
}

// 添加角色的权限
func (rbac *RBAC) AddRolePermission(role Role, permissions []string) {
	rbac.rolePermissions[role] = permissions
}

// 检查用户是否有权限
func (rbac *RBAC) HasPermission(user User, permission string) bool {
	for _, role := range user.Roles {
		permissions, found := rbac.rolePermissions[role]
		if !found {
			continue
		}

		for _, p := range permissions {
			if p == permission {
				return true
			}
		}
	}

	return false
}

func main() {
	// 初始化RBAC系统
	rbac := NewRBAC()

	// 添加角色权限
	rbac.AddRolePermission(Admin, []string{"create", "read", "update", "delete"})
	rbac.AddRolePermission(Editor, []string{"create", "read", "update"})
	rbac.AddRolePermission(Viewer, []string{"read"})

	// 创建几个用户并赋予不同角色
	user1 := User{Username: "user1", Roles: []Role{Admin}}
	user2 := User{Username: "user2", Roles: []Role{Editor}}
	user3 := User{Username: "user3", Roles: []Role{Viewer}}

	// 检查用户权限
	fmt.Println(user1.Username, "has 'create' permission:", rbac.HasPermission(user1, "create"))
	fmt.Println(user2.Username, "has 'update' permission:", rbac.HasPermission(user2, "update"))
	fmt.Println(user3.Username, "has 'delete' permission:", rbac.HasPermission(user3, "delete"))
}
