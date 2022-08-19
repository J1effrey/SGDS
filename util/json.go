package util

import (
	"encoding/json"
	"fmt"
	"github.com/jeffrey/SGDS/model"
	"net"
	"time"
)

//发送请求时格式化json包有用的工具
func getAddToClusterMessage(source model.Node, dest model.Node, message string) model.AddToClusterMessage {
	return model.AddToClusterMessage{
		Source: model.Node{
			Id:     source.Id,
			IpAddr: source.IpAddr,
			Port:   source.Port},
		Dest: model.Node{
			Id:     dest.Id,
			IpAddr: dest.IpAddr,
			Port:   dest.Port},
		Message: message,
	}
}

func ConnectToCluster(me model.Node, dest model.Node) bool {
	//连接到socket的相关细节信息
	connOut, err := net.DialTimeout("tcp", dest.IpAddr+":"+dest.Port, time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("不能连接到集群", me.Id)
			return false
		}
	} else {
		fmt.Println("连接到集群")
		text := "Hi nody.. 请添加我到集群"
		requestMessage := getAddToClusterMessage(me, dest, text)
		if err = json.NewEncoder(connOut).Encode(&requestMessage); err != nil {
			return false
		}

		decoder := json.NewDecoder(connOut)
		var responseMessage model.AddToClusterMessage
		if err = decoder.Decode(&responseMessage); err != nil {
			return false
		}
		fmt.Println("得到数据响应:\n" + responseMessage.String())
		return true
	}
	return false
}
