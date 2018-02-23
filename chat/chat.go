package chat

import (
	"github.com/treychua/beatricethetelegrambot/lunchvenue"
	"github.com/treychua/beatricethetelegrambot/request"
	"gopkg.in/mgo.v2/bson"
)

type chat struct {
	ChatID int64                  `bson:"chatid"`
	Venues lunchvenue.LunchVenues `bson:"venues"`
}

// attempts to get a chat struct from the db. if the chat doesn't exist, creates it
func getChatFromDB(req *request.Request) (*chat, error) {
	s := req.Session.Copy()
	defer s.Close()

	collection := s.DB("beatricedb").C("chats")

	var ch chat
	err := collection.Find(bson.M{"chatid": req.ChatID}).One(&ch)
	if nil != err {
		errString := err.Error()
		if "not found" != errString {
			return nil, err
		}

		ch = chat{req.ChatID, lunchvenue.LunchVenues{}}
		err = collection.Insert(&ch)
		if nil != err {
			return nil, err
		}
	}

	return &ch, nil
}
