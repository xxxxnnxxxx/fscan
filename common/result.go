package common

// port status
const (
	PortStatus_Close = iota
	PortStauts_Open
)

type DetectResult struct {
	Port   int    `json:"port"`
	Status int    `json:"status"`
	Info   string `json:"info"`
}

type ResultNode struct {
	IP         string         `json:"ip"`
	Url        string         `json:"url"`
	Mode       string         `json:"mode"`
	DetectInfo []DetectResult `json:"results"`
}

type ResultSet struct {
	Nodes []ResultNode `json:"nodes"`
}

func NewResultSet() *ResultSet {
	return &ResultSet{
		Nodes: make([]ResultNode, 0),
	}
}
