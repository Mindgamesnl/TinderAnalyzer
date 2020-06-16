package main

import (
	"github.com/Mindgamesnl/TinderAnalyzer/models"
	"github.com/sirupsen/logrus"
)

func main() {
	account := models.LoadAccountData("data.json")

	logrus.Info("Loaded profile " + account.User.Name)
}
