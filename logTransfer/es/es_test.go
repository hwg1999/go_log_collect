package es

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestEs(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")

	// 新增记录
	p1 := User{Name: "jerry", Age: 30, Email: "hzde0128@live.cn"}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
