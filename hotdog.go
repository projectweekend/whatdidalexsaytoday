package hotdog

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var (
	thatsWhatHeSaid = []string{
		"My chiropractor wanted to massage my glutes with a large spoon",
		"Miley Cyrus is generally accepted as attractive",
		"Bieber is a decent looking guy",
		"You gotta get some scented candles bro",
		"That's actually very flattering, thank you",
		"Casinos and gay bars. Last night got weird.",
		"Hey, fuck egg shells. If one of those little fuckers falls in your scrambled eggs it's almost impossible to fish it out.",
		"My evenings are pretty booked up",
		"Follow my live tweets bro?",
		"I've got to make my kilt payment",
		"Oh, doing some lunges here? I'll do a couple.",
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(thatsWhatHeSaid))

	tmpl, err := template.New("index.html").ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, thatsWhatHeSaid[index])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/", handler)
}
