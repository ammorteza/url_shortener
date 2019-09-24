package controllers

import (
	"encoding/json"
	"github.com/ammorteza/url_shortener/db"
	"github.com/ammorteza/url_shortener/models"
	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
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

func (c *UrlController) MakeUrl(w http.ResponseWriter, r *http.Request){
	urlModel, err := models.NewUrlModel(c.dbConnection)
	if err != nil{
		http.Error(w, "error in make url model: " + err.Error(), http.StatusInternalServerError)
	}else{
		requestBody := requestBody{}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "error in get request params: " + err.Error(), http.StatusInternalServerError)
		}else{
			responseBody := ResponseBody{}
			hashedUrl, err := uuid.NewV4()

			if err != nil {
				http.Error(w, "error in make unique url (uuid):" + err.Error(), http.StatusInternalServerError)
			} else {
				responseBody.Short_url = "/" + hashedUrl.String()

				if err := urlModel.Insert(requestBody.Base_url, hashedUrl.String()); err != nil {
					http.Error(w, "error in insert new record in db: " + err.Error(), http.StatusInternalServerError)
				} else {
					response, err := json.Marshal(responseBody)
					if err != nil {
						http.Error(w, "error in make response: " + err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "error in make url model: " + err.Error(), http.StatusInternalServerError)
	}else{
		vars := mux.Vars(r)
		mainUrl, err := urlModel.GetMainUrl(vars["unique_url_id"])
		if err != nil {
			http.Error(w, "error in get url params: " + err.Error(), http.StatusInternalServerError)
		}else{
			http.Redirect(w, r, mainUrl, http.StatusSeeOther)
		}
	}
}
