package models

import (
	"fmt"
	"log"
	"todo/config"
	"todo/types/pagination"
	"todo/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var articleModel = ArticleModel{}
var tagModel = TagModel{}
var articleTagRefModel = ArticleTagRefModel{}

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(config.Mysql), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	fmt.Println("MYSQL: " + config.Mysql + " CONNECTED")
}

func getTxOrDb(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return db
}

func GetArticles(condition ArticleCondition, keyword string, p *pagination.Pagination, tagIds []int) (articles []Article, count int64) {
	if tagIds != nil {
		refCondition := ArticleTagRefCondition{TagIds: tagIds}
		refs := articleTagRefModel.getArticleTagRefs(refCondition)
		if condition.Ids == nil {
			condition.Ids = []int{}
		}
		for _, ref := range refs {
			condition.Ids = append(condition.Ids, ref.ArticleId)
		}
		condition.Ids = utils.UniqInts(condition.Ids)
	}
	articles, count = articleModel.getArticles(condition, keyword, p)
	return
}

func CreateArticle(article *Article, tagIds []int) {
	db.Transaction(func(tx *gorm.DB) error {
		if err := articleModel.createArticles([]*Article{article}, tx); err != nil {
			return err
		}
		if tagIds != nil {
			refs := []*ArticleTagRef{}
			for _, tagId := range tagIds {
				refs = append(refs, &ArticleTagRef{ArticleId: article.Id, TagId: tagId})
			}
			if err := articleTagRefModel.createArticleTagRefs(refs, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateArticle(updation Article, tagIds []int) {
	db.Transaction(func(tx *gorm.DB) error {
		if err := articleModel.updateArticles(ArticleCondition{Ids: []int{updation.Id}}, "", updation, tx); err != nil {
			return err
		}
		if tagIds != nil {
			tagIdsSet := make(map[int]bool)
			for _, tagId := range tagIds {
				if _, ok := tagIdsSet[tagId]; !ok {
					tagIdsSet[tagId] = true
				}
			}
			refs := articleTagRefModel.getArticleTagRefs(ArticleTagRefCondition{ArticleIds: []int{updation.Id}})
			toCreateRefs := []*ArticleTagRef{}
			toDeleteRefsCondition := ArticleTagRefCondition{Ids: []int{}}
			for _, ref := range refs {
				if _, ok := tagIdsSet[ref.TagId]; !ok {
					toCreateRefs = append(toCreateRefs, &ArticleTagRef{ArticleId: updation.Id, TagId: ref.TagId})
				} else {
					toDeleteRefsCondition.Ids = append(toDeleteRefsCondition.Ids, ref.Id)
				}
			}
			if err := articleTagRefModel.createArticleTagRefs(toCreateRefs, tx); err != nil {
				return err
			}
			if err := articleTagRefModel.deleteArticleTagRefs(toDeleteRefsCondition, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteArticles(condition ArticleCondition, keyword string, hardDel bool) {
	if hardDel {
		articles, _ := articleModel.getArticles(condition, keyword, nil)
		articleIds := []int{}
		for _, article := range articles {
			articleIds = append(articleIds, article.Id)
		}
		db.Transaction(func(tx *gorm.DB) error {
			if err := articleModel.deleteArticles(condition, keyword, hardDel, tx); err != nil {
				return err
			}
			if err := articleTagRefModel.deleteArticleTagRefs(ArticleTagRefCondition{ArticleIds: articleIds}, tx); err != nil {
				return err
			}
			return nil
		})
	} else {
		articleModel.deleteArticles(condition, keyword, hardDel, nil)
	}
}

func GetTags(condition TagCondition, keyword string, p *pagination.Pagination, articleIds []int) (tags []Tag, count int64) {
	if articleIds != nil {
		refCondition := ArticleTagRefCondition{ArticleIds: articleIds}
		refs := articleTagRefModel.getArticleTagRefs(refCondition)
		if condition.Ids == nil {
			condition.Ids = []int{}
		}
		for _, ref := range refs {
			condition.Ids = append(condition.Ids, ref.TagId)
		}
		condition.Ids = utils.UniqInts(condition.Ids)
	}
	tags, count = tagModel.getTags(condition, keyword, p)
	return
}

func CreateTags(tags []*Tag) error {
	return tagModel.createTags(tags, nil)
}

func UpdateTags(condition TagCondition, keyword string, updation Tag) error {
	return tagModel.updateTags(condition, keyword, updation, nil)
}

func UpdateTagsByEntities(tags []*Tag) error {
	return tagModel.updateTagsByEntities(tags, nil)
}

func DeleteTags(condition TagCondition, keyword string) error {
	tags, _ := tagModel.getTags(condition, keyword, nil)
	tagIds := []int{}
	for _, tag := range tags {
		tagIds = append(tagIds, tag.Id)
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tagModel.deleteTags(condition, keyword, tx); err != nil {
			return err
		}
		if err := articleTagRefModel.deleteArticleTagRefs(ArticleTagRefCondition{TagIds: tagIds}, tx); err != nil {
			return err
		}
		return nil
	})
}

func GetArticleTagRefs(condition ArticleTagRefCondition) (article_tag_refs []ArticleTagRef) {
	article_tag_refs = articleTagRefModel.getArticleTagRefs(condition)
	return
}
