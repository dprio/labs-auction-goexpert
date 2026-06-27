package auction_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateAuction(t *testing.T) {

	t.Run("should create an auction successfully and close it", func(t *testing.T) {
		//given
		mockAuctionRepository := newMockauctionRepositoryInterface(t)
		mockBidRepository := newMockbidRepositoryInterface(t)
		target := NewAuctionUseCase(mockAuctionRepository, mockBidRepository)

		input := AuctionInputDTO{
			ProductName: "Product 1",
			Category:    "Category 1",
			Description: "Description 1",
			Condition:   0,
		}

		duration := 20 * time.Second

		var auc *auction_entity.Auction

		mockAuctionRepository.EXPECT().
			CreateAuction(context.Background(), mock.Anything).
			Return(nil).
			Once()

		mockAuctionRepository.EXPECT().
			UpdateAuction(mock.Anything, mock.Anything).
			Return(nil).
			Run(func(ctx context.Context, auction *auction_entity.Auction) {
				auc = auction
			}).
			Once()

		//when
		err := target.CreateAuction(context.Background(), input)
		assert.Nil(t, err)

		//then

		<-time.After(duration + 3*time.Second) // Wait for the auction to close

		assert.Equal(t, auction_entity.Completed, auc.Status)

		mockAuctionRepository.AssertExpectations(t)

	})

}
