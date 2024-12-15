package alchemy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func (a *Alchemy) makeCall(method string, requestBody interface{}, expectedResponse interface{}) (interface{}, error) {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest(method, a.AlchemyUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// var data interface{}
	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	fmt.Println("Erreur lors du Unmarshal :", err)
	// 	return nil, nil
	// }
	// var prettyJSON bytes.Buffer
	// err = json.Indent(&prettyJSON, body, "", "  ")
	// if err != nil {
	// 	fmt.Println("Erreur lors du formatage JSON :", err)
	// 	return nil, nil
	// }
	// fmt.Println("JSON Formaté :")
	// fmt.Println(prettyJSON.String())

	responseValue := reflect.New(reflect.TypeOf(expectedResponse).Elem()).Interface()
	if err := json.Unmarshal(body, &responseValue); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return responseValue, nil
}
