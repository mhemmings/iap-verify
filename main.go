package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dogenzaka/go-iap/appstore"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		printUsage()
		os.Exit(1)
	}

	filePath := args[len(args)-1]

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Unable to open file: %s\n", filePath)
		os.Exit(1)
	}

	receipt := string(fileBytes)

	sandbox := flag.Bool("sandbox", false, "True if this is a sandbox receipt")
	secret := flag.String("secret", "", "Shared secret which is needed to validate auto-renewing subscriptions")
	flag.Parse()

	appStoreConf := appstore.Config{
		IsProduction: !*sandbox,
	}

	client := appstore.NewWithConfig(appStoreConf)

	req := appstore.IAPRequest{
		ReceiptData: strings.TrimSpace(receipt),
	}

	if len(*secret) > 0 {
		req.Password = *secret
	}

	resp := &appstore.IAPResponse{}

	err = client.Verify(req, resp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = appstore.HandleError(resp.Status)
	if err != nil {
		fmt.Printf("%d: %s\n", resp.Status, err.Error())
		os.Exit(0)
	}

	js, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(js))
}

func printUsage() {
	fmt.Println("USAGE: iap-verify [-sandbox] [-secret] path-to-receipt")
}
