// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id              uint64      `json:"id"              ` //
	Account        string      `json:"account"        ` //
	Phone           string      `json:"phone"           ` //
	Email           string      `json:"email"           ` //
	Password        string      `json:"password"        ` //
	Nickname        string      `json:"nickname"        ` //
	Avatar          string      `json:"avatar"          ` // 头像
	Gender          uint        `json:"gender"          ` // 性别 0 未知 1 男 2 女
	Signature       string      `json:"signature"       ` // 签名
	EmailVerifiedAt *gtime.Time `json:"emailVerifiedAt" ` //
	RememberToken   string      `json:"rememberToken"   ` //
	CreatedAt       *gtime.Time `json:"createdAt"       ` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` //
}
