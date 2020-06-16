package TinderAnalyzer

import (
	"io/ioutil"
	"os"
)

func main() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll("./data.json")
	
}
