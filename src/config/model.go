package config

type Track struct {
	Name        string `json:"name"`
	Config      string `json:"config"`
	Description string `json:"description"`
	MaxSlots    int    `json:"max_slots"`
}

type Car struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Paintings   []string `json:"paintings"`
}
