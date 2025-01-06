# cf-ddns-go

cf-ddns-go is a Go application that updates Cloudflare DNS records with the current public IP address.

## Features

- Fetches the current public IP address from `https://api.ipify.org`
- Updates Cloudflare DNS records with the fetched IP address
- Uses Cloudflare API to manage DNS records

## Prerequisites

- Go 1.23 or higher
- Cloudflare API token with permission to modify DNS records

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/mayanksoni1996/cf-ddns-go.git
   ```
2. Set the Environment variables
   - CF_API_TOKEN: Cloudflare API Token
   - CF_ZONE_NAME: Zone Name for the Cloudflare
   - CF_SUBDOMAIN: subdomain to update
  
## Usage
Run the application
```sh
./cf-ddns-linux-amd64
```
