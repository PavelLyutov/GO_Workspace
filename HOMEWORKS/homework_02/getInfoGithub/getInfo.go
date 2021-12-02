package getInfoGithub

import (
	"encoding/json"
	"homework_02/data"
	"homework_02/users"
	"log"
	"net/http"
	"strconv"
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
	user.DistributionWork = data.GetYearsByPercent(data.SortByValue(GetInfoFromContribution(username)))

	return user
}

func GetInfoFromContribution(username string) (map[string]int){
	years := []int{2017, 2018, 2019, 2020, 2021}
	contInfo := make(map[string]int)

	for _, v := range years {

		resp, err := http.Get("https://skyline.github.com/" + username + "/" + strconv.Itoa(v) + ".json")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var contribution users.Contribution
		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&contribution)
		var total int
		for i := 0; i < len(contribution.Count); i++ {
			for j := 0; j < len(contribution.Count[i].Days); j++ {
				total += contribution.Count[i].Days[j].Count
			}
		}
		contInfo[strconv.Itoa(v)] += total

	}
	return  contInfo
}

//func GetLastFiveYears(username string) PairList {
//	cont := make(PairList, 5)
//	years := []int{2017, 2018, 2019, 2020, 2021}
//	i := 0
//
//	for _, v := range years {
//		var key string
//		var value int
//		key, value = getInfoGithub.GetInfoFromContribution(username, v)
//		cont[i] = Pair{key, value}
//	}
//
//	return cont
//}

