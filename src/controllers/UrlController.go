package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"url-shortener/src/models"
	"github.com/nu7hatch/gouuid"
)

type UrlController struct{
	Controller
}

type requestBody struct {
	Base_url string `json:"base_url"`
}

type ResponseBody struct {
	Short_url string `json:"shorter_url"`
} 

func (c *UrlController)MakeUrl(w http.ResponseWriter, r *http.Request) {
	urlModel := models.UrlModel{}
	urlModel.Init(c.dbConnection)

	requestBody := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil{
		panic(err)
	}

	responseBody := ResponseBody{}
	hashedUrl, err := uuid.NewV4()

	if err != nil{
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}else{
		responseBody.Short_url = "/" + hashedUrl.String()

		if err := urlModel.Insert(requestBody.Base_url, hashedUrl.String()); err != nil{
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}else{
			response, err := json.Marshal(responseBody)
			if err != nil{
				panic(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(nil)
			}else{
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(response)
			}
		}
	}
}

func (c *UrlController)RedirectShortenUrl(w http.ResponseWriter, r *http.Request)  {
	urlModel := models.UrlModel{}
	urlModel.Init(c.dbConnection)

	vars := mux.Vars(r)
	mainUrl, err := urlModel.GetMainUrl(vars["unique_url_id"])
	if err != nil{
		panic(err)
	}
	fmt.Println(mainUrl)
	http.Redirect(w, r, mainUrl, http.StatusSeeOther)
}
