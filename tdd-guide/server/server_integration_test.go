package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newScoreRequest(player, http.MethodGet))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)

		want := []Player{
			{"Pepper", 3},
		}

		assertLeague(t, got, want)
	})
}
