//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/0xanpham/nft-collection/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	collections []*model.Collection
	users []*model.User
	nfts []*model.Nft
}
