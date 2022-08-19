package model

type Node struct {
	Id     int    `json:"id"`
	IpAddr string `json:"ipAddr"`
	Port   string `json:"port"`
}

type AddToClusterMessage struct {
	Source  Node   `json:"source"`
	Dest    Node   `json:"dest"`
	Message string `json:"message"`
}
