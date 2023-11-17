package request

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func init() {
}

// Sent formdata type requesr, use basic auth
// url: request url
// method: request method
// username: basic auth username
// password: basic auth password
// formData: formdata map
// return: response body
func FormDataReq(url, username, password string, formData map[string]string) (responseData map[string]interface{}) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range formData {
		_ = writer.WriteField(k, v)
	}
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth(username, password)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Sent get type request, use basic auth
// url: request url
// username: basic auth username
// password: basic auth password
// queryData: query data map
func GET(url, username, password string, queryData map[string]string) (responseData map[string]interface{}) {
	if queryData != nil {
		url += "?"
		for k, v := range queryData {
			url += k + "=" + v + "&"
		}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth(username, password)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		log.Fatal(err)
	}
	return
}
