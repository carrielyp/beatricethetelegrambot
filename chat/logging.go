package chat

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/treychua/beatricethetelegrambot/request"
)

// LoggingMiddleware is another layer to perform structured logging
type LoggingMiddleware struct {
	Logger log.Logger
	Svc    Service
}

//HandleRequest to perform structured logging for all requests
func (mw LoggingMiddleware) HandleRequest(r *request.Request) (s string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "HandleRequest",
			"input", fmt.Sprintf("r: %#v", r),
			"output", fmt.Sprintf("s: %v", s),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	s, err = handleRequest(r, mw)
	return
}

func (mw LoggingMiddleware) getChat(r *request.Request) (c *chat, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "getChat",
			"input", fmt.Sprintf("r: %#v", r),
			"output", c,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	c, err = mw.Svc.getChat(r)
	return
}

func (mw LoggingMiddleware) handleList(c *chat) (s string) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "handleList",
			"input", fmt.Sprintf("c: %#v", c),
			"output", s,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	s = mw.Svc.handleList(c)
	return
}

func (mw LoggingMiddleware) handleRand(c *chat) (s string) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "handleRand",
			"input", fmt.Sprintf("c: %#v", c),
			"output", s,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	s = mw.Svc.handleRand(c)
	return
}

func (mw LoggingMiddleware) handleAdd(c *chat, r *request.Request) (s string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "handleAdd",
			"input", fmt.Sprintf("c: %#v, r: %#v", c, r),
			"output", s,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	s, err = mw.Svc.handleAdd(c, r)
	return
}

func (mw LoggingMiddleware) handleDelete(c *chat, r *request.Request) (s string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "handleDelete",
			"input", fmt.Sprintf("c: %#v, r: %#v", c, r),
			"output", s,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	s, err = mw.Svc.handleDelete(c, r)
	return
}
