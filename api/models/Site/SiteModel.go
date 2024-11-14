package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	SiteName      string `gorm:"unique;not null"`
	Domain        string `gorm:"unique;not null"`
	IsActive      bool
	NginxServerID uint //`gorm:"not null"`
	//Configurations []*Configuration.Configuration `gorm:"many2many:configuration_site;"`
	//Permission     Permission.Permission
	//Location       []*Location.Location `gorm:"many2many:location_site;"`
	//SSlcerificate  []SSLcertificat.SSL
}

func CreateSite(db *gorm.DB, site *Site) error {
	site.IsActive = false
	result := db.Create(site)
	return result.Error
}

func GetSite(db *gorm.DB, id uint) (Site, error) {
	var site Site
	if err := db.Where("id = ?", id).First(&site).Error; err != nil {
		return Site{}, fmt.Errorf("Site does not exist: %w", err)
	}
	return site, nil
}

func GetSitesAll(db *gorm.DB) ([]Site, error) {
	var sites []Site
	if err := db.Find(&sites).Error; err != nil {
		return nil, fmt.Errorf("error getting sites: %w", err)
	}
	return sites, nil
}

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

func ActivateOrUnactivateSite(db *gorm.DB, id uint) error {
	var site Site
	site, err := GetSite(db, id)
	if err != nil {
		return fmt.Errorf("error finding site: %w", err)
	}
	if site.IsActive == true {
		site.IsActive = false
		if err := db.Save(&site).Error; err != nil {
			return fmt.Errorf("error saving server: %w", err)
		}
	} else {
		site.IsActive = true
		if err := db.Save(&site).Error; err != nil {
			return fmt.Errorf("error saving server: %w", err)
		}
	}
	return nil
}
