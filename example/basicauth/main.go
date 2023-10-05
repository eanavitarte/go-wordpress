package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/spf13/pflag"

	"github.com/qorpress/go-wordpress"
)

var (
	username string
	password string
	endpoint string
	help     bool
)

func main() {

	// read .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pflag.StringVarP(&username, "username", "", os.Getenv("WORDPRESS_USERNAME"), "wordpress' username.")
	pflag.StringVarP(&password, "password", "", os.Getenv("WORDPRESS_PASSWORD"), "wordpress' password.")
	pflag.StringVarP(&endpoint, "endpoint", "", os.Getenv("WORDPRESS_API_ENDPOINT"), "wordpress api endpoint (eg. https://x0rzkov.com/wp-json).")
	pflag.BoolVarP(&help, "help", "h", false, "help info")
	pflag.Parse()
	if help {
		pflag.PrintDefaults()
		os.Exit(1)
	}

	tp := wordpress.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	// create wp-api client
	client, _ := wordpress.NewClient(endpoint, tp.Client())

	ctx := context.Background()

	// get the currently authenticated users details
	authenticatedUser, _, err := client.Users.Me(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	pp.Printf("Authenticated user %+v", authenticatedUser)
}
