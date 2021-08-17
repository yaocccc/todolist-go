package models

import (
	"fmt"
	"log"
	"todo/config"
	"todo/types/pagination"
	"todo/utils"

	"github.com/chenhg5/collection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var articleModel = ArticleModel{}
var tagModel = TagModel{}
var articleTagRefModel = ArticleTagRefModel{}

type ArticleCreation struct {
	Article
	TagIds []int
}

type ArticleUpdation struct {
	Article
	TagIds []int
}

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(config.Mysql), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	utils.Logger(" SYSTEM ", "INFO", "MYSQL: "+config.Mysql+" CONNECTED")
}

func getTxOrDb(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return db
}

func GetArticles(condition ArticleCondition, keyword string, p *pagination.Pagination, tagIds []int) (articles []Article, count int64) {
	/** if query tagIds -> filter articleIds by articleTagRefs */
	if tagIds != nil {
		refCondition := ArticleTagRefCondition{TagIds: tagIds}
		refs := articleTagRefModel.getArticleTagRefs(refCondition)
		if condition.Ids == nil {
			condition.Ids = []int{}
		}
		for _, ref := range refs {
			condition.Ids = append(condition.Ids, ref.ArticleId)
		}
		condition.Ids = collection.Collect(condition.Ids).Unique().ToIntArray()
	}
	articles, count = articleModel.getArticles(condition, keyword, p)
	return
}

func CreateArticles(creations []*ArticleCreation) error {
	return db.Transaction(func(tx *gorm.DB) error {
		articles := []*Article{}
		for _, creation := range creations {
			articles = append(articles, &creation.Article)
		}
		if err := articleModel.createArticles(articles, tx); err != nil {
			return err
		}

		refs := []*ArticleTagRef{}
		for _, creation := range creations {
			for _, tagId := range creation.TagIds {
				refs = append(refs, &ArticleTagRef{ArticleId: creation.Id, TagId: tagId})
			}
		}
		if err := articleTagRefModel.createArticleTagRefs(refs, tx); err != nil {
			return err
		}
		return nil
	})
}

func UpdateArticles(updations []ArticleUpdation) error {
	articles := []*Article{}
	articleIds := []int{}
	for _, updation := range updations {
		articles = append(articles, &updation.Article)
		articleIds = append(articleIds, updation.Id)
	}

	toCreateRefs := []*ArticleTagRef{}
	toDeleteRefsCondition := ArticleTagRefCondition{Ids: []int{}}
	refs := articleTagRefModel.getArticleTagRefs(ArticleTagRefCondition{ArticleIds: articleIds})
	refMapByKey := make(map[string]bool) /** key: articleId-tagId */
	for _, ref := range refs {
		refMapByKey[fmt.Sprintf("%d-%d", ref.ArticleId, ref.TagId)] = true
	}
	for _, updation := range updations {
		if updation.TagIds != nil {
			updationMapByKey := make(map[string]bool) /** key: articleId-tagId */
			for _, tagId := range updation.TagIds {
				updationMapByKey[fmt.Sprintf("%d-%d", updation.Id, tagId)] = true
				if _, ok := refMapByKey[fmt.Sprintf("%d-%d", updation.Id, tagId)]; !ok {
					toCreateRefs = append(toCreateRefs, &ArticleTagRef{ArticleId: updation.Id, TagId: tagId})
				}
			}
			for _, ref := range refs {
				if _, ok := updationMapByKey[fmt.Sprintf("%d-%d", ref.ArticleId, ref.TagId)]; !ok {
					toDeleteRefsCondition.Ids = append(toDeleteRefsCondition.Ids, ref.Id)
				}
			}
		}
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := articleModel.updateArticles(articles, tx); err != nil {
			return err
		}
		if err := articleTagRefModel.createArticleTagRefs(toCreateRefs, tx); err != nil {
			return err
		}
		if err := articleTagRefModel.deleteArticleTagRefs(toDeleteRefsCondition, tx); err != nil {
			return err
		}
		return nil
	})
}

func DeleteArticles(condition ArticleCondition, keyword string) error {
	articles, _ := articleModel.getArticles(condition, keyword, nil)
	articleIds := []int{}
	for _, article := range articles {
		articleIds = append(articleIds, article.Id)
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := articleModel.deleteArticles(condition, keyword, tx); err != nil {
			return err
		}
		if err := articleTagRefModel.deleteArticleTagRefs(ArticleTagRefCondition{ArticleIds: articleIds}, tx); err != nil {
			return err
		}
		return nil
	})
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
		condition.Ids = collection.Collect(condition.Ids).Unique().ToIntArray()
	}
	tags, count = tagModel.getTags(condition, keyword, p)
	return
}

func CreateTags(tags []*Tag) error {
	return tagModel.createTags(tags, nil)
}

func UpdateTags(tags []Tag) error {
	return tagModel.updateTags(tags, nil)
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
