// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Post is the golang structure of table post for DAO operations like Where/Data.
type Post struct {
	g.Meta    `orm:"table:post, do:true"`
	Id        interface{} //
	Content   interface{} //
	UserId    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
