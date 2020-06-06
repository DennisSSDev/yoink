package epicstorecommunicatorserver

import (
	"context"

	"github.com/twitchtv/twirp"

	"github.com/dennisssdev/yoink/graphql/query"
	pb "github.com/dennisssdev/yoink/rpc/epicstorecommunicator"
)

/*
PingForFreeGames makes a graphql call to the epic store api
and asks for data about the latest free games.
It returns the names of the games and a bit of data
that partains to the promotion
*/
func (s *Server) PingForFreeGames(ctx context.Context, empty *pb.EmptyReq) (resp *pb.FreeEpicStoreGamesResp, err error) {
	games := []*pb.EpicStoreGame{}
	result := &pb.FreeEpicStoreGamesResp{
		Games: games,
	}
	if gqlError := s.GraphQLClient.Query(
		ctx,
		&query.EpicStoreFreeGames,
		nil,
	); gqlError != nil {
		return result,
			twirp.NewError(
				twirp.Unavailable,
				"The Epic Store query has resulted in an error: "+gqlError.Error(),
			)
	}

	elemLength := len(query.EpicStoreFreeGames.Catalog.SearchStore.Elements)
	gameQueryElements := query.EpicStoreFreeGames.Catalog.SearchStore.Elements

	for i := 0; i < elemLength; i++ {
		if gameQueryElements[i].Title == "Mystery Game" {
			// this is not an actual game move on
			continue
		}

		gameItem := gameQueryElements[i]
		gamePromotion := gameItem.Promotions.PromotionalOffers

		if len(gamePromotion) == 0 ||
			len(gamePromotion[0].PromotionalOffers) == 0 {
			// this is not a timed promotion game, probably free to play
			continue
		}

		games = append(games,
			&pb.EpicStoreGame{
				Name:         string(gameItem.Title),
				ImageUrl:     string(gameItem.KeyImages[i].URL),
				PromoEndDate: string(gameItem.Promotions.PromotionalOffers[0].PromotionalOffers[0].EndDate),
			},
		)
	}
	if len(games) == 0 {
		return result,
			twirp.NewError(
				twirp.NotFound, "Epic Store doesn't have any free promotions",
			)
	}
	result.Games = games
	return result, nil
}
