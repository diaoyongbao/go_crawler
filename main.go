package main

import (
	_ "github.com/gocolly/colly/proxy"
	"go_crawler/controller"
)


//func get_proxy() string {
//	r,err := http.Get("http://172.18.63.111:5010/get/")
//	if err !=nil{
//		panic(err.Error())
//	}
//	b ,err := ioutil.ReadAll(r.Body)
//	if err != nil{
//		panic(err.Error())
//	}
//	var p Proxy
//	e := json.Unmarshal([]byte(b),&p)
//	if e != nil {
//		panic(e.Error())
//	}
//	return p.Proxy
//}


func main() {

	s := controller.ProxyAll()
	//fmt.Println(s)
	for i:=0 ; i<len(s);i++ {
		controller.CheckProxy(s[i])
	}
	//c := colly.NewCollector(
	//	colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	//)
	//
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})
	//
	//c.OnError(func(_ *colly.Response, err error) {
	//	fmt.Println("Something went wrong:", err)
	//})
	//
	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("Visited", r.Request.URL)
	//})
	//
	//c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})
	//
	//c.OnScraped(func(r *colly.Response) {
	//	fmt.Println("Finished", r.Request.URL)
	//})
	//fmt.Println(get_proxy())
	////if p, err := proxy.RoundRobinProxySwitcher(
	////	"http://"+get_proxy(),
	////); err == nil {
	//c.SetProxyFunc(randomProxySwitcher)
	////}
	//
	//c.Visit("https://movie.douban.com/top250?start=0&filter=")
}

