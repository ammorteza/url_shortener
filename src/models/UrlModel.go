package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Url struct {
	gorm.Model
	UniqueHashUrl	string `gorm:"unique;not null;size:64"`
	MainUrl 		string `gorm:"not null"`
}

type UrlModel struct {
	Model
}

func (urlModel *UrlModel)Insert(mainUrl string, uniqueHashUrl string) {
	url := Url{MainUrl: mainUrl, UniqueHashUrl: uniqueHashUrl}
	urlModel.db.Create(&url)
}

func (urlModel *UrlModel)GetMainUrl(uniqueHashUrl string)  (string, error){
	url := Url{}
	fmt.Println("hashUrl:" + uniqueHashUrl)
	urlModel.db.Where("unique_hash_url = ?" , uniqueHashUrl).First(&url)
	if url.MainUrl == ""{
		return "" , errors.New("does not exist!")
	}
	return url.MainUrl, nil
}