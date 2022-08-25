package backend

import (
    "context"
    "fmt"

    "around/constants"

    "github.com/olivere/elastic/v7"
)

// 这里是一个全局变量
var (
    ESBackend *ElasticsearchBackend
)



// 创立数据库的连接，struct类比java当中的class
type ElasticsearchBackend struct {
    client *elastic.Client // build一个类似session factory的东西
    // 通过client操控数据库，类似于mySQL的session factory，只能初始化一次
    // 这里使用 * pointer类型是因为
    
}

// initialization: 初始化backend的信息
func InitElasticsearchBackend() {
    client, err := elastic.NewClient(
        elastic.SetURL(constants.ES_URL),
        elastic.SetBasicAuth(constants.ES_USERNAME, constants.ES_PASSWORD))
    if err != nil {
        panic(err)
    }

    exists, err := client.IndexExists(constants.POST_INDEX).Do(context.Background())
    if err != nil {
        panic(err)
    }

    if !exists {
        // index是目录，帮助做搜索上的优化
        // index : false 没有索引的优化，搜的时候还是一行一行搜，我们这里认为常用搜索的只有id user message，节省底层开销
        mapping := `{
            "mappings": {
                "properties": {
                    "id":       { "type": "keyword" },
                    "user":     { "type": "keyword" },
                    "message":  { "type": "text" },
                    "url":      { "type": "keyword", "index": false }, 
                    "type":     { "type": "keyword", "index": false }
                }
            }
        }`
        _, err := client.CreateIndex(constants.POST_INDEX).Body(mapping).Do(context.Background())
        if err != nil {
            panic(err)
        }
    }

    exists, err = client.IndexExists(constants.USER_INDEX).Do(context.Background())
    if err != nil {
        panic(err)
    }

    if !exists {
        mapping := `{
                        "mappings": {
                                "properties": {
                                        "username": {"type": "keyword"},
                                        "password": {"type": "keyword"},
                                        "age":      {"type": "long", "index": false},
                                        "gender":   {"type": "keyword", "index": false}
                                }
                        }
                }`
        _, err = client.CreateIndex(constants.USER_INDEX).Body(mapping).Do(context.Background())
        if err != nil {
            panic(err)
        }
    }
    fmt.Println("Indexes are created.")

    ESBackend = &ElasticsearchBackend{client: client} // 赋值
    // client: client 第一个client是struct里的field name，相当于ElasticserachBackend.client = client，go的特殊赋值方式
    // 这里不会重名，因为第一个client不算variable，只是一个field
    // 这里就相当于Java调用constructor
}

// 我们需要实现两个功能，search by user, search by keyword
// 所以我们只需要用query和index（DB）
func (backend *ElasticsearchBackend) ReadFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
    searchResult, err := backend.client.Search().
        Index(index).
        Query(query).
        Pretty(true).
        Do(context.Background())
    if err != nil {
        return nil, err
    }

    return searchResult, nil
}

func (backend *ElasticsearchBackend) SaveToES(i interface{}, index string, id string) error {
    _, err := backend.client.Index().
        Index(index).
        Id(id).
        BodyJson(i). // 为什么传进来的是interface，不是POST对象？因为我们之后想复用，这样除了POST还可以插入其他的数据类型
        Do(context.Background())
    return err
}

func (backend *ElasticsearchBackend) DeleteFromES(query elastic.Query, index string) error {
    _, err := backend.client.DeleteByQuery().
        Index(index).
        Query(query).
        Pretty(true).
        Do(context.Background())

    return err
}