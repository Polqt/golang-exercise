package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-8678837880403-8661680762295-ZxAXTYh2jAlUGyB2Q24qKHFq")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A08KPSUJDCN-8674073199941-c2e0e05553e28eb7573f715b7cd7c7faa2eb0d092d7715eeb9d35ae72070f89a")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example: "my yob is 2004",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				println("Error converting year to int")
			}
			age := 2023 - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}