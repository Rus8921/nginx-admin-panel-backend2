package Location

//
//import (
//	"gorm.io/gorm"
//)
//
//// Location represents the location model
//type Location struct {
//	gorm.Model
//	body string
//	//Site []*models.Site `gorm:"many2many:location_site;"` // Many-to-many relationship with Site
//}
//
//// CreateLocation creates a new location in the database
//// db: Database connection
//// location: Location to be created
//// Returns an error if the creation fails
//func CreateLocation(db *gorm.DB, location *Location) error {
//	result := db.Create(location)
//	return result.Error
//}
//
//// GetLocation retrieves a location by ID
//// db: Database connection
//// id: Location ID
//// Returns the location or an error if retrieval fails
//func GetLocation(db *gorm.DB, id uint) (Location, error) {
//	var location Location
//	if err := db.Where("id = ?", id).First(&location).Error; err != nil {
//		return Location{}, err
//	}
//	return location, nil
//}
//
//// DeleteLocation deletes a location by ID
//// db: Database connection
//// id: Location ID
//// Returns an error if deletion fails or if the location is not found
//func DeleteLocation(db *gorm.DB, id uint) error {
//	result := db.Delete(&Location{}, id)
//	if result.Error != nil {
//		return result.Error
//	}
//	if result.RowsAffected == 0 {
//		return nil
//	}
//	return nil
//}
//
//func UpdateLocation(db *gorm.DB, id uint, updatedLocation Location) (Location, error) {
//	var location Location
//	if err := db.Where("id = ?", id).First(&location).Error; err != nil {
//		return Location{}, err
//	}
//	location.body = updatedLocation.body
//	result := db.Save(&location)
//	return location, result.Error
//}
