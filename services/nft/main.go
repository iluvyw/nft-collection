package nft

import (
	"encoding/json"
	"strconv"

	"github.com/0xanpham/nft-collection/graph/model"
	"github.com/0xanpham/nft-collection/redis"
)

type NftService struct {
		db *redis.DB
		redisIndex int
}

func Init() *NftService {
		return &NftService{
				redisIndex: 0,
				db: redis.Connect(0),
		}
}

func (nftService *NftService) Create(newNft model.NewNft) *model.Nft{
    id := strconv.Itoa(nftService.db.Size(nftService.redisIndex))

    address := newNft.Address
    tokenId := newNft.TokenID
    tokenURI := "https://token.uri.com"

    nft := &model.Nft{
        ID: id,
        Address: address,
        TokenID: tokenId,
        TokenURI: tokenURI,
    }

		data := map[string]interface{}{
			"id": id,
			"address": address,
			"tokenId": tokenId,
			"tokenUri": tokenURI,
		}

    out, error := json.Marshal(data)
		if error != nil {
			panic(error)
		}

		nftService.db.Set(string(id), string(out))
    
    return nft
}

// func (db *DB) getNfts() *[]model.Nft {
//     ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
//     id := strconv.FormatInt(db.client.DBSize(ctx).Val() + 1, 10)
//     nfts := []model.Nft{}
//     // nft := db.client.HGetAll(ctx, fmt.Sprintf("nft:%s",id))
    
//     return &nfts
// }