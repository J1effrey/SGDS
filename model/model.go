package model

import (
	"strconv"
)

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

//将节点信息格式化输出
func (node *Node) String() string {
	return "NodeInfo {nodeId:" + strconv.Itoa(node.Id) + ", nodeIpAddr:" + node.IpAddr + ", port:" + node.Port + "}"
}

//将添加节点信息格式化
func (req AddToClusterMessage) String() string {
	return "AddToClusterMessage:{\n  source:" + req.Source.String() + ",\n  dest: " + req.Dest.String() + ",\n  message:" + req.Message + " }"
}
