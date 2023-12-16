package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"bytes"
)

func main() {
	// Question 1
	url := "http://httpbin.org/get"
	firstParam := "FIRSTPARAM"
	firstVal := "FIRSTVAL"

	fullURL := fmt.Sprintf("%s?%s=%s", url, firstParam, firstVal)

	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("GET request successful!")
	
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
	
		fmt.Println("Response Body:")
		fmt.Println(string(body))
	} else {
		fmt.Printf("GET request failed with status code: %d\n", response.StatusCode)
	}

	// Question 2
	url = "http://diptest.com/http-post/"
	params := "FIRSTPARAM=FIRSTVAL"
	body := bytes.NewBufferString(params)

	response, err = http.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.Status)

	fmt.Println("Response Headers:", response.Header)

	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Body:", responseBody.String())

	// Question 3
	imageURL := "https://httpbingo.org/image"
	outputFilePath := "downloaded_image.jpg"

	err = downloadImage(imageURL, outputFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image downloaded successfully to", outputFilePath)
}

func downloadImage(url, outputPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		return err
	}

	return nil
}
