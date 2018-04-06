package wxspider

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/sda/wechat"
)

//SpiderArticle 采集文章并保存到本地
func SpiderArticle(urlStr string) error {
	var a Article
	article, err := wechat.Find(urlStr)
	if err == nil {

		if article.URL == "" {
			return errors.New("不支持该链接！")
		}

		a.GetArticleByURL(article.URL)
		a.AppID = article.AppID
		a.AppName = article.AppName
		a.RoundHead = article.RoundHead
		a.OriHead = article.OriHead
		a.URL = article.URL
		a.Title = article.Title
		a.Intro = article.Intro
		a.Cover = article.Cover
		a.Author = article.Author
		a.PubAt = article.PubAt
		a.Save()
	}
	return nil
}

//PublishArticle 采集文章并保存到本地
func PublishArticle() error {
	// if post.State == 0 { // 检查提交状态
	var a Article

	rows := a.GetPlanPublushArticle()
	for _, row := range rows {
		PostArticle(row)
	}

	// article, err := wechat.Find(urlStr)
	// if err == nil {

	// 	if article.URL == "" {
	// 		return errors.New("不支持该链接！")
	// 	}

	// 	a.GetArticleByURL(article.URL)
	// 	a.AppID = article.AppID
	// 	a.AppName = article.AppName
	// 	a.RoundHead = article.RoundHead
	// 	a.OriHead = article.OriHead
	// 	a.URL = article.URL
	// 	a.Title = article.Title
	// 	a.Intro = article.Intro
	// 	a.Cover = article.Cover
	// 	a.Author = article.Author
	// 	a.PubAt = article.PubAt
	// 	// i64, err := strconv.ParseInt(article.PubAt, 10, 64)
	// 	// if err != nil {
	// 	// 	// fmt.Println(err)
	// 	// 	return errors.New("时间转化失败")
	// 	// }
	// 	// // a.PublishAt = time.Unix(i64, 0)
	// 	// a.PubAt = i64

	// 	// panic(a.ID)

	// 	a.Save()
	// 	// fmt.Println(a)
	// }
	// }
	return nil
}

//PostArticle 采集文章并保存到本地
func PostArticle(article Article) error {
	client := http.Client{}
	data := make(url.Values)
	data["title"] = []string{article.Title}
	resp, err := client.PostForm("http://wxapi.oo/api/v1/article", data)
	if err != nil {
		log.Printf("登录时提交数据异常")
	}
	formPost, err := goquery.NewDocumentFromReader(resp.Body)

	resp.Body.Close()

	postMsg, err := formPost.Html()
	// panic(err)
	log.Println(" %s  ", postMsg)
	return nil
}