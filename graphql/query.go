package graphql

import (
	"context"
	"github.com/shurcooL/graphql"
)

type Queryer interface {
	Query(ctx context.Context, query any, vars map[string]any) error
}

type graphqlQueryer struct {
	client *graphql.Client
}

func NewQueryer(client *graphql.Client) Queryer {
	return &graphqlQueryer{
		client: client,
	}
}

func (g *graphqlQueryer) Query(ctx context.Context, query any, vars map[string]any) error {
	return g.client.Query(ctx, query, vars)
}
