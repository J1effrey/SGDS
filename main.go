package main

import (
	"flag"
	"fmt"
	"github.com/jeffrey/SGDS/model"
	"github.com/jeffrey/SGDS/util"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	makeMasterOnError := flag.Bool("makeMasterOnError", false, "make this node master if unable to connect to the cluster ip provided.")
	clusterip := flag.String("clusterip", "127.0.0.1:8001", "ip address of any node to connnect")
	myport := flag.String("myport", "8001", "ip address to run this node on. default is 8001.")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano()) //种子
	myid := rand.Intn(9999999)

	//获取ip地址
	myIp, _ := net.InterfaceAddrs()

	//创建nodeInfo结构体
	me := model.Node{Id: myid, IpAddr: myIp[13].String(), Port: *myport}
	dest := model.Node{Id: -1, IpAddr: strings.Split(*clusterip, ":")[0], Port: strings.Split(*clusterip, ":")[1]}
	fmt.Println("我的节点信息：", me.String())
	//尝试连接到集群，在已连接的情况下向集群发送请求
	ableToConnect := util.ConnectToCluster(me, dest)

	//如果dest节点不存在，则me节点为主节点启动，否则直接退出系统
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("将启动me节点为主节点")
		}
		if err := util.ListenOnPort(me); err != nil {
			fmt.Printf("There is an error: %v", err)
			return
		}
	} else {
		fmt.Println("正在退出系统，请设置me节点为主节点")
	}
}
