// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta          `orm:"table:users, do:true"`
	Id              interface{} //
	Account         interface{} //
	Phone           interface{} //
	Email           interface{} //
	Password        interface{} //
	Nickname        interface{} //
	Avatar          interface{} // 头像
	Gender          interface{} // 性别 0 未知 1 男 2 女
	Signature       interface{} // 签名
	EmailVerifiedAt *gtime.Time //
	RememberToken   interface{} //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
