package structs

type Bioskop struct {
	ID		int		`json:"id"`
	Nama	string	`json:"nama"`
	Lokasi	string	`json:"lokasi"`
	Rating	float64	`json:"rating"`
}