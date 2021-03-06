package main

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
	//	"strings"
	"errors"
)

func main() {

	/////////////////////////////////////
	//Initialize the users in the chat
	/////////////////////////////////////
	InitUsers()
	fmt.Println("Initialized the Users")

	bot, err := tgbotapi.NewBotAPI("348577595:AAHHm0zC10iUEp6l5o2RP-pJ45Hl9f9DOnU")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	/////////////////////////////////////
	// Wait for updates to the inline query
	/////////////////////////////////////
	updates, err := bot.GetUpdatesChan(u)
	response := "Please Get started!"
	sendMessage := "Erro: Money not sent!"
	length := 0
	userName := ""
	moneyStart := 0
	money := 0
	end := 0
	var lastUpdate tgbotapi.Update

	fmt.Println("About to run loop!")
	for update := range updates {
		fmt.Println("Running Loop!")

		if update.InlineQuery == nil {
			fmt.Println("NIL")
			break
		}

		if update.InlineQuery.Query == "" {
			fmt.Println("QUERY IS EMPTY")
			continue
		}

		fmt.Println("Input string:", update.InlineQuery.Query)

		inputString := update.InlineQuery.Query
		length = len(inputString)

		if length < 3 {
			response = "Error: Not a valid command"
		} else if firstThree := inputString[0:3]; firstThree == "pay" {
			fmt.Println("you hit the pay command")
			if length == 3 {
				response = "Please enter your friends name"
			}

			if inputString[len(inputString)-2] == ' ' {
				end = len(inputString) - 2
				userName = inputString[5 : end+1]
				response = "Now enter a '$' followed by an amount!"
			}

			if inputString[len(inputString)-1] == '$' {
				moneyStart = len(inputString)
			}

			money, _ = strconv.Atoi(inputString[moneyStart:len(inputString)])

		}

		fmt.Println("Current amount:", money)
		fmt.Println("Current username:", userName)
		fmt.Println("made it past that!")

		// The article to display and message to return
		sendMessage = "You sent @" + userName + " $" + strconv.Itoa(money)
		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Click here when you're ready!", sendMessage)

		//Set the body of the article after each update
		article.Description = response

		// Not sure what this does yet
		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       []interface{}{article},
		}

		fmt.Println("made it to the end")

		// Don't take any errors!
		_, _ = bot.AnswerInlineQuery(inlineConf)
		lastUpdate = update
	}

	//maybe just throw the API call down here That might be the bets bet

	//Get Telegram ID for Sender
	senderID := ""
	receiverID := ""

	for _, u := range users {
		fmt.Println("Tele Name LEN", len(u.TelegramUsername))
		fmt.Println("user Name LEN", len(userName))
		if u.TelegramUsername == userName {
			fmt.Println("MATCH")
			receiverID = u.CapitalOneAccountID
		}

		if u.TelegramID == strconv.Itoa(lastUpdate.InlineQuery.From.ID) {
			senderID = u.CapitalOneAccountID
		}
	}

	fmt.Println("done with for loop!")
	fmt.Println("Rec ID: ", receiverID)
	fmt.Println("Send ID: ", senderID)

	if receiverID == "" || senderID == "" {
		return
	}

	SendPayment(senderID, receiverID, money)
	fmt.Println("MADEW THE CLALL MY MANE")
	fmt.Println("MADEW THE CLALL MY MANE")
	fmt.Println("MADEW THE CLALL MY MANE")
	fmt.Println("MADEW THE CLALL MY MANE")

	//length = 0
	//userName = ""
	//moneyStart = 0
	//money = 0
	//end = 0
}

func pay(messageArray []string) string {
	returnString := "default"

	if len(messageArray) < 3 {
		return "Error: There are less than three arguments"
	}

	//Get the username, and check that its valid
	atUser := messageArray[1]
	valid := ValidUser(atUser)
	if valid == false {
		return "Error: Username not found"
	}

	//Get the dollar amount and check that its valid
	amountToSend, err := strconv.Atoi(messageArray[2])
	err = validPrice(amountToSend)
	if err != nil {
		returnString = "Error: Invalid Price"
		return returnString
	}

	returnString = "Ready to pay!"

	return returnString
}

func validPrice(price int) error {

	if price < 0 {
		return errors.New("invalid price: negative value")
	}

	return nil
}

func ValidUser(user string) bool {
	valid := false
	user = user[1:len(user)]
	for i := 0; i < len(users); i++ {
		if user == users[i].TelegramUsername {
			valid = true
		}
	}
	return valid
}
