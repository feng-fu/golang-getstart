package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		nil,
	}
	server := NewPlayerServer(&store)

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
		nil,
	}

	server := NewPlayerServer(&store)

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

		assertDeepEqual(t, got, want)
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{nil, nil, wantedLeague}
	server := NewPlayerServer(&store)

	t.Run("it returns the league table as JSON", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertDeepEqual(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
	})
}

// func TestFileSystemStore(t *testing.T) {
// 	t.Run("/league from a reader", func(t *testing.T) {
// 		database := strings.NewReader(`[
// 			{"Name": "Cleo", "Wins": 10},
// 			{"Name": "Chris", "Wins": 33}
// 		]`)

// 		store := FileSystemStore{database}
// 		got := store.GetLeague()

// 		want := []Player{
// 			{"Cleo", 10},
// 			{"Chris", 33},
// 		}
// 		assertLeague(t, got, want)
// 	})
// }

func newScoreRequest(name string, method string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got '%d' want '%d'", got, want)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	var got []Player

	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("UNable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return got
}

func assertDeepEqual(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != "application/json" {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}
}
