package main

import (
	"github.com/go-playground/log/v8"
	"github.com/kordondev/automatic-website-backup/web"
)

func main() {
	base := "http://feuerwehr.karlsbad.de"
	url := "/website/"
	website := web.FetchAndParse(base, url)
	log.Info(website.ToString(false))
	website.Save()
}
