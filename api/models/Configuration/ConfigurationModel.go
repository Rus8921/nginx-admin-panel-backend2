package Configuration

import (
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Upstreams"
	"gorm.io/gorm"
)

// Configuration represents the configuration model
type Configuration struct {
	gorm.Model
	Parametrs string               `gorm:"not null"` // Configuration parameters
	SiteID    uint                 // One-to-one relationship with Site
	Upstream  []Upstreams.Upstream // One-to-many relationship with Upstreams
}

// CreateConfiguration creates a new configuration in the database
// db: Database connection
// configuration: Configuration to be created
// Returns an error if the creation fails
func CreateConfiguration(db *gorm.DB, configuration *Configuration) error {
	result := db.Create(configuration)
	return result.Error
}

// GetConfiguration retrieves a configuration by ID
// db: Database connection
// id: Configuration ID
// Returns the configuration or an error if retrieval fails
func GetConfiguration(db *gorm.DB, id uint) (Configuration, error) {
	var configuration Configuration
	if err := db.Where("id = ?", id).First(&configuration).Error; err != nil {
		return Configuration{}, err
	}
	return configuration, nil
}

func GetConfigurationAll(db *gorm.DB) ([]Configuration, error) {
	var configurations []Configuration
	if err := db.Find(&configurations).Error; err != nil {
		return nil, err
	}
	return configurations, nil
}

// DeleteConfiguration deletes a configuration by ID
// db: Database connection
// id: Configuration ID
// Returns an error if deletion fails or if the configuration is not found
func DeleteConfiguration(db *gorm.DB, id uint) error {
	result := db.Delete(&Configuration{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

// UpdateConfiguration updates a configuration's details
// db: Database connection
// id: Configuration ID
// updatedConfiguration: Configuration with updated details
// Returns the updated configuration or an error if the update fails
func UpdateConfiguration(db *gorm.DB, id uint, updatedConfiguration Configuration) (Configuration, error) {
	var configuration Configuration
	configuration, err := GetConfiguration(db, id)
	if err != nil {
		return Configuration{}, err
	}
	if updatedConfiguration.Parametrs != "" {
		configuration.Parametrs = updatedConfiguration.Parametrs
	}
	if err := db.Save(&configuration).Error; err != nil {
		return Configuration{}, err
	}
	return configuration, nil
}
