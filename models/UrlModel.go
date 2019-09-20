package models

import (
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

func (urlModel *UrlModel)Insert(mainUrl string, uniqueHashUrl string) error{
	url := Url{MainUrl: mainUrl, UniqueHashUrl: uniqueHashUrl}
	return urlModel.db.Create(&url).Error
}

func (urlModel *UrlModel)GetMainUrl(uniqueHashUrl string)  (string, error){
	url := Url{}
	fmt.Println("hashUrl:" + uniqueHashUrl)
	err := urlModel.db.Where("unique_hash_url = ?" , uniqueHashUrl).First(&url).Error
	if err != nil{
		return "" , err
	}
	return url.MainUrl, nil
}