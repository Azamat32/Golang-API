package internal

import (
	"encoding/json"
	"fmt"
	"group-tracker/pkg"
	"io"
	"net/http"
)

func JsonHandler(groups *[]pkg.GroupMain, jsonStr string) {
	response, err := http.Get(jsonStr)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Cant read JSON File:", err)
		return
	}
	err = json.Unmarshal(body, &groups)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
	}
}

func JsonHandler2(places *pkg.GroupLocations, jsonStr string) {
	response, err := http.Get(jsonStr)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Cant read JSON File:", err)
		return
	}
	err = json.Unmarshal(body, &places)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
	}
}

func JsonHandler3(dates *pkg.GroupDates, jsonStr string) {
	response, err := http.Get(jsonStr)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Cant read JSON File:", err)
		return
	}
	err = json.Unmarshal(body, &dates)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
	}
}
