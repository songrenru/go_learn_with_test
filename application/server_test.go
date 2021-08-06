package poker

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"reflect"
	"io"
)

func TestServer(t *testing.T) {// 测试handler方法
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
            "Floyd":  10,
		},
		nil,
		nil,
	}
	// server := &PlayerServer{&store}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) { 
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		// PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		// PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	// server := &PlayerServer{&store}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("wang '%d' but got '%d'", 1, len(store.winCalls))
		}
	})

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper1"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 2 {
			t.Errorf("wang '%d' but got '%d'", 1, len(store.winCalls))
		}

		if store.winCalls[1] != player {
			t.Errorf("wang '%s' but got '%s'", player, store.winCalls[1])
		}
	})
}

func TestLeague(t *testing.T) {

	t.Run("it return 200 on /league", func(t *testing.T) {
		store := StubPlayerStore{}
		// server := &PlayerServer{&store}
		server := NewPlayerServer(&store)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
            {"Chris", 20},
            {"Tiest", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeague,}
		server := NewPlayerServer(&store)
		
		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

        assertLeague(t, got, wantedLeague)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/" + name, nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/" + name, nil)
	return request
}

func newGetLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wang '%s' but got '%s'", want, got)
	}	
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wang '%d' but got '%d'", want, got)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

const jsonContentType = "application/json"

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.HeaderMap)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
        t.Fatalf ("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
    }

    return
}