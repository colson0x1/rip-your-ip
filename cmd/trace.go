package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP.",
	Long:  `Trace the IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				// fmt.Println(ip)
				showData(ip)
			}
		} else {
			fmt.Println("Provide an IP address ayy!")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

// retrieved from HTTP GET postman fetched at http://ipinfo.io/1.1.1.1/geo
// {
//     "ip": "1.1.1.1",
//     "hostname": "one.one.one.one",
//     "anycast": true,
//     "city": "Los Angeles",
//     "region": "California",
//     "country": "US",
//     "loc": "34.0522,-118.2437",
//     "org": "AS13335 Cloudflare, Inc.",
//     "postal": "90076",
//     "timezone": "America/Los_Angeles",
//     "readme": "https://ipinfo.io/missingauth"
// }

// mapping into JSON
type IP = struct {
	IP       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Location string `json::"location"`
	Postal   string `json::"postal"`
	Timezone string `json::"timezone"`
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := IP{}

	// saving the response into JSON
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Oops! Failed to unmarshal HTTP response.")
	}

	fmt.Println("Data retrieved successfully.")

	fmt.Printf("IP address: %s\n City: %s\n Region: %s\n Country: %s\n Location: %s\n Postal: %s\n Timezone: %s\n",
		data.IP,
		data.City,
		data.Region,
		data.Country,
		data.Location,
		data.Postal,
		data.Timezone,
	)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Oops! Failed to get HTTP response.")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Oops! Failed to read HTTP response.")
	}

	return responseByte
}
