package main

import (
	"github.com/go-playground/log/v8"
	"github.com/kordondev/automatic-website-backup/web"
)

func main() {
	base := "http://feuerwehr.karlsbad.de"

	savedWebsites := make(map[string]bool)
	websitesToSave := []string{"/website/"}

	i := 0
	for i < len(websitesToSave) {
		url := websitesToSave[i]
		log.Info(url)
		i++
		if _, exists := savedWebsites[url]; exists {
			continue
		}
		savedWebsites[url] = true
		website := web.FetchAndParse(base, url)
		log.Info(website.ToString(false))
		website.Save()
		websitesToSave = append(websitesToSave, *website.Links...)
	}
}
