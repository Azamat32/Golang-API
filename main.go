package main

import (
	"fmt"
	"group-tracker/internal"
	"group-tracker/pkg"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var (
	groups        []pkg.GroupMain
	place         pkg.GroupLocations
	dates         pkg.GroupDates
	groupInfo     pkg.GroupInfo
	isDataFetched bool
)

// Хэндлит 404

// Хэндлит главную страницу index.html
func mainHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(
		"./templates/index.html",
		"./templates/template.html",
	)
	if err != nil {
		fmt.Println("Cant Parse HTML files")
		return
	}
	if r.URL.Path != "/" {
		internal.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// Проверка на повторную загрузку JSON

	if !isDataFetched {
		internal.JsonHandler(&groups, "https://groupietrackers.herokuapp.com/api/artists")
		isDataFetched = true
	}

	err = html.Execute(w, groups)
	if err != nil {
		fmt.Println("Cant Execute data to send")
		return
	}
}

// Handle the Personal info
func showMoreHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) > 3 {
		internal.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// taking the id of group
	groupId, err := strconv.Atoi(r.URL.Path[len("/show/"):])
	html, err := template.ParseFiles(
		"./templates/about.html",
		"./templates/template.html",
	)
	if err != nil {
		fmt.Println("Cant Parse HTML files")
		return
	}

	var group pkg.GroupMain
	for _, elements := range groups {
		if groupId == elements.ID {
			group = elements
			break
		}
	}
	// checking for correct id
	if group.Name == "" {
		internal.ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	internal.JsonHandler2(&place, group.Locations)
	internal.JsonHandler3(&dates, group.ConcertDates)
	groupInfo.GroupPersonal = group
	groupInfo.Dates = dates
	groupInfo.Places = place
	err = html.Execute(w, groupInfo)

	if err != nil {
		fmt.Println("Cant Execute data to send", err)
		return
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/static/style.css", internal.CssHandler)
	http.HandleFunc("/show/", showMoreHandler)
	fmt.Println("Running on http://localhost:8080/ ")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
		return
	}
}
