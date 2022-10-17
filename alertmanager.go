package main

import "time"

type (

	// Timestamp is a helper for (un)marhalling time
	Timestamp time.Time

	// HookMessage is the message we receive from Alertmanager
	HookMessage struct {
		Version           string            `json:"version"`
		GroupKey          string            `json:"groupKey"`
		Status            string            `json:"status"`
		Receiver          string            `json:"receiver"`
		GroupLabels struct {
			Alertname string `json:"alertname"`
		} `json:"groupLabels"`
		CommonLabels struct {
			Alertname string `json:"alertname"`
		} `json:"commonLabels"`
		CommonAnnotations struct {
			Summary string `json:"summary"`
		} `json:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL"`
		Alerts            []Alert           `json:"alerts"`
	}

	// Alert is a single alert.
	Alert struct {
		Labels      map[string]string `json:"labels"`
		Annotations struct {
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		StartsAt    string            `json:"startsAt,omitempty"`
		EndsAt      string            `json:"endsAt,omitempty"`
		GeneratorURL string			  `json:"generatorURL,omitempty"`
		Status string			  `json:"status,omitempty"`
	}
)
