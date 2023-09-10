// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserExtends is the golang structure of table user_extends for DAO operations like Where/Data.
type UserExtends struct {
	g.Meta     `orm:"table:user_extends, do:true"`
	Id         interface{} //
	UserId     interface{} //
	ViewNums   interface{} // 浏览次数
	FansNums   interface{} // 粉丝数
	FollowNums interface{} // 关注人数
	AdminRole  interface{} // 后台角色，用户可自由切换
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
