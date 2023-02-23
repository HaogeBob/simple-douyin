package main

import (
	"context"
	"log"
	"time"

	"github.com/simple/douyin/kitex_gen/feed"
	"github.com/simple/douyin/kitex_gen/feed/feedservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := feedservice.NewClient("feed", client.WithHostPorts("127.0.0.1:8081"))
	if err != nil {
		log.Fatal(err)
	}

	req := &feed.FeedRequest{LatestTime: int64(-62135596800), Token: "9ee0c09b80d47742d56a91994b8dd040"}
	resp, err := c.Feed(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

}
