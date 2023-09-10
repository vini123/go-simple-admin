// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserExtends is the golang structure for table user_extends.
type UserExtends struct {
	Id         uint64      `json:"id"         ` //
	UserId     uint64      `json:"userId"     ` //
	ViewNums   uint        `json:"viewNums"   ` // 浏览次数
	FansNums   uint        `json:"fansNums"   ` // 粉丝数
	FollowNums uint        `json:"followNums" ` // 关注人数
	AdminRole  string      `json:"adminRole"  ` // 后台角色，用户可自由切换
	CreatedAt  *gtime.Time `json:"createdAt"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` //
}
