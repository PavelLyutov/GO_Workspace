package getInfoGithub

import (
	"encoding/json"
	"homework_02/data"
	"homework_02/users"
	"log"
	"net/http"
)

func GetFromGithub(username string) (users.UserInfo) {

	resp, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	resp1, err1 := http.Get("https://api.github.com/users/" + username + "/repos")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp1.Body.Close()

	var user users.UserInfo
	var repos []users.RepoInfo

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&user)

	decoder1 := json.NewDecoder(resp1.Body)
	decoder1.Decode(&repos)

	allLang := make(map[string]int)
	for i := 0; i < len(repos); i++ {
		user.Forks += repos[i].Forks

		resp2, err := http.Get("https://api.github.com/repos/" + username + "/" + repos[i].Name + "/languages")
		if err != nil {
			log.Fatal(err)
		}
		defer resp2.Body.Close()

		langInfo := make(map[string]int)
		decoder := json.NewDecoder(resp2.Body)
		decoder.Decode(&langInfo)

		for k, v := range langInfo {
			allLang[k] += v
		}
	}

		user.DistributionLanguage = data.GetTopFiveByPercent(data.SortByValue(allLang))
	return user
}
