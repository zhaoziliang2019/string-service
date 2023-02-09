package plugins

import (
	"github.com/go-kit/kit/log"
	"github.com/zhaoziliang2019/string-service/service"
	"time"
)

type loggingMiddleware struct {
	service.Service
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) service.ServiceMiddleware {
	return func(next service.Service) service.Service {
		return loggingMiddleware{next, logger}
	}
}
func (mw loggingMiddleware) Concat(a, b string) (ret string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Concat",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(begin))
	}(time.Now())
	ret, err = mw.Service.Concat(a, b)
	return
}
func (mw loggingMiddleware) Diff(a, b string) (ret string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Diff",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(begin))
	}(time.Now())
	ret, err = mw.Service.Diff(a, b)
	return
}
func (mw loggingMiddleware) HealthCheck() (result bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "HealthCheck",
			"result", result,
			"took", time.Since(begin))
	}(time.Now())
	result = mw.Service.HealthCheck()
	return
}
