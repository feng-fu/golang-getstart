package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("return Pepper's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newScoreRequest("Pepper", http.MethodGet)
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newScoreRequest("Floyd", http.MethodGet)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newScoreRequest("JackLi", http.MethodGet)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestPost(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}

	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Lucy"
		request := newScoreRequest(player, http.MethodPost)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{map[string]int{
		// "Pepper": 0,
	}}
	server := PlayerServer{&store}
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, newScoreRequest(player, http.MethodGet))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}

func newScoreRequest(name string, method string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
