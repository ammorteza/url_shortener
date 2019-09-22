package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/src/controllers"
	"url-shortener/src/db"
	"url-shortener/src/interfaces"
)

func getDbConnection() interfaces.DbConnection{
	var dbConnection interfaces.DbConnection = db.MysqlConnection{}
	return dbConnection
}

func TestMakeUrl(t *testing.T){
	urlController := controllers.UrlController{}
	urlController.Init(getDbConnection())

	var jsonRequest = []byte(`{"base_url" : "/url/new/make"}`)
	req, err := http.NewRequest("POST" , "/url/shorten/make", bytes.NewBuffer(jsonRequest))
	if err != nil{
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	rtr := httptest.NewRecorder()
	handler := http.HandlerFunc(urlController.MakeUrl)
	handler.ServeHTTP(rtr, req)

	var responseJson controllers.ResponseBody
	if err := json.Unmarshal([]byte(rtr.Body.String()), responseJson); err != nil{
		t.Fatal(err)
	}

	if rtr.Code != http.StatusOK || responseJson.Short_url == ""{
		t.Error("make url api has an Error!")
	}
}
