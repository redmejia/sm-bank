package bank

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	bankNumber = "111122223333"
	minCard    = 1000
	maxCard    = 3000
	minCv      = 100
	maxCv      = 200
)

// This function generate new card the first the 12 number are bank.
// The bank number are 111122223333 total 12. this is just for practice
// but I provably generate the last 8 number so they can be differnt number
// displaying 1111222267895312.
func generateCard() (card string) {
	rand.Seed(time.Now().UnixNano())
	// generate the last four numbers
	var generateNumber int = rand.Intn(maxCard-minCard) + minCard
	var lastNumber string = strconv.Itoa(generateNumber)
	card = bankNumber + lastNumber
	return
}

// This function will generate the cv numbers random everytime function is call
func generateCvNumber() (cv uint8) {
	rand.Seed(time.Now().UnixNano())
	var generateCvNum int = rand.Intn(maxCv-minCv) + minCv
	cv = uint8(generateCvNum)
	return
}

// createCard ... this can take Account
func createCard() (clientCard *card) {
	// generate the cards here.
	var cardType string = "debit"
	var cardNumber string = generateCard()
	var cv uint8 = generateCvNumber()
	clientCard = &card{
		cardType:   cardType,
		cardNumber: cardNumber,
		cvNumber:   cv,
	}
	return
}
