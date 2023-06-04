package pkg

type GroupMain struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type GroupLocations struct {
	Places []string `json:"locations"`
}

type GroupDates struct {
	Dates []string `json:"dates"`
}

type GroupInfo struct {
	Places        GroupLocations
	Dates         GroupDates
	GroupPersonal GroupMain
}
