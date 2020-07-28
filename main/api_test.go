package main

import (
	"github.com/thanhgit/docmg/main/db"
	"testing"
)

func TestGetServers(tester *testing.T)  {
	var servers []db.Server
	instance := db.GetInstance()
	db := instance.Find(&servers)
	if db.Error != nil{
		tester.Error("TestGetServers:" + db.Error.Error())
	}
}

func TestGetWebsites(tester *testing.T) {
	var websites []db.Website
	instance := db.GetInstance()
	db := instance.Find(&websites)
	if db.Error != nil {
		tester.Error("TestGetWebsites: " + db.Error.Error())
	}
}