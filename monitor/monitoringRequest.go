package monitor

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"uptime/colors"
	"uptime/config"
	"uptime/httpClient"
	"uptime/notifications"
	"uptime/utils"
)

func sendRequest(url string, route config.Route) {
	client := httpClient.NewHttpClient()

	req, err := http.NewRequest(route.Method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if route.Headers != nil {
		AddRequestHeaders(req, route)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	statusCode := res.StatusCode
	if statusCode != 200 {
		errorMsg := fmt.Sprintf("Error: %s. Status code: %s", url, strconv.Itoa(statusCode))
		fmt.Println(colors.ColorRed, errorMsg, colors.ColorReset)
		if os.Getenv("EMAILS_ENABLED") != "" {
			notifications.SendEmailNotification(url, statusCode)
		}
		if os.Getenv("SLACK_ENABLED") != "" {
			slackErrorMsg := fmt.Sprintf("Error while monitoring %s.  Status code: %s", url, strconv.Itoa(statusCode))
			notifications.SendSlackNotification(slackErrorMsg)
		}
		return
	}

	successMessage := fmt.Sprintf("Successfully: %s", url)
	fmt.Println(colors.ColorGreen, successMessage, colors.ColorReset)
}

func SendMonitoringRequest(baseUrl string, route config.Route){
	fullUrl := utils.CreateFullUrl(baseUrl, route.Path)
	sendRequest(fullUrl, route)
}