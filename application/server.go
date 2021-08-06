package poker

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	store PlayerStore
	// router  *http.ServeMux
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	// p := &PlayerServer{
	// 	store,
	// 	http.NewServeMux(),
	// }

	// p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	// p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// player := r.URL.Path[len("/players/"):]

// 	// switch r.Method {
// 	// case http.MethodGet:
// 	// 	p.showScore(w, player)
// 	// case http.MethodPost:
// 	// 	p.processWin(w, player)
// 	// }
		
	
// 	// router := http.NewServeMux()

// 	// router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	// router.Handle("/players/", http.HandlerFunc(p.playersHandler))

// 	// router.ServeHTTP(w, r)
	
// 	p.router.ServeHTTP(w, r)
// }

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	// leagueTable := []Player{
	// 	{"Chris", 999},
	// }
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.getLeagueTable())
}

func (p *PlayerServer) getLeagueTable() []Player {
	return p.store.GetLeague()
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := r.URL.Path[len("/players/"):]

// 	fmt.Fprint(w, GetPlayerScore(player))
// }

// func GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}

// 	if name == "Floyd" {
// 		return "10"
// 	}

// 	return "0"
// }

func (p *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := p.store.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, p.store.GetPlayerScore(name))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, name string) {
	p.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}