// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostDao is the data access object for table post.
type PostDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns PostColumns // columns contains all the column names of Table for convenient usage.
}

// PostColumns defines and stores column names for table post.
type PostColumns struct {
	Id        string //
	Content   string //
	UserId    string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// postColumns holds the columns for table post.
var postColumns = PostColumns{
	Id:        "id",
	Content:   "content",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPostDao creates and returns a new DAO object for table data access.
func NewPostDao() *PostDao {
	return &PostDao{
		group:   "default",
		table:   "post",
		columns: postColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PostDao) Columns() PostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
