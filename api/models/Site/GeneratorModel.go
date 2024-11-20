package models

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type GeneratorModel struct {
	MainConfTemplatePath string
	SiteConfTemplatePath string
	MainConfPath         string
	SiteConfPath         string
}

func (g *GeneratorModel) GenerateSiteConfig(site Site) error {
	// Read the site template
	siteTemplate, err := ioutil.ReadFile(g.SiteConfTemplatePath)
	if err != nil {
		return err
	}

	// Replace placeholders with site data
	siteConfig := string(siteTemplate)
	siteConfig = strings.ReplaceAll(siteConfig, "$SITE_NAME", site.SiteName)
	siteConfig = strings.ReplaceAll(siteConfig, "$SITE_DOMAIN", site.Domain)

	// Write the new site config
	siteConfPath := fmt.Sprintf("%s/%s.conf", g.SiteConfPath, site.SiteName)
	err = ioutil.WriteFile(siteConfPath, []byte(siteConfig), 0644)
	if err != nil {
		return err
	}

	// Read the main template
	mainTemplate, err := ioutil.ReadFile(g.MainConfTemplatePath)
	if err != nil {
		return err
	}

	// Add include statement for the new site config
	mainConfig := string(mainTemplate)
	includeStatement := fmt.Sprintf("include %s;", g.SiteConfPath)
	mainConfig = strings.Replace(mainConfig, "# $SOME_DOMAIN_ENTRY1", includeStatement, 1)

	// Write the new main config
	err = ioutil.WriteFile(g.MainConfPath, []byte(mainConfig), 0644)
	if err != nil {
		return err
	}

	return nil
}
