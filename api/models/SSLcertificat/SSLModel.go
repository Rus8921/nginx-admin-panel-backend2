package SSLcertificat

import "gorm.io/gorm"

type SSL struct {
	gorm.Model
	FileCrt  string
	FileKey  string
	IsActive bool
	SiteID   uint
}

func CreateSSL(db *gorm.DB, ssl *SSL) error {
	result := db.Create(ssl)
	return result.Error
}

func GetSSL(db *gorm.DB, id uint) (SSL, error) {
	var ssl SSL
	if err := db.Where("id = ?", id).First(&ssl).Error; err != nil {
		return SSL{}, err
	}
	return ssl, nil
}

func DeleteSSL(db *gorm.DB, id uint) error {
	result := db.Delete(&SSL{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func UpdateSSL(db *gorm.DB, id uint, updatedSSL SSL) (SSL, error) {
	var ssl SSL
	ssl, err := GetSSL(db, id)
	if err != nil {
		return SSL{}, err
	}
	if updatedSSL.FileCrt != "" {
		ssl.FileCrt = updatedSSL.FileCrt
	}
	if updatedSSL.FileKey != "" {
		ssl.FileKey = updatedSSL.FileKey
	}
	if err := db.Save(&ssl).Error; err != nil {
		return SSL{}, err
	}
	return ssl, nil
}
