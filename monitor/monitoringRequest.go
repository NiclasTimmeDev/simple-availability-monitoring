package monitor

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"uptime/colors"
	"uptime/config"
	"uptime/httpClient"
	"uptime/notifications"
	"uptime/utils"
)

// SendMonitoringRequest sends request to a configured route from config.yml
// and sends error notifications to the admin if errors occur during
// the request.
func SendMonitoringRequest(baseUrl string, route config.Route){
	url := utils.CreateFullUrl(baseUrl, route.Path)
	client := httpClient.NewHttpClient()

	req, err := http.NewRequest(route.Method, url, nil)
	if err != nil {
		message :=  fmt.Sprintf("Error while monitoring %s.\nError: %s", url, err.Error())
		notifyViaAllChannels(url, message)
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
		errorMsg := fmt.Sprintf("Error while monitoring: %s.\nStatus code: %s", url, strconv.Itoa(statusCode))
		fmt.Println(colors.ColorRed, errorMsg, colors.ColorReset)
		notifyViaAllChannels(url, errorMsg)
		return
	}

	successMessage := fmt.Sprintf("Successfully: %s", url)
	fmt.Println(colors.ColorGreen, successMessage, colors.ColorReset)
}

// notifyViaAllChannels sends notifications to the admin
// via all notification channels (e.g., email and slack).
func notifyViaAllChannels(url string, message string){
	notifications.SendEmailNotification(url, message)
	notifications.SendSlackNotification(message)
}