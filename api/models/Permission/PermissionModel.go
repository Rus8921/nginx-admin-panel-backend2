package Permission

import "gorm.io/gorm"

// Permission represents the permission model with fields for
// ApproveAdminID, UserID, and SiteID. It embeds gorm.Model for
// common model fields.
type Permission struct {
	gorm.Model
	ApproveAdminID uint // 1 к 1
	UserID         uint // 1 к 1
	SiteID         uint // 1 к 1
}

// CreatePermission creates a new permission record in the database.
// It takes a gorm.DB instance and a pointer to a Permission struct.
// Returns an error if the creation fails.
func CreatePermission(db *gorm.DB, permission *Permission) error {
	result := db.Create(permission)
	return result.Error
}

// GetPermission retrieves a permission record by its ID from the database.
// It takes a gorm.DB instance and the ID of the permission to retrieve.
// Returns the Permission struct and an error if the retrieval fails.
func GetPermission(db *gorm.DB, id uint) (Permission, error) {
	var permission Permission
	if err := db.Where("id = ?", id).First(&permission).Error; err != nil {
		return Permission{}, err
	}
	return permission, nil
}

// DeletePermission deletes a permission record by its ID from the database.
// It takes a gorm.DB instance and the ID of the permission to delete.
// Returns an error if the deletion fails.
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

func GetPermissionAll(db *gorm.DB) ([]Permission, error) {
	var permissions []Permission
	if err := db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func CheckPermission(db *gorm.DB, userID uint, siteID uint) (bool, error) {
	var permission Permission
	if err := db.Where("user_id = ? AND site_id = ?", userID, siteID).First(&permission).Error; err != nil {
		return false, err
	}
	return true, nil
}
