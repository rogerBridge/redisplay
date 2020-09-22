package main

import (
	"go_redis/controllers"
	"go_redis/router"
	"log"

	"github.com/valyala/fasthttp"
)

func init() {
	start()
	//// 搞一些闲置的redis连接
	//var wg sync.WaitGroup
	//for i := 0; i < 5000; i++ {
	//	wg.Add(2)
	//	go newConn(&wg, redis_config.Pool.Get())
	//	go newConn(&wg, redis_config.Pool1.Get())
	//}
	//wg.Wait()
	//log.Println("预热redis链接成功")
}

//
//// 预热一下客户端, 减少之后的redisPool的链接的内存分配建立连接导致的时间消耗
//func newConn(w *sync.WaitGroup, conn redis.Conn) {
//	defer conn.Close()
//	_, err := conn.Do("get", "name")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	w.Done()
//}

// preload func set
func start() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	err := controllers.InitStore()
	if err != nil {
		log.Println(err)
		return
	}
	// 加载MySQL中的limit到全局变量和redis中
	err = controllers.LoadLimit()
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	//u := new(structure.UserLogin)
	//u.Username = "hello"
	//u.Password = "12345678"
	//err := users.InsertUsers(u)
	//if err!=nil {
	//	log.Printf("%s\n", err)
	//}
	//err := users.VerifyUsers(u)
	//if err!=nil {
	//	log.Printf("user %v not exist\n", u)
	//}
	//orders.InsertOrders("xxxxxx", "leo2n", 123, 1, time.Now(), "process")
	// receive.Receive(common.Ch)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/buy", buy)
	//// "/cancelBuy" 这个接口只能由后台来调用
	//mux.HandleFunc("/cancelBuy", cancelBuy)
	//log.Println("Listening on 0.0.0.0:4000")
	//err := http.ListenAndServe("0.0.0.0:4000", mux)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//r := router.New()
	////r := mux.NewRouter()
	////r.Handle(fasthttp.MethodPost, "/buy", buy)
	//r.POST("/syncGoodsLimit", syncGoodsLimit)
	//r.GET("/goodsList", goodsList)
	//r.POST("/syncGoodsFromMysql2Redis", syncGoodsFromMysql2Redis)
	//r.POST("/syncGoodsFromRedis2Mysql", syncGoodsFromRedis2Mysql)
	//r.POST("/buy", auth.MiddleAuth(buy))
	//r.POST("/cancelBuy", cancelBuy)
	//r.POST("/login", Login)
	//r.POST("/logout", Logout)
	//r.POST("/register", Register)
	//mux := func(ctx *fasthttp.RequestCtx) {
	//	switch string(ctx.Path()) {
	//	case "/buy":
	//		buy(ctx)
	//	case "/cancelBuy":
	//		cancelBuy(ctx)
	//	default:
	//		ctx.Error("not found", fasthttp.StatusNotFound)
	//	}
	//}

	r := router.ThisRouter()
	log.Println("Listen on :4000")
	log.Fatalln(fasthttp.ListenAndServe(":4000", r.Handler))
}
