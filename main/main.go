package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thanhgit/docmg/main/db"
	_ "github.com/thanhgit/docmg/main/db"
	_ "github.com/thanhgit/docmg/main/docs"
	"github.com/thanhgit/docmg/main/dto"
	_ "github.com/thanhgit/docmg/main/dto"
	"golang.org/x/oauth2"
	"log"
	"sync"
)

func user_group(userName string, groupId int)  {
	user := db.User{
		Name:      userName,
	}
	group := db.GroupUser{
		ID: groupId,
	}

	db.CreateUser(user)

	db.CreateUser_group(
			db.User_group{
				User:       user,
				Group:      group,
			},
		)
}


func connectDB() {
	instance := db.GetInstance()

	if instance != nil {
		println("Connect MySQL OK")

	} else {
		println("Fail to connect")
	}

	//Create20Group()
	//Create1000User()
	requests := make(chan db.Document, 10)
	quit := make(chan bool)
	go func(r chan db.Document, quit chan bool) {
		for {
			select {
			case doc := <-r:
				db.CreateDocument(doc)
				break
			case <-quit:
				return
			}
		}
	}(requests, quit)

	//Create1000Doc(2000, requests)
	//Create1000Doc(3000, requests)
	//
	//for index:=100; index<300; index++ {
	//	Create1000Doc(index*1000, requests)
	//}
	//
	//<- quit

	//for index:= 1; index < 10; index++ {
	//	CreatePattern(index)
	//}

	//db.CreateDocument(db.Document{
	//	Author_id: 1,
	//	Parent_id: 10,
	//	Obj_type:  fmt.Sprintf("File"),
	//	Title:     fmt.Sprintf("Test.png"),
	//	Level: 10,
	//})

}

func Create20Group()  {
	// create 20 group
	wg := sync.WaitGroup{}
	wg.Add(20)
	for index:=0; index< 20; index++ {
		go func(_index int) {
			defer wg.Done()
			groupname := fmt.Sprintf("Group%d", _index)
			db.CreateGroup(db.GroupUser{
				Name:      groupname,
			})
		}(index)
	}

	wg.Wait()
}

func Create1000User()  {
	max := 1000
	wg := sync.WaitGroup{}
	wg.Add(max)
	for index:=0; index < max; index++ {
		username := fmt.Sprintf("Nguyen van A %d", index)
		go func(_index int, _username string) {
			defer wg.Done()
			user_group(_username, _index%20)
		}(index, username)
	}
	wg.Wait()
}

func Create1000Doc(from int, in chan db.Document)  {
	max := 1000
	wg := sync.WaitGroup{}
	wg.Add(max)
	for index:= from; index < max+from; index++ {
		go func(_index int) {
			defer wg.Done()
			in <- db.Document{
						Author_id: _index%1000,
						Parent_id: _index%10,
						Obj_type:  fmt.Sprintf("Type%d", _index%6),
						Title:     fmt.Sprintf("My title %d", _index),
					}

		}(index)
	}

	wg.Wait()
}

func CreatePattern(level int) {
	//for index:=0; index < 10; index++ {
	//	t := ""
	//	if index%1==0 {
	//		t = "Directory"
	//	} else {
	//		t = "File"
	//	}
		db.CreateDocument(db.Document{
			Author_id: 1,
			Parent_id: level - 1,
			Obj_type:  fmt.Sprintf("Type Directory"),
			Title:     fmt.Sprintf("Directory level %d", level),
			Level: 		level,
		})
	//}
}

func connectHydra(context context.Context) {
	conf := &oauth2.Config{
		ClientID:     "app-app",
		ClientSecret: "consumer-secret",
		Scopes:       []string{"openid", "offline"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://192.168.1.151:9000/oauth2/auth",
			TokenURL: "http://192.168.1.151:9000/oauth2/token",
		},
		RedirectURL: "https://webhook.site/543d9bdb-bb70-496e-baf6-6e2ba2e6e3fd",
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(context, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(context, tok)
	client.Get("...")
}

func setupGlobalMiddleware(c *gin.Context) *gin.Context {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	return c
}

func main()  {
	//connectDB()
	sitemap := db.GetSiteMap("main/db/project1.json")
	//fmt.Printf("Sitemap: %v", *sitemap)

	db.SaveSiteMap(*sitemap)
}

func main_temp() {
	context := context.Background()
	connectDB()
	connectHydra(context)

	r := gin.New()

	url := ginSwagger.URL("http://localhost:7000/swagger/doc.json") // The url endpoint
	// API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// authentication
	r.POST("/api/login", func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		var acc dto.Account
		if c.BindJSON(&acc) != nil {
			c.JSON(200, gin.H{
				"Status": "Bad request",
				"Code":   200,
			})
		} else {
			c.JSON(200, gin.H{
				"Status": "string",
				"Code": 200,
			})
		}
	})

	r.Run(":7000")
}
