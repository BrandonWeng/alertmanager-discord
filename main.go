package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	webhookUrl    = flag.String("webhook.url", os.Getenv("DISCORD_WEBHOOK"), "Discord WebHook URL.")
	listenAddress = flag.String("listen.address", os.Getenv("LISTEN_ADDRESS"), "Address:Port to listen on.")
)

func main() {
	flag.Parse()

	discordClient := getDiscordClient(*webhookUrl)

	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/alerts", discordClient.alertsHandler)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}
