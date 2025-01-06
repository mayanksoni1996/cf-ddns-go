package main

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func filterRecord(records []cloudflare.DNSRecord, recordName string) cloudflare.DNSRecord {
	for _, record := range records {
		if record.Name == recordName {
			return record
		}
	}
	return cloudflare.DNSRecord{}
}
func main() {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatalf("Error getting IP: %s", err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading IP: %s", err)
	}
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CF_API_TOKEN"))
	zoneName := os.Getenv("CF_ZONE_NAME")
	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatalf("Error getting zone ID: %s", err)
	}
	recordName := os.Getenv("CF_SUBDOMAIN") + "." + zoneName
	//records, err := api.ListDNSRecords(context.Background(), zoneID, cloudflare.DNSRecord{Name: recordName})
	records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{Type: "A"})
	if err != nil {
		log.Fatalf("Error getting DNS records: %s", err)
	}
	record := filterRecord(records, recordName)
	record.Content = string(ip)
	resultSet, _ := api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateDNSRecordParams{
		Type:     record.Type,
		Name:     recordName,
		Content:  record.Content,
		Data:     record.Data,
		ID:       record.ID,
		Priority: record.Priority,
		TTL:      record.TTL,
		Proxied:  record.Proxied,
		Tags:     record.Tags,
	})
	log.Printf("Updated record: %s", resultSet.Name)
}
