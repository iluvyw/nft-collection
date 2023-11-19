package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/0xanpham/nft-collection/graph/model"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	client *redis.Client
}

func Connect() *DB {
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
    return &DB{
        client: client,
    }
}

func (db *DB) CreateNft(newNft model.NewNft) *model.Nft{
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
    id := strconv.FormatInt(db.client.DBSize(ctx).Val() + 1, 10)

    address := newNft.Address
    tokenId := newNft.TokenID
    tokenURI := "https://token.uri.com"

    nft := &model.Nft{
        ID: id,
        Address: address,
        TokenID: tokenId,
        TokenURI: tokenURI,
    }

    db.client.HSet(
        ctx,
        fmt.Sprintf("nft:%s",id),
        "id", nft.ID,
        "address", nft.Address,
        "tokenId", nft.TokenID,
        "tokenUri", nft.TokenURI,
    )
    
    return nft
}

func (db *DB) CreateCollection(newCollection model.NewCollection) *model.Collection{
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
    id := strconv.FormatInt(db.client.DBSize(ctx).Val() + 1, 10)

    name := newCollection.Name
    authorId := newCollection.AuthorID
    nfts := newCollection.Nfts

    collection := &model.Collection{
        ID: id,
        Name: name,
        AuthorID: authorId,
        NftIds: nfts,
    }

    db.client.HSet(
        ctx,
        fmt.Sprintf("collection:%s",id),
        "id", collection.ID,
        "name", collection.Name,
        "authorId", collection.AuthorID,
        "nfts", collection.NftIds,
    )
    
    return collection
}