package auction

import (
	"context"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"

	"go.mongodb.org/mongo-driver/bson"
)

func (ar *AuctionRepository) UpdateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	logger.Info("Updating auction with id = " + auctionEntity.Id)

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.ReplaceOne(ctx, bson.M{"_id": auctionEntity.Id}, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to update auction", err)
		return internal_error.NewInternalServerError("Error trying to update auction")
	}

	return nil
}
