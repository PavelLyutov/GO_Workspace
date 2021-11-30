package users

import "homework_02/data"

type UserInfo struct{
	Login string `json:"login"`
	Id int `json:"id"`
	Repos int `json:"public_repos"`
	Forks int
	DistributionLanguage data.PairList
	Followers int `json:"followers"`
}