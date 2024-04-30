package constants

type Artist struct {
	ID           uint
	Image        string
	Name         string
	Members      []string
	CreationDate uint
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Locations struct {
	ID        uint
	Locations []string
	Dates     string
}

type Dates struct {
	ID    uint
	Dates []string
}

type Relations struct {
	ID             uint
	DatesLocations map[string][]string
}
