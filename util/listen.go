package util

import (
	"encoding/json"
	"fmt"
	"github.com/jeffrey/SGDS/model"
	"net"
)

//me节点连接其它节点成功或者自身成为主节点之后开始监听别的节点在未来可能对它自身的连接
func ListenOnPort(me model.Node) error {
	//监听即将到来的信息
	ln, _ := net.Listen("tcp", fmt.Sprint(":"+me.Port))
	//接受连接
	for {
		connIn, err := ln.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("Error received while listening.", me.Id)
			}
		} else {
			var requestMessage model.AddToClusterMessage
			if err = json.NewDecoder(connIn).Decode(&requestMessage); err != nil {
				return err
			}
			fmt.Println("Got request:\n" + requestMessage.String())

			text := "已添加你到集群"
			responseMessage := getAddToClusterMessage(me, requestMessage.Source, text)
			if err = json.NewEncoder(connIn).Encode(&responseMessage); err != nil {
				return err
			}
			if err = connIn.Close(); err != nil {
				return err
			}
		}
	}
}
