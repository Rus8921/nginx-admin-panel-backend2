package Permission

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ApproveAdminID uint
	UserID         uint
	SiteID         uint
}

func CreatePermission(db *gorm.DB, permission *Permission) error {
	result := db.Create(permission)
	return result.Error
}

func GetPermission(db *gorm.DB, id uint) (Permission, error) {
	var permission Permission
	if err := db.Where("id = ?", id).First(&permission).Error; err != nil {
		return Permission{}, err
	}
	return permission, nil
}

func DeletePermission(db *gorm.DB, id uint) error {
	result := db.Delete(&Permission{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}
