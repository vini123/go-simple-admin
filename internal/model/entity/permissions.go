// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure for table permissions.
type Permissions struct {
	Id         uint64      `json:"id"          ` //
	Name       string      `json:"name"        ` //
	GuardName  string      `json:"guard_name"  ` //
	Path       string      `json:"path"        ` // 对应 vue path
	Icon       string      `json:"icon"        ` // 对应 vue icon
	Title      string      `json:"title"       ` // 对应 vue title
	Hidden     int         `json:"hidden"      ` // 对应 vue 是否不显示在侧边栏中,详情，编辑等页面会用到
	KeepAlive  int         `json:"keep_alive"  ` // 对应 vue 是否使用 keep-alive
	AlwaysShow int         `json:"always_show" ` // 当子菜单只有一个时，是否要显示父菜单
	Component  string      `json:"component"   ` // 对应 vue component，为空时取组合 path name 对应的组件
	Link       string      `json:"link"        ` // 如果有值，则另外打开窗口弹窗
	Iframe     string      `json:"iframe"      ` // 如果有值，则展现 iframe
	Redirect   string      `json:"redirect"    ` // 对应 vue 面包屑重定向,为空的时候父级指向子集第一个
	ParentId   uint64      `json:"parent_id"   ` //
	Order      uint        `json:"order"       ` // 排序,同层次排序才有意义
	CreatedAt  *gtime.Time `json:"created_at"  ` //
	UpdatedAt  *gtime.Time `json:"updated_at"  ` //
}
