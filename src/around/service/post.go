package service

import (
	"mime/multipart"
    "reflect"

    "around/backend"
    "around/constants"
    "around/model"

    "github.com/olivere/elastic/v7"
)

func SearchPostsByUser(user string) ([]model.Post, error) {
    query := elastic.NewTermQuery("user", user)
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)
    if err != nil {
        return nil, err
    }
    return getPostFromSearchResult(searchResult), nil
}

func SearchPostsByKeywords(keywords string) ([]model.Post, error) {
    query := elastic.NewMatchQuery("message", keywords)
    query.Operator("AND")
    if keywords == "" {
        query.ZeroTermsQuery("all")
    }
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)
    if err != nil {
        return nil, err
    }
    return getPostFromSearchResult(searchResult), nil
}

func getPostFromSearchResult(searchResult *elastic.SearchResult) []model.Post {
    var ptype model.Post
    var posts []model.Post

    for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
        p := item.(model.Post)
        posts = append(posts, p)
    }
    return posts
}

func SavePost(post *model.Post, file multipart.File) error {
    medialink, err := backend.GCSBackend.SaveToGCS(file, post.Id)
    if err != nil {
        return err
    }
    post.Url = medialink

    err = backend.ESBackend.SaveToES(post, constants.POST_INDEX, post.Id)
    
    return err // 如果SaveToGCS成功，SaveToES失败，代表网盘里存了很多没人refer的文件（无效文件）
    // 对用户使用网站没有影响，影响后台运行成本，如何解决？？没有现成的library（想要达到mysql的transaction的效果）
}

func DeletePost(id string, user string) error {
    query := elastic.NewBoolQuery()
    query.Must(elastic.NewTermQuery("id", id))
    query.Must(elastic.NewTermQuery("user", user))

    return backend.ESBackend.DeleteFromES(query, constants.POST_INDEX)
}
