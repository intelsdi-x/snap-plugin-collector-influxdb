package influxdb

type stats struct {
	Results []struct {
		Series []struct {
			Name    string            `json:"name"`
			Columns []string          `json:"columns"`
			Values  [][]int           `json:"values"`
			Tags    map[string]string `json:"tags,omitempty"`
		} `json:"series"`
	} `json:"results"`
}
