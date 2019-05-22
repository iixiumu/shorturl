package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MyURL struct {
	ID        uint64 `gorm:"Column:id;Type:BIGINT UNSIGNED AUTO_INCREMENT NOT NULL;PRIMARY_KEY" json:"id"`
	ShortUrl  string `gorm:"Column:short_url;Type:CHAR(6) NOT NULL" json:"short_url"`
	LongUrl   string `gorm:"Column:long_url;Type:VARCHAR(256) NOT NULL" json:"long_url"`
	CreatedAt time.Time
}

func (MyURL) TableName() string {
	return "my_url"
}

func (MyURL) Long2Short(longUrl string) (string, error) {
	var err error
	obj := &MyURL{}
	err = MyDB.Where("long_url = ?", longUrl).First(obj).Error
	if err != nil {
		return "", err
	} else {
		return obj.ShortUrl, nil
	}
}

func (MyURL) Short2Long(shortUrl string) (string, error) {
	var err error
	obj := &MyURL{}
	err = MyDB.Where("short_url = ?", shortUrl).First(obj).Error
	if err != nil {
		return "", err
	} else {
		return obj.LongUrl, nil
	}
}

func (MyURL) Save(shortUrl, longUrl string) error {
	var err error
	obj := &MyURL{
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	}
	err = MyDB.Save(obj).Error
	return err
}
