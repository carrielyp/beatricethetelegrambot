package request

import mgo "gopkg.in/mgo.v2"

// Request contains the messages received from telegram
type Request struct {
	Session *mgo.Session
	ChatID  int64
	Message []string
}
