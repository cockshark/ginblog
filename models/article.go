package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` // gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
	Tag   Tag `json:"tag"`                 // 实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

// ExistArticleByID 文章是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?").First(&article)

	return article.ID > 0

}

// GetArticleTotal 获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

/*
Preload是什么东西，为什么查询可以得出每一项的关联Tag？
Preload就是一个预加载器，它会执行两条SQL，
分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，
那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询

*/

// GetArticles 获取多个文章
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

/*  我们的Article是如何关联到Tag？？？
能够达到关联，首先是gorm本身做了大量的约定俗成
Article有一个结构体成员是TagID，就是外键。gorm 会通过类名+ID的方式去找到这两个类之间的关联关系
Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
*/

// GetArticle 获取单个文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

// EditArticle 编辑修改文章
func EditArticle(id int, data interface{}) {
	db.Model(&Article{}).Where("id = ?", id).Update(data)

}

/*
v.(I) 是什么？
v表示一个接口值，I表示接口类型。这个实际就是Go中的类型断言，
用于判断一个接口值的实际类型是否为某个类型，或一个非接口值的类型是否实现了某个接口类型
*/

// AddArticle 新增文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

// DeleteArticle 删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
