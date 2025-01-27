package match

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bardic/gocrib/queries/queries"

	conn "github.com/bardic/gocrib/server/db"
	"github.com/bardic/gocrib/server/route/helpers"

	"github.com/labstack/echo/v4"
)

// Return match details if the player is able to join the match
//
//	@Summary	Join match by id
//	@Description
//	@Tags		match
//	@Accept		json
//	@Produce	json
//	@Param		matchId	path		int	true	"match id"'
//	@Success	200		{object}	vo.MatchDetailsResponse
//	@Failure	400		{object}	error
//	@Failure	404		{object}	error
//	@Failure	500		{object}	error
//	@Router		/match/{matchId}/deal [put]
func Deal(c echo.Context) error {
	matchId, err := strconv.Atoi(c.Param("matchId"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db := conn.Pool()
	defer db.Close()
	q := queries.New(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	players, err := q.GetPlayersInMatch(ctx, &matchId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	numOfPlayers := len(players)

	cardsToDealPerPlayer := 6
	if numOfPlayers > 2 {
		cardsToDealPerPlayer = 5
	}

	// Deal cards

	cards, err := q.GetMatchCards(ctx, &matchId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for i := 0; i < numOfPlayers; i++ {
		for j := 0; j < cardsToDealPerPlayer; j++ {
			card := cards[(i*cardsToDealPerPlayer)+j]
			q.UpdateMatchCardState(ctx, queries.UpdateMatchCardStateParams{
				State:     queries.CardstateHand,
				Origowner: players[i].ID,
				Currowner: players[i].ID,
				ID:        card.Card.ID,
			})
		}
	}

	match, err := helpers.GetMatch(&matchId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, match)
}