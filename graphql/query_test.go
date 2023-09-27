package graphql

import (
	"context"
	"fmt"
	"github.com/shurcooL/graphql"
	"testing"
)

func TestQuery(t *testing.T) {
	// {
	//  betResults(first: 5) {
	//    id
	//    requestId
	//    callerAddress
	//    gameResult
	//  }
	// }

	client := graphql.NewClient("https://api.studio.thegraph.com/query/33456/oddeven/version/latest", nil)
	var query struct {
		BetResults []struct {
			Id             graphql.String
			Amount         graphql.String
			RequestId      graphql.String
			CallerAddress  graphql.String
			GameResult     graphql.Boolean
			BlockTimestamp graphql.String
		} `graphql:"betResults(first:5 where: {callerAddress: $addr})"`
	}
	vars := map[string]any{
		"addr": "0x11924b8135a636ee14f65e338ceaab8917b68b7b",
	}
	err := client.Query(context.Background(), &query, vars)
	if err != nil {
		// Handle error.
		fmt.Println(err)
	}
	fmt.Println(query.BetResults[0].RequestId)
	fmt.Println(query.BetResults[0].CallerAddress)
	fmt.Println(query.BetResults[0].GameResult)
	fmt.Println(query.BetResults[0].Amount)
	fmt.Println("总共查到", len(query.BetResults), "条数据")
}
