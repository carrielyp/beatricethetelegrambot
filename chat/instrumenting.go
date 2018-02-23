package chat

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/treychua/beatricethetelegrambot/request"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Svc            Service
}

func (mw InstrumentingMiddleware) getChat(r *request.Request) (c *chat, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getChat", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	c, err = mw.Svc.getChat(r)
	return
}

// func (mw InstrumentingMiddleware) HandleRequest(r *request.Request) (s string, err error) {
// 	defer func(begin time.Time) {
// 		lvs := []string{"method", "HandleRequest", "error", fmt.Sprint(err != nil)}
// 		mw.RequestCount.With(lvs...).Add(1)
// 		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	s, err = mw.Svc.HandleRequest(r)
// 	return
// }
