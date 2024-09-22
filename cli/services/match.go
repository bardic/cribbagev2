package services

import (
	"net/http"
	"strconv"

	"github.com/bardic/cribbagev2/cli/state"
	tea "github.com/charmbracelet/bubbletea"
)

func GetPlayerMatch() tea.Msg {
	return url(EndPointMatch+"/?id="+strconv.Itoa(state.ActiveMatchId), http.MethodGet, "")
}

func PutPlayerMatch(id int) tea.Msg {
	return url(EndPointMatch, http.MethodPut, "")
}

func PostPlayerMatch() tea.Msg {
	return url(EndPointMatch, http.MethodPost, "")
}

func GetPlayerMatchCard() tea.Msg {
	return url(EndPointMatchCard, http.MethodPut, "")
}

func PutPlayerMatchCard(id int) tea.Msg {
	return url(EndPointMatchCard, http.MethodPut, "")
}

func PostPlayerMatchCard(ids []int) tea.Msg {
	return url(EndPointMatchCard, http.MethodPost, "")
}