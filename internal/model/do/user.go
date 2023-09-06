// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta          `orm:"table:user, do:true"`
	Id              interface{} //
	Passport        interface{} //
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
