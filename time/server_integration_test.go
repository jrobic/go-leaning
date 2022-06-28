package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, `[]`)
	defer cleanDatabase()

	store, _ := NewFileSystemPlayerStore(database)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(player))
		AssertStatus(t, response.Code, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		AssertLeague(t, got, want)
	})
}

func BenchmarkGetLeague(b *testing.B) {
	database, cleanDatabase := CreateTempFile(b, `[
	{"Name": "Cleo", "Wins": 10},
	{"Name": "Chris", "Wins": 33}
	]`)

	defer cleanDatabase()

	store, _ := NewFileSystemPlayerStore(database)
	server := NewPlayerServer(store)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
	}
}
