package controller

import (
	"bytes"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/jmoiron/sqlx"
	"go_crawler/model"
	"go_crawler/utils"
	"log"
	//"./model/Proxy"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//从获取的proxy 再次提交 请求 验证是否 可用

func ProxyAll() []string{
	reponse,err := http.Get("http://172.18.63.111:5010/get_all/")
	if err !=nil {
		fmt.Println("get faild")
	}
	fmt.Println("data")
	b ,err := ioutil.ReadAll(reponse.Body)
	if err !=nil {
		panic(err.Error())
	}
	var  p model.Proxy
	err1 := json.Unmarshal([]byte(b), &p)
	if err1 != nil {
		panic(err1.Error())
	}
	result := make([]string,len(p))
	for  i :=0;i< len(p); i++ {
		//fmt.Println(p[i].Proxy)

		result[i] = "http://"+p[i].Proxy
	}

	return  result
}

//发送请求到https://httpbin.org/ip
//确定代理是否可用
func CheckProxy(p string){
	c := colly.NewCollector(colly.AllowURLRevisit())
	rp, err := proxy.RoundRobinProxySwitcher(p)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})
	err1 := c.Visit("https://httpbin.org/ip")
	if err1 !=nil {
		fmt.Println(err1.Error())
	}else {
		//如果成功 则直接入库，方便后续调用
		fmt.Println(p)
		var Db *sqlx.DB = utils.ConnectMysql()
		defer Db.Close()
		utils.AddProxy(Db,p)
	}
}

