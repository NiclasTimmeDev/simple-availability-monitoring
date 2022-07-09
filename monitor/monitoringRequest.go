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

var allowedHttpMethods = []string{
	"GET",
	"POST",
	"PATCH",
	"PUT",
	"DELETE",
	"OPTIONS",
	"CONNECT",
	"HEAD",
	"TRACE",
}

// SendMonitoringRequest sends request to a configured route from config.yml
// and sends error notifications to the admin if errors occur during
// the request.
func SendMonitoringRequest(baseUrl string, route config.Route){

	// Verify that the request method is valid.
	isRequestMethodValid := utils.StringSliceContains(allowedHttpMethods, route.Method)
	
	url := utils.CreateFullUrl(baseUrl, route.Path)
	
	if isRequestMethodValid == false {
		errMsg := fmt.Sprintf("Invalid http method %s for url %s. Aborting", route.Method, url)
		panic(errMsg)
	}

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
	expectedStatusCode := route.ExpectedStatusCode
	if expectedStatusCode == 0 {
		expectedStatusCode = 200
	}

	if statusCode != expectedStatusCode {
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