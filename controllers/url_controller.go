package controllers

import (
	"encoding/json"
	"github.com/ammorteza/url_shortener/db"
	"github.com/ammorteza/url_shortener/models"
	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

type UrlController struct {
	Controller
}

type requestBody struct {
	Base_url string `json:"base_url"`
}

type ResponseBody struct {
	Short_url string `json:"short_url"`
}

func NewUrlController(connection db.DbConnection) *UrlController{
	urlController := &UrlController{}
	urlController.dbConnection = connection
	return urlController
}

func (c *UrlController) MakeUrl(w http.ResponseWriter, r *http.Request) {
	urlModel, err := models.NewUrlModel(c.dbConnection)
	if err != nil{
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}else{
		requestBody := requestBody{}
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err = json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}else{
			responseBody := ResponseBody{}
			hashedUrl, err := uuid.NewV4()

			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(nil)
			} else {
				responseBody.Short_url = "/" + hashedUrl.String()

				if err := urlModel.Insert(requestBody.Base_url, hashedUrl.String()); err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(nil)
				} else {
					response, err := json.Marshal(responseBody)
					if err != nil {
						log.Fatal(err)
						w.WriteHeader(http.StatusInternalServerError)
						w.Write(nil)
					} else {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)
						w.Write(response)
					}
				}
			}
		}
	}
}

func (c *UrlController) RedirectShortenUrl(w http.ResponseWriter, r *http.Request) {
	urlModel, err := models.NewUrlModel(c.dbConnection)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		vars := mux.Vars(r)
		mainUrl, err := urlModel.GetMainUrl(vars["unique_url_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}else{
			http.Redirect(w, r, mainUrl, http.StatusSeeOther)
		}
	}
}
