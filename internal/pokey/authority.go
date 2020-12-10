package pokey

type Authority struct {
	ID          string `json:"id,omitempty"`
	Label       string `json:"label,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	NextSerial  int    `json:"next_serial,omitempty"`
}
