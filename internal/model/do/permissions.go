// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure of table permissions for DAO operations like Where/Data.
type Permissions struct {
	g.Meta     `orm:"table:permissions, do:true"`
	Id         interface{} //
	Name       interface{} //
	GuardName  interface{} //
	Path       interface{} // 对应 vue path
	Icon       interface{} // 对应 vue icon
	Title      interface{} // 对应 vue title
	Hidden     interface{} // 对应 vue 是否不显示在侧边栏中,详情，编辑等页面会用到
	KeepAlive  interface{} // 对应 vue 是否使用 keep-alive
	AlwaysShow interface{} // 当子菜单只有一个时，是否要显示父菜单
	Component  interface{} // 对应 vue component，为空时取组合 path name 对应的组件
	Link       interface{} // 如果有值，则另外打开窗口弹窗
	Iframe     interface{} // 如果有值，则展现 iframe
	Redirect   interface{} // 对应 vue 面包屑重定向,为空的时候父级指向子集第一个
	ParentId   interface{} //
	Order      interface{} // 排序,同层次排序才有意义
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
