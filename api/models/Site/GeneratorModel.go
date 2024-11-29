package models

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

type GeneratorModel struct {
	MainConfTemplatePath string
	SiteConfTemplatePath string
	MainConfPath         string
	SiteConfPath         string
}

func NewGeneratorModel(mainConfTemplatePath, siteConfTemplatePath, mainConfPath, siteConfPath string) *GeneratorModel {
	return &GeneratorModel{
		MainConfTemplatePath: mainConfTemplatePath,
		SiteConfTemplatePath: siteConfTemplatePath,
		MainConfPath:         mainConfPath,
		SiteConfPath:         siteConfPath,
	}
}

func GenerateRandomFileName(prefix, suffix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%d%s", prefix, rand.Intn(1000000), suffix)
}

func (g *GeneratorModel) CreateConfigCopies() (string, string, error) {
	mainConfCopy := filepath.Join("NginxConfigurators", GenerateRandomFileName("main_", ".conf"))
	siteConfCopy := filepath.Join("NginxConfigurators", GenerateRandomFileName("site_", ".conf"))

	if err := copyFile(g.MainConfTemplatePath, mainConfCopy); err != nil {
		return "", "", err
	}
	if err := copyFile(g.SiteConfTemplatePath, siteConfCopy); err != nil {
		return "", "", err
	}

	return mainConfCopy, siteConfCopy, nil
}

func (g *GeneratorModel) UpdateSiteConfig(siteConfPath, siteName, siteDomain string) error {
	content, err := ioutil.ReadFile(siteConfPath)
	if err != nil {
		return err
	}

	updatedContent := strings.ReplaceAll(string(content), "$SITE_NAME", siteName)
	updatedContent = strings.ReplaceAll(updatedContent, "$SITE_DOMAIN", siteDomain)

	return ioutil.WriteFile(siteConfPath, []byte(updatedContent), 0644)
}

func (g *GeneratorModel) IncludeSiteConfigInMain(mainConfPath, siteConfPath string) error {
	content, err := ioutil.ReadFile(mainConfPath)
	if err != nil {
		return err
	}

	includeDirective := fmt.Sprintf("include %s;", siteConfPath)
	updatedContent := strings.ReplaceAll(string(content), "# $SOME_DOMAIN_ENTRY1", includeDirective)

	return ioutil.WriteFile(mainConfPath, []byte(updatedContent), 0644)
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dst, input, 0644); err != nil {
		return err
	}

	return nil
}
