package main

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/viper"
)

func main() {
	// Setup configuration with viper
	viper.SetDefault("API_ENDPOINT", "ovh-eu")
	viper.BindEnv("API_ENDPOINT")
	viper.BindEnv("APPLICATION_KEY")
	viper.BindEnv("APPLICATION_SECRET")
	viper.BindEnv("CONSUMER_KEY")
	viper.BindEnv("DNS_ZONE")
	viper.BindEnv("RECORD_ID")
	viper.BindEnv("SUBDOMAIN")
	viper.BindEnv("TARGET")
	viper.BindEnv("TTL")

	apiEndpoint := viper.GetString("API_ENDPOINT")
	appKey := viper.GetString("APPLICATION_KEY")
	appSecret := viper.GetString("APPLICATION_SECRET")
	consumerKey := viper.GetString("CONSUMER_KEY")

	client, err := ovh.NewClient(
		apiEndpoint,
		appKey,
		appSecret,
		consumerKey,
	)
	if err != nil {
		panic(err)
	}

	type Params struct {
		SubDomain string `json:"subDomain,omitempty"`
		Target    string `json:"target,omitempty"`
		TTL       int64  `json:"ttl,omitempty"`
	}

	var reqBody = &Params{
		SubDomain: viper.GetString("SUBDOMAIN"),
		Target:    viper.GetString("TARGET"),
		TTL:       viper.GetInt64("TTL"),
	}

	dnsZone := viper.GetString("DNS_ZONE")
	recordId := viper.GetString("RECORD_ID")
	query := fmt.Sprintf("/domain/zone/%v/record/%v", dnsZone, recordId)

	if err = client.Put(query, reqBody, nil); err != nil {
		panic(err)
	}
}
