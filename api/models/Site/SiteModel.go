package models

import (
	"fmt"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Configuration"

	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/SSLcertificat"
	"gorm.io/gorm"
)

// Site represents the site model with fields for site name, domain,
// active status, and associated Nginx server ID.
type Site struct {
	gorm.Model
	SiteName       string                      `gorm:"unique;not null"` // Unique and non-null site name
	Domain         string                      `gorm:"unique;not null"` // Unique and non-null domain
	IsActive       bool                        // Active status of the site
	NginxServerID  uint                        // Associated Nginx server ID
	Configurations Configuration.Configuration // one-to-one relationship
	Permission     Permission.Permission       // one-to-one relationship
	Location       []*Location                 `gorm:"many2many:location_site;"` // many-to-many relationship
	SSlcerificate  []SSLcertificat.SSL         // multiple certificates for one site
}

// CreateSite creates a new site in the database with the given site details.
// The site is initially set to inactive.
func CreateSite(db *gorm.DB, site *Site) error {
	site.IsActive = false
	result := db.Create(site)
	return result.Error
}

// GetSite retrieves a site from the database by its ID.
// Returns the site and an error if the site does not exist.
func GetSite(db *gorm.DB, id uint) (Site, error) {
	var site Site
	if err := db.Where("id = ?", id).First(&site).Error; err != nil {
		return Site{}, fmt.Errorf("Site does not exist: %w", err)
	}
	return site, nil
}

// GetSitesAll retrieves all sites from the database.
// Returns a slice of sites and an error if any occurs.
func GetSitesAll(db *gorm.DB) ([]Site, error) {
	var sites []Site
	if err := db.Find(&sites).Error; err != nil {
		return nil, fmt.Errorf("error getting sites: %w", err)
	}
	return sites, nil
}

func GetAllSSLCertificates(db *gorm.DB, id uint) ([]SSLcertificat.SSL, error) {
	var site Site
	site, err := GetSite(db, id)
	if err != nil {
		return nil, fmt.Errorf("error finding server: %w", err)
	}
	var ssl []SSLcertificat.SSL
	err = db.Model(&site).Association("SSlcerificate").Find(&ssl)
	return ssl, err
}

// DeleteSite deletes a site from the database by its ID.
// Returns an error if the site does not exist or if any other error occurs.
func DeleteSite(db *gorm.DB, id uint) error {
	result := db.Delete(&Site{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting site: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("site with Id %d not found", id)
	}
	return nil
}

// UpdateSite updates the details of an existing site in the database by its ID.
// Returns the updated site and an error if the site does not exist or if any other error occurs.
func UpdateSite(db *gorm.DB, id uint, updatedSite Site) (Site, error) {
	var site Site
	site, err := GetSite(db, id)
	if err != nil {
		return Site{}, fmt.Errorf("error finding site: %w", err)
	}
	if updatedSite.Domain != "" {
		site.Domain = updatedSite.Domain
	}
	if updatedSite.SiteName != "" {
		site.SiteName = updatedSite.SiteName
	}
	if err := db.Save(&site).Error; err != nil {
		return Site{}, fmt.Errorf("error saving site: %w", err)
	}
	return site, nil
}

// ActivateOrUnactivateSite toggles the active status of a site by its ID.
// Returns an error if the site does not exist or if any other error occurs.
func ActivateOrUnactivateSite(db *gorm.DB, id uint) error {
	var site Site
	site, err := GetSite(db, id)
	if err != nil {
		return fmt.Errorf("error finding site: %w", err)
	}
	if site.IsActive == true {
		site.IsActive = false
		if err := db.Save(&site).Error; err != nil {
			return fmt.Errorf("error saving site changes: %w", err)
		}
	} else {
		site.IsActive = true
		if err := db.Save(&site).Error; err != nil {
			return fmt.Errorf("error saving site changes: %w", err)
		}
	}
	return nil
}

// Location represents the location model
type Location struct {
	gorm.Model
	Body string
	Site []*Site `gorm:"many2many:location_site;"` // Many-to-many relationship with Site
}

// CreateLocation creates a new location in the database
// db: Database connection
// location: Location to be created
// Returns an error if the creation fails
func CreateLocation(db *gorm.DB, location *Location) error {
	result := db.Create(location)
	return result.Error
}

// GetLocation retrieves a location by ID
// db: Database connection
// id: Location ID
// Returns the location or an error if retrieval fails
func GetLocation(db *gorm.DB, id uint) (Location, error) {
	var location Location
	if err := db.Where("id = ?", id).First(&location).Error; err != nil {
		return Location{}, err
	}
	return location, nil
}

func GetLocationALL(db *gorm.DB) ([]Location, error) {
	var locations []Location
	if err := db.Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

// DeleteLocation deletes a location by ID
// db: Database connection
// id: Location ID
// Returns an error if deletion fails or if the location is not found
func DeleteLocation(db *gorm.DB, id uint) error {
	result := db.Delete(&Location{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func UpdateLocation(db *gorm.DB, id uint, updatedLocation Location) (Location, error) {
	var location Location
	if err := db.Where("id = ?", id).First(&location).Error; err != nil {
		return Location{}, err
	}
	location.Body = updatedLocation.Body
	result := db.Save(&location)
	return location, result.Error
}
