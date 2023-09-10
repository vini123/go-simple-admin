// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserExtendsDao is the data access object for table user_extends.
type UserExtendsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns UserExtendsColumns // columns contains all the column names of Table for convenient usage.
}

// UserExtendsColumns defines and stores column names for table user_extends.
type UserExtendsColumns struct {
	Id         string //
	UserId     string //
	ViewNums   string // 浏览次数
	FansNums   string // 粉丝数
	FollowNums string // 关注人数
	AdminRole  string // 后台角色，用户可自由切换
	CreatedAt  string //
	UpdatedAt  string //
}

// userExtendsColumns holds the columns for table user_extends.
var userExtendsColumns = UserExtendsColumns{
	Id:         "id",
	UserId:     "user_id",
	ViewNums:   "view_nums",
	FansNums:   "fans_nums",
	FollowNums: "follow_nums",
	AdminRole:  "admin_role",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewUserExtendsDao creates and returns a new DAO object for table data access.
func NewUserExtendsDao() *UserExtendsDao {
	return &UserExtendsDao{
		group:   "default",
		table:   "user_extends",
		columns: userExtendsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserExtendsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserExtendsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserExtendsDao) Columns() UserExtendsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserExtendsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserExtendsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserExtendsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
