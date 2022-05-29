package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/winebarrel/ssocreds"
)

func init() {
	log.SetFlags(0)
}

func main() {
	profile := os.Getenv("AWS_PROFILE")

	if profile == "" {
		log.Fatal("AWS_PROFILE is not set")
	}

	startUrl, err := ssocreds.SsoStartUrl(profile)

	if err != nil {
		log.Fatal(err)
	}

	accessToken, region, err := ssocreds.LastAccessTokenAndRegion(startUrl)

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	account, permissionSet, err := ssocreds.AccountAndPermissionSet(cfg)

	if err != nil {
		log.Fatal(err)
	}

	accessKeyId, secretAccessKey, sessionToken, err := ssocreds.SsoCredentials(cfg, account, permissionSet, accessToken, region)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("export AWS_ACCESS_KEY_ID='%s'\n", accessKeyId)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY='%s'\n", secretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN='%s'\n", sessionToken)
}
