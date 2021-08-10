package models

import (
	"fmt"
	"log"
	"todo/config"
	"todo/types/pagination"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(config.Mysql), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	fmt.Println("MYSQL: " + config.Mysql + " CONNECTED")

	// now := time.Now().Unix()
	// article := Article{
	// 	Type:          1,
	// 	Status:        0,
	// 	Title:         "TEST3",
	// 	Content:       "TEST",
	// 	CreatedTime:   now,
	// 	UpdatedTime:   now,
	// 	CompletedTime: 0,
	// 	DeletedTime:   0,
	// 	IsDeleted:     0,
	// }
	// article2 := Article{
	// 	Type:          1,
	// 	Status:        0,
	// 	Title:         "TEST4",
	// 	Content:       "TEST2",
	// 	CreatedTime:   now,
	// 	UpdatedTime:   now,
	// 	CompletedTime: 0,
	// 	DeletedTime:   0,
	// 	IsDeleted:     0,
	// }
	// CreateArticles([]*Article{&article, &article2})
	// fmt.Printf("%+v\n", article)
	// fmt.Printf("%+v\n", article2)

	condition := ArticleCondition{}
	condition.Ids = []int{1, 2, 4, 5, 6, 7, 8}
	pagination := pagination.Pagination{
		Offset:   0,
		Limit:    2,
		OrderBy:  "id",
		OrderDir: pagination.DESC,
	}
	articles, count := GetArticles(condition, "TEST", &pagination)
	fmt.Printf("%+v, %d\n", articles, count)

	// condition := ArticleCondition{}
	// condition.Ids = []int{1, 2, 4}
	// DeleteArticles(condition, "")
}
