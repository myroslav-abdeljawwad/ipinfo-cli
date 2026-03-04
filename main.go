package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Built by Myroslav Mokhammad Abdeljawwad
const (
	apiURLTemplate = "https://ipinfo.io/%s/json"
	userAgent      = "ipinfo-cli/1.0"
	timeout        = 5 * time.Second
)

type ipInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Hostname string `json:"hostname"`
}

func fetchIPInfo(ip string) (*ipInfo, error) {
	url := fmt.Sprintf(apiURLTemplate, ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info ipInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

func printIPInfo(info *ipInfo) {
	fmt.Printf("IP:       %s\n", info.IP)
	fmt.Printf("City:     %s\n", info.City)
	fmt.Printf("Region:   %s\n", info.Region)
	fmt.Printf("Country:  %s\n", info.Country)
	fmt.Printf("Location: %s\n", info.Loc)
	if info.Hostname != "" {
		fmt.Printf("Hostname: %s\n", info.Hostname)
	}
	if info.Org != "" {
		fmt.Printf("Org:      %s\n", info.Org)
	}
}

func main() {
	var (
		ip   string
		help bool
	)

	flag.StringVar(&ip, "ip", "", "IP address to lookup (defaults to your public IP)")
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()

	if help || flag.NArg() > 0 {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	if ip == "" {
		ip = "me"
	}

	info, err := fetchIPInfo(ip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching IP info: %v\n", err)
		os.Exit(1)
	}
	printIPInfo(info)
}