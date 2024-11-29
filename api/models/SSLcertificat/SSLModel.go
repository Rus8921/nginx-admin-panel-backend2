package SSLcertificat

import (
	"fmt"
	"gorm.io/gorm"
)

// SSL represents an SSL certificate model with fields for the certificate file, key file,
// active status, and associated site ID.
type SSL struct {
	gorm.Model
	FileCrt  string // Path to the certificate file
	FileKey  string // Path to the key file
	IsActive bool   // Indicates if the SSL certificate is active
	SiteID   uint   // ID of the site associated with the SSL certificate
}

// CreateSSL inserts a new SSL certificate record into the database.
// Parameters:
// - db: A pointer to the gorm.DB instance.
// - ssl: A pointer to the SSL struct to be created.
// Returns:
// - error: An error object if the operation fails, otherwise nil.
func CreateSSL(db *gorm.DB, ssl *SSL) error {
	ssl.IsActive = false
	result := db.Create(ssl)
	return result.Error
}

// GetSSL retrieves an SSL certificate record from the database by its ID.
// Parameters:
// - db: A pointer to the gorm.DB instance.
// - id: The ID of the SSL certificate to retrieve.
// Returns:
// - SSL: The retrieved SSL struct.
// - error: An error object if the operation fails, otherwise nil.
func GetSSL(db *gorm.DB, id uint) (SSL, error) {
	var ssl SSL
	if err := db.Where("id = ?", id).First(&ssl).Error; err != nil {
		return SSL{}, err
	}
	return ssl, nil
}

func GetSSLCertificatesAll(db *gorm.DB) ([]SSL, error) {
	var ssls []SSL
	if err := db.Find(&ssls).Error; err != nil {
		return nil, err
	}
	return ssls, nil
}

// DeleteSSL deletes an SSL certificate record from the database by its ID.
// Parameters:
// - db: A pointer to the gorm.DB instance.
// - id: The ID of the SSL certificate to delete.
// Returns:
// - error: An error object if the operation fails, otherwise nil.
func DeleteSSL(db *gorm.DB, id uint) error {
	result := db.Delete(&SSL{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting ssl: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("ssl with ID %d not found", id)
	}
	return nil
}

// UpdateSSL updates an existing SSL certificate record in the database.
// Parameters:
// - db: A pointer to the gorm.DB instance.
// - id: The ID of the SSL certificate to update.
// - updatedSSL: The SSL struct containing updated fields.
// Returns:
// - SSL: The updated SSL struct.
// - error: An error object if the operation fails, otherwise nil.
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

func ActivateOrUnactivateSSL(db *gorm.DB, id uint) error {
	var ssl SSL
	ssl, err := GetSSL(db, id)
	if err != nil {
		return err
	}
	if ssl.IsActive {
		ssl.IsActive = false
		if err := db.Save(&ssl).Error; err != nil {
			return fmt.Errorf("error saving ssl changes: %w", err)
		}
	} else {
		ssl.IsActive = true
		if err := db.Save(&ssl).Error; err != nil {
			return fmt.Errorf("error saving ssl changes: %w", err)
		}
	}
	return nil
}
