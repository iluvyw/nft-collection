package collection

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/0xanpham/nft-collection/graph/model"
	"github.com/0xanpham/nft-collection/redis"
)

type CollectionService struct {
		db *redis.DB
		redisIndex int
}

func Init() *CollectionService {
		return &CollectionService{
				redisIndex: 1,
				db: redis.Connect(1),
		}
}

func (collectionService *CollectionService) Create(newCollection model.NewCollection) *model.Collection{
    id := strconv.Itoa(collectionService.db.Size(collectionService.redisIndex))

    name := newCollection.Name
    authorId := newCollection.AuthorID
    nfts := newCollection.Nfts

    collection := &model.Collection{
        ID: id,
        Name: name,
        AuthorID: authorId,
        NftIds: nfts,
    }

    nftIds, error := json.Marshal(collection.NftIds)
    if error != nil {
        panic(error)
    }

    data := map[string]interface{}{
        "id": collection.ID,
        "name": collection.Name,
        "authorId": collection.AuthorID,
        "nfts": nftIds,
    }

		out, error := json.Marshal(data)
		if error != nil {
			panic(error)
		}

		fmt.Println(id)

    collectionService.db.Set(id, string(out))
    
    return collection
}