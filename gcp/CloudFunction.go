package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mmcdole/gofeed"
)

const (
	BucketName           = "web3_blog_commitment"
	ObjectName           = "rss_feeds"
	InfuraKey            = ""
	PrivateKey           = ""
	PrivateKeyPassphrase = ""
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func FunctionEntry(ctx context.Context, m PubSubMessage) error {
	data, err := downloadFeeds()

	if err != nil {
		log.Fatal(err)
	}

	feeds := strings.Split(string(data), "\n")
	feedParser := gofeed.NewParser()
	contractSession := getContractSession()

	for _, feed := range feeds {
		if feed == "" {
			continue
		}

		parsed, err := feedParser.ParseURL(feed)

		if err != nil {
			continue
		}

		length := big.NewInt(int64(len(parsed.Items)))
		contractSession.SetCurrentLengthOfFeed(feed, length)
	}

	return nil
}

func downloadFeeds() ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket(BucketName).Object(ObjectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	fmt.Println("Existing object downloaded.")
	return data, nil
}

func getContractSession() *RssFeedLengthOracleSession {
	client, _ := ethclient.Dial("https://goerli.infura.io/v3/" + InfuraKey)
	rssOracle, err := NewRssFeedLengthOracle(common.HexToAddress("0x8DA1b37231C6A4562Df6C33F02937FE4E7b91787"), client)

	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(PrivateKey), PrivateKeyPassphrase)

	if err != nil {
		log.Fatal(err)
	}

	return &RssFeedLengthOracleSession{
		Contract:     rssOracle,
		TransactOpts: *auth,
	}
}
