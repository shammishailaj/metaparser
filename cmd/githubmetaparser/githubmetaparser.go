package main

import (
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"time"
)

type GithubMetaSSHKeyFingerprints struct {
	MD5_RSA    string `json:"MD5_RSA"`
	MD5_DSA    string `json:"MD5_DSA"`
	SHA256_RSA string `json:"SHA256_RSA"`
	SHA256_DSA string `json:"SHA256_DSA"`
}

type GithubMeta struct {
	VerifiablePasswordAuthentication bool                         `json:"verifiable_password_authentication"`
	SSHKeyFingerprints               GithubMetaSSHKeyFingerprints `json:"ssh_key_fingerprints"`
	HooksIPs                         []string                     `json:"hooks"`
	WebIPs                           []string                     `json:"web"`
	ApiIPs                           []string                     `json:"api"`
	GitIPs                           []string                     `json:"git"`
	PagesIPs                         []string                     `json:"pages"`
	ImporterIPs                      []string                     `json:"importer"`
}

func main() {
	fmt.Printf("# Getting the Github Meta Data at: %s", time.Now().String())
	url := "https://api.github.com/meta"
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().EnableTrace().Get(url)

	if err != nil {
		fmt.Printf("# Error making request to %s. %s\n", url, err.Error())
		return
	}

	if resp.StatusCode() != 200 {
		fmt.Printf("# HTTP Status is %d (not 200). Won't continue\n", resp.StatusCode())
		return
	}

	var responseData GithubMeta
	responseDataErr := json.Unmarshal([]byte(resp.String()), &responseData)
	if responseDataErr != nil {
		fmt.Printf("# Error Unmarshalling response from %s. %s\n", url, responseDataErr.Error())
		return
	}

	fmt.Printf("# Now outputting IP data for nginx\n# Github IPs/CIDRs data for nginx\n")
	fmt.Printf("# Hooks IPs\n=================\n")
	for _, v := range responseData.HooksIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Web IPs\n=================\n")
	for _, v := range responseData.WebIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# API IPs\n=================\n")
	for _, v := range responseData.ApiIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Git IPs\n=================\n")
	for _, v := range responseData.GitIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Pages IPs\n=================\n")
	for _, v := range responseData.PagesIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Importer IPs\n=================\n")
	for _, v := range responseData.ImporterIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")
}