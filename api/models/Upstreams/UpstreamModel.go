package Upstreams

import "gorm.io/gorm"

// Upstream represents the upstream model with GORM's built-in fields and custom fields.
type Upstream struct { // подвязать к локейшен, а локейшен к конфигурации, а конфигу к сайту. У многих локейшенов может быть 1 апстрим
	// пускай есть дефолтный апстрим у сайта, к сайту мы можем привзяывать локейшены, которые будут наследовать апстримы сайта, но можно его изменить
	gorm.Model
	ConfigurationID uint   // ID of the configuration
	Priority        uint   `gorm:"not nul"` // Priority of the upstream
	Parametr        string `gorm:"not nul"` // Parameter of the upstream
}

// CreateUpstream creates a new upstream record in the database.
// Parameters:
// - db: The GORM database connection.
// - upstream: The upstream object to be created.
// Returns:
// - error: An error object if any error occurs, otherwise nil.
func CreateUpstream(db *gorm.DB, upstream *Upstream) error {
	result := db.Create(upstream)
	return result.Error
}

// GetUpstream retrieves an upstream record by its ID.
// Parameters:
// - db: The GORM database connection.
// - id: The ID of the upstream to retrieve.
// Returns:
// - Upstream: The retrieved upstream object.
// - error: An error object if any error occurs, otherwise nil.
func GetUpstream(db *gorm.DB, id uint) (Upstream, error) {
	var upstream Upstream
	if err := db.Where("id = ?", id).First(&upstream).Error; err != nil {
		return Upstream{}, err
	}
	return upstream, nil
}

// GetUpstreamsAll retrieves all upstream records from the database.
// Parameters:
// - db: The GORM database connection.
// Returns:
// - []Upstream: A slice of all upstream objects.
// - error: An error object if any error occurs, otherwise nil.
func GetUpstreamsAll(db *gorm.DB) ([]Upstream, error) {
	var upstreams []Upstream
	if err := db.Find(&upstreams).Error; err != nil {
		return nil, err
	}
	return upstreams, nil
}

// DeleteUpstream deletes an upstream record by its ID.
// Parameters:
// - db: The GORM database connection.
// - id: The ID of the upstream to delete.
// Returns:
// - error: An error object if any error occurs, otherwise nil.
func DeleteUpstream(db *gorm.DB, id uint) error {
	result := db.Delete(&Upstream{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

// UpdateUpstream updates an existing upstream record by its ID.
// Parameters:
// - db: The GORM database connection.
// - id: The ID of the upstream to update.
// - updatedUpstream: The updated upstream object with new values.
// Returns:
// - Upstream: The updated upstream object.
// - error: An error object if any error occurs, otherwise nil.
func UpdateUpstream(db *gorm.DB, id uint, updatedUpstream Upstream) (Upstream, error) {
	var upstream Upstream
	upstream, err := GetUpstream(db, id)
	if err != nil {
		return Upstream{}, err
	}
	if updatedUpstream.Parametr != "" {
		upstream.Parametr = updatedUpstream.Parametr
	}
	if updatedUpstream.Priority != 0 {
		upstream.Priority = updatedUpstream.Priority
	}
	if err := db.Save(&upstream).Error; err != nil {
		return Upstream{}, err
	}
	return upstream, nil
}
