package main

import (
	"log"

	feed "github.com/simple/douyin/kitex_gen/feed/feedservice"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
