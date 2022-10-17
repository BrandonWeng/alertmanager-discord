package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

const defaultListenAddress = "127.0.0.1:9094"

var (
	webhookUrl    = flag.String("webhook", os.Getenv("DISCORD_WEBHOOK"), "Discord WebHook URL.")
	listenAddress = flag.String("listen.address", os.Getenv("LISTEN_ADDRESS"), "Address:Port to listen on.")
)

func main() {
	flag.Parse()
	if *listenAddress == "" {
		*listenAddress = defaultListenAddress
	}

	discordClient := getDiscordClient(*webhookUrl)

	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/alerts", discordClient.alertsHandler)

	log.Printf("Listening on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}
