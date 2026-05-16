package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Player holds player response
type Player struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	HighScore      int      `json:"highScore"`
	IsOnline       bool     `json:"isOnline"`
	Location       string   `json:"location"`
	LevelsUnlocked []string `json:"levelsUnlocked"`
}

// players 模拟从数据库查询后的数据
var players = []Player{
	{
		ID:        123,
		Name:      "Pablo",
		HighScore: 1100,
		IsOnline:  true,
		Location:  "Italy",
	},
	{
		ID:        230,
		Name:      "Dora",
		HighScore: 2100,
		IsOnline:  false,
		Location:  "Germany",
	},
}

func main() {
	// Schema
	playerFields := graphql.Fields{
		"id":             &graphql.Field{Type: graphql.Int},
		"name":           &graphql.Field{Type: graphql.String},
		"highScore":      &graphql.Field{Type: graphql.Int},
		"isOnline":       &graphql.Field{Type: graphql.Boolean},
		"location":       &graphql.Field{Type: graphql.String},
		"levelsUnlocked": &graphql.Field{Type: graphql.NewList(graphql.String)},
	}

	// 根查询: 定义了查询时的根对象
	playerQuery := graphql.ObjectConfig{Name: "Player", Fields: playerFields}

	// 根对象
	playerObject := graphql.NewObject(playerQuery)

	// Schema
	playersFields := graphql.Fields{
		"players": &graphql.Field{
			Type:        graphql.NewList(playerObject),
			Description: "All players",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return players, nil // players 模拟从数据库查询后的数据
			},
		},
	}
	playersQuery := graphql.ObjectConfig{Name: "Players", Fields: playersFields}
	playersObject := graphql.NewObject(playersQuery)

	// 模式配置: 从模式配置创建一个新的模式
	schemaConfig := graphql.SchemaConfig{Query: playersObject}

	// 模式: 定义了 GraphQL 响应的结构
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // 一个交互式的 GraphQL 编辑器, 用于启用暴露的 API 的文档
	})

	http.Handle("/graphql", h)
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		log.Fatalf("failed to start graphql server, error: %v", err)
	}
}
