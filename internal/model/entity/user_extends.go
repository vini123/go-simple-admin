// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserExtends is the golang structure for table user_extends.
type UserExtends struct {
	Id         uint64      `json:"id"          ` //
	UserId     uint64      `json:"user_id"     ` //
	ViewNums   uint        `json:"view_nums"   ` // 浏览次数
	FansNums   uint        `json:"fans_nums"   ` // 粉丝数
	FollowNums uint        `json:"follow_nums" ` // 关注人数
	AdminRole  string      `json:"admin_role"  ` // 后台角色，用户可自由切换
	CreatedAt  *gtime.Time `json:"created_at"  ` //
	UpdatedAt  *gtime.Time `json:"updated_at"  ` //
}
