package main

import (
	"go.uber.org/zap"
	"time"
	"fmt"
)

func main() {
	sugar()
	logger()
	//https://godoc.org/go.uber.org/zap#pkg-files
	newProd()

}

func sugar() {
	// In contexts where performance is nice,
	// but not critical, use the SugaredLogger. 
	//It's 4-10x faster than other structured logging packages 
	//and supports both structured and printf-style logging.

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow(
	"failed to fetch URL",
	 "url", "http://www.google.com",
	 "attempt", 3,
	 "backoff", time.Second,
	)		

sugar.Infof("failed to fetch URL: %s", "http://www.google.com")

}

func logger() {
	//In the rare contexts where every microsecond and
	//every allocation matter, use the Logger. It's even 
	//faster than the SugaredLogger and allocates far less,
	// but it only supports strongly-typed, structured logging.

	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info(
		"failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)	
}

func newProd() {
	logger, err := zap.NewProduction()
	if err != nil {
		  logger.Info("can't initialize zap logger:")
		  logger.Info(fmt.Sprintf("serving on none"))
	}
	defer logger.Sync()
}
