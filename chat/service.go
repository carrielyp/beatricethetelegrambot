package chat

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"github.com/treychua/beatricethetelegrambot/lunchvenue"

	"gopkg.in/mgo.v2/bson"

	"github.com/treychua/beatricethetelegrambot/request"
)

// Service handles any commands or inputs from the chat
type Service interface {
	HandleRequest(r *request.Request) (string, error)

	getChat(r *request.Request) (*chat, error)
	handleAdd(c *chat, r *request.Request) (string, error)
	handleDelete(c *chat, r *request.Request) (string, error)
	handleList(c *chat) string
	handleRand(c *chat) string
}

// ServiceImpl is an empty struct that'll have the methods for fulfilling Service's interface
type ServiceImpl struct{}

func (cs ServiceImpl) getChat(r *request.Request) (*chat, error) {
	return getChatFromDB(r)
}

func (cs ServiceImpl) handleAdd(c *chat, r *request.Request) (string, error) {
	if 2 > len(r.Message) {
		return "Sorry? I didn't catch what you said! Use '/add <your venue>'!", nil
	}

	location := strings.Join(r.Message[1:], " ")
	if location == "ERRORSTRING" {
		return "HUHHHHHHHHHH I BROKE", errors.New("GOT INVALID INPUT PANIC PANIC")
	}

	err := c.Venues.Add(location)
	if _, ok := err.(*lunchvenue.LocationAlreadyExistsError); ok {
		return `Location already exists~! Try somewhere else!  ¯\_(ツ)_/¯`, nil
	}
	err = updateChatTable(c, r)
	if nil != err {
		return "", err
	}

	return location + " added successfully~ Ehehe~", nil
}

func (cs ServiceImpl) handleDelete(c *chat, r *request.Request) (string, error) {
	if 2 > len(r.Message) {
		return "Huh? What'd you want removed, again? Try telling me '/delete <your venue>' again!", nil
	}

	location := strings.Join(r.Message[1:], " ")
	ok, err := c.Venues.Delete(location)

	if !ok {
		if _, errMsg := err.(*lunchvenue.LocationNotFoundError); errMsg {
			return "You didn't provide a venue that's in the list! >:(", nil
		}
	}

	err = updateChatTable(c, r)
	if nil != err {
		return "", err
	}

	reply := location + " removed successfully~\n"
	if 0 == len(c.Venues) {
		reply += "You have no other venues remaining! Add more soon, okay?"
	} else {
		reply += "You have the remaining venues~ \n"
		for i, v := range c.Venues {
			reply += strconv.Itoa(i+1) + ": " + v.Location + "\n"
		}
	}

	return reply, nil
}

func (cs ServiceImpl) handleList(c *chat) string {
	if 0 == len(c.Venues) {
		return "You have no venues for lunch~ :("
	}

	reply := "Here are some of the places you've added!\n"
	for i, v := range c.Venues {
		reply += strconv.Itoa(i+1) + ": " + v.Location + "\n"
	}

	return reply
}

func (cs ServiceImpl) handleRand(c *chat) string {
	if 0 == len(c.Venues) {
		return "Hey... you don't have any venues! >:("
	}

	i := rand.Intn(len(c.Venues))

	reply := "Okay... it'll have to be " + c.Venues[i].Location + " for today! :)"

	return reply
}

// HandleRequest handles a request given by Telegram and returns a reply message.
func (cs ServiceImpl) HandleRequest(r *request.Request) (string, error) {
	return handleRequest(r, cs)
}

// =============================================================================
// helpers
// =============================================================================
func handleRequest(r *request.Request, s Service) (string, error) {
	if 0 == len(r.Message) {
		return "", errors.New("chatService::handleRequest() - message of length 0")
	}

	c, err := s.getChat(r)
	if nil != err {
		return "", errors.New("chatService::handleRequest() - chat cannot be retrieved")
	}

	var reply string
	switch r.Message[0] {

	case "/add":
		reply, err = s.handleAdd(c, r)
		if err != nil {
			return reply, err
		}

	case "/list":
		reply = s.handleList(c)

	case "/del":
		reply, err = s.handleDelete(c, r)
		if err != nil {
			return reply, err
		}

	case "/rand":
		reply = s.handleRand(c)
	}

	return reply, nil
}

func updateChatTable(c *chat, r *request.Request) error {
	s := r.Session.Copy()
	defer s.Close()

	collection := s.DB("beatricedb").C("chats")
	return collection.Update(bson.M{"chatid": c.ChatID}, &c)
}
