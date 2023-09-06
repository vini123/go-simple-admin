// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"goSimpleAdmin/api/user"
)

type ControllerAdmin struct{}

func NewAdmin() user.IUserAdmin {
	return &ControllerAdmin{}
}

