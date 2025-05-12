type Record struct {
	Name string `json:"name" csv:"name"`
	IP   string `json:"ip" csv:"IP"`
	// Puedes agregar más campos si los tiene tu CSV
}

type DataNet struct {
	Records []Record `json:"datalist"`
}
