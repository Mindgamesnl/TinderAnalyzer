package main

import (
	"github.com/Mindgamesnl/TinderAnalyzer/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	account := models.LoadAccountData("data.json")
	logrus.Info("Loaded profile " + account.User.Name)

	days := account.GetDays()

	// calculate average likes per day
	{
		totalLikes := 0
		for i := range days {
			totalLikes += days[i].Likes
		}
		average := totalLikes / len(days)
		logrus.Info(" - average of " + strconv.Itoa(average) + " likes per day, with a total of " + strconv.Itoa(totalLikes) + " likes.")
	}

	// calculate average passes per day
	{
		totalPasses := 0
		for i := range days {
			totalPasses += days[i].Passes
		}
		average := totalPasses / len(days)
		logrus.Info(" - average of " + strconv.Itoa(average) + " passes per day, with a total of " + strconv.Itoa(totalPasses) + " passes.")
	}

	// calculate matches
	{
		totalMatches := 0
		for i := range days {
			totalMatches += days[i].Matches
		}
		average := totalMatches / len(days)
		logrus.Info(" - average of " + strconv.Itoa(average) + " matches per day, with a total of " + strconv.Itoa(totalMatches) + " matches.")
	}
}
