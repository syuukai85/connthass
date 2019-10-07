package seed

import (
	"github.com/connthass/connthass/api/entity"
	"github.com/connthass/connthass/api/infrastructure/orm/model"
	"github.com/connthass/connthass/api/interface/gateway"
	"github.com/jinzhu/gorm"
)

// SysRoles シードデータ配列
type SysRoles []model.SysRole

// NewSysRole コンストラクタ
func NewSysRole() SysRoles {
	return SysRoles{
		{
			model.Base{
				ID: gateway.EntitySysRoleIDToUint(entity.SystemAdminID),
			},
			model.BaseRole{
				Name: "systemAdmin",
			},
		},
		{
			model.Base{
				ID: gateway.EntitySysRoleIDToUint(entity.GeneralUserID),
			},
			model.BaseRole{
				Name: "generalUser",
			},
		},
	}
}

func (ar SysRoles) seed(db *gorm.DB) {
	for _, sysRole := range ar {
		if !db.First(&model.SysRole{}, *sysRole.ID).RecordNotFound() {
			continue
		}
		db.Create(&sysRole)
	}
}
