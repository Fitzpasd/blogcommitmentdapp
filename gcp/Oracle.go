package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BucketName = "web3_blog_commitment"
	ObjectName = "rss_feeds"
)

func main() {
	infuraKey := flag.String("infuraApiKey", "", "Infura API Key")
	contractAddress := flag.String("contractAddr", "", "Address of smart contract which emits events")
	flag.Parse()

	if *infuraKey == "" {
		log.Fatal(errors.New("Arg 'infuraApiKey' is required."))
	}

	if *contractAddress == "" {
		log.Fatal(errors.New("Arg 'contractAddr' is required."))
	}

	listenToEvents(*infuraKey, *contractAddress)
}

func listenToEvents(infuraKey, contractAddr string) {
	client := getEthClient(infuraKey)
	contractAbi := getContractAbi()
	logs, sub := getLogs(contractAddr, client)

	fmt.Println("Listening for events...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			event := struct {
				Rss_Feed string
			}{}

			err := contractAbi.UnpackIntoInterface(&event, "NewCommitment", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			go addEntryToBucket(event.Rss_Feed)
		}
	}
}

func getEthClient(infuraKey string) *ethclient.Client {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/" + infuraKey)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getContractAbi() abi.ABI {
	content, err := ioutil.ReadFile("./contract.json")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(content)))
	if err != nil {
		log.Fatal(err)
	}

	return contractAbi
}

func getLogs(contractAddr string, client *ethclient.Client) (chan types.Log, ethereum.Subscription) {
	contractAddress := common.HexToAddress(contractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	return logs, sub
}

func addEntryToBucket(entry string) {
	data, err := downloadExistingObject()

	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(data)
	buf.WriteString(entry)
	buf.WriteString("\n")

	uploadNewObject(buf)
}

func downloadExistingObject() ([]byte, error) {
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

func uploadNewObject(buf *bytes.Buffer) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(BucketName).Object(ObjectName).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err = io.Copy(wc, buf); err != nil {
		return err
	}

	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return err
	}

	fmt.Println("Uploaded new object.")
	return nil
}
