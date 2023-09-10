// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionsDao is the data access object for table permissions.
type PermissionsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PermissionsColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionsColumns defines and stores column names for table permissions.
type PermissionsColumns struct {
	Id         string //
	Name       string //
	GuardName  string //
	Path       string // 对应 vue path
	Icon       string // 对应 vue icon
	Title      string // 对应 vue title
	Hidden     string // 对应 vue 是否不显示在侧边栏中,详情，编辑等页面会用到
	KeepAlive  string // 对应 vue 是否使用 keep-alive
	AlwaysShow string // 当子菜单只有一个时，是否要显示父菜单
	Component  string // 对应 vue component，为空时取组合 path name 对应的组件
	Link       string // 如果有值，则另外打开窗口弹窗
	Iframe     string // 如果有值，则展现 iframe
	Redirect   string // 对应 vue 面包屑重定向,为空的时候父级指向子集第一个
	ParentId   string //
	Order      string // 排序,同层次排序才有意义
	CreatedAt  string //
	UpdatedAt  string //
}

// permissionsColumns holds the columns for table permissions.
var permissionsColumns = PermissionsColumns{
	Id:         "id",
	Name:       "name",
	GuardName:  "guard_name",
	Path:       "path",
	Icon:       "icon",
	Title:      "title",
	Hidden:     "hidden",
	KeepAlive:  "keep_alive",
	AlwaysShow: "always_show",
	Component:  "component",
	Link:       "link",
	Iframe:     "iframe",
	Redirect:   "redirect",
	ParentId:   "parent_id",
	Order:      "order",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewPermissionsDao creates and returns a new DAO object for table data access.
func NewPermissionsDao() *PermissionsDao {
	return &PermissionsDao{
		group:   "default",
		table:   "permissions",
		columns: permissionsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionsDao) Columns() PermissionsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PermissionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
