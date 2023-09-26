GraphQL client

https://github.com/shurcooL/graphql

```go
package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shurcooL/graphql"
	"os"
)
func main() {
	// {
	//  betResults(first: 5) {
	//    id
	//    requestId
	//    callerAddress
	//    gameResult
	//  }
	// }
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	client := graphql.NewClient(os.Getenv("GRAPH_ENDPOINT_DEV"), nil)
	var query struct {
		BetResults []struct {
			Id            graphql.String
			RequestId     graphql.String
			CallerAddress graphql.String
			GameResult    graphql.Boolean
		} `graphql:"betResults(where: {gameResult: true})"`
	}

	err = client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
		fmt.Println(err)
	}
	fmt.Println(query.BetResults[0].RequestId)
	fmt.Println(query.BetResults[0].CallerAddress)
	fmt.Println(query.BetResults[0].GameResult)
	fmt.Println("总共查到", len(query.BetResults), "条数据")
}
```