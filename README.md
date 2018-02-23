# beatricethetelegrambot
She's a nice lady(bot) who'll tell you where to go for lunch.

# how to start
1. Have docker up and running
2. run `docker-compose up` 
3. modify your telegram bot api key within main.go. it looks like this:
```golang
bot, err := tgbotapi.NewBotAPI("YOUR_KEY_OBTAINED_FROM_BOTFATHER_HERE")
```
4. run `go run main.go`

# features
1. add location
2. list locations
3. delete location

# future features
1. randomly obtain location
2. nearby location tracking

#test
