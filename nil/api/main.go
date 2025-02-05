package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type Song struct {
	Group       string
	SongName    string
	SongDetails SongDetails
}

type SongDetails struct {
	ReleaseDate string `json:"ReleaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

var arr = []Song{a1, a2, a3, a4, a5, a6}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/info", foo).Methods("GET")

	http.ListenAndServe(":8081", r)
}

func foo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		slog.Any("ERROR", err)
	}

	group := r.Form.Get("group")
	song := r.Form.Get("song")

	for i := 0; i < len(arr); i++ {
		if arr[i].Group == group && arr[i].SongName == song {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(arr[i])
		}
	}
}

var a1 = Song{
	Group:    "Sokoninaru",
	SongName: "Metalin",
	SongDetails: SongDetails{
		ReleaseDate: "15.05.2018",
		Text:        "Metalin, boku no tamashii\nKowakunai, ikite iru\nKono sekai de, boku wa ikiru",
		Link:        "https://www.youtube.com/watch?v=1",
	},
}

var a2 = Song{
	Group:    "Sokoninaru",
	SongName: "Break out!!!",
	SongDetails: SongDetails{
		ReleaseDate: "22.11.2019",
		Text:        "Break out!!!, boku wa tobitatsu\nKowakunai, yume o shinjiru\nKono jikan ni, subete o kakete",
		Link:        "https://www.youtube.com/watch?v=2",
	},
}

var a3 = Song{
	Group:    "Muse",
	SongName: "SuperMassive",
	SongDetails: SongDetails{
		ReleaseDate: "16.07.2006",
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	},
}

var a4 = Song{
	Group:    "A Crowd of Rebellion",
	SongName: "Crocus",
	SongDetails: SongDetails{
		ReleaseDate: "18.09.2018",
		Text:        "Crocus, boku no yume\nKimi to issho ni, ikite iru\nKowakunai, boku wa mamoru",
		Link:        "https://www.youtube.com/watch?v=4",
	},
}

var a5 = Song{
	Group:    "A Crowd of Rebellion",
	SongName: "ZENITH",
	SongDetails: SongDetails{
		ReleaseDate: "05.08.2020",
		Text:        "ZENITH, boku wa kakeru\nKowakunai, yume o shinjiru\nKono toki ni, subete o shinjiru",
		Link:        "https://www.youtube.com/watch?v=5",
	},
}

var a6 = Song{
	Group:    "A Crowd of Rebellion",
	SongName: "Hello, Mr. Judgement.",
	SongDetails: SongDetails{
		ReleaseDate: "12.01.2021",
		Text:        "Hello, Mr. Judgement, boku wa ikiru\nKowai kedo, boku wa tobitatsu\nKono jikan ni, subete o kakete",
		Link:        "https://www.youtube.com/watch?v=6",
	},
}
