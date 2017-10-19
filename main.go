package main

import (
	"fmt"
	"github.com/cdipaolo/sentiment"
)

func main() {
	model, err := sentiment.Train()
	if err != nil {
	    panic(fmt.Sprintf("There was an issue training the model.\n\t%v\n", err))
	}

	analysis := model.SentimentAnalysis("You're mother is an great lady", sentiment.English)
	fmt.Print(analysis, "\n")
}
