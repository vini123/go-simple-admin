// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	Id        uint64      `json:"id"         ` //
	Name      string      `json:"name"       ` //
	GuardName string      `json:"guard_name" ` //
	Title     string      `json:"title"      ` // 角色名称
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}
