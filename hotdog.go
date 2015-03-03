package hotdog

import (
	"github.com/gorilla/mux"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	thatsWhatHeSaid = []string{
		"My chiropractor wanted to massage my glutes with a large spoon.",
		"Miley Cyrus is generally accepted as attractive.",
		"Bieber is a decent looking guy.",
		"I've got man hands.",
		"You gotta get some scented candles, bro.",
		"That's actually very flattering, thank you.",
		"Casinos and gay bars. Last night got weird.",
		"Hey, fuck egg shells. If one of those little fuckers falls in your scrambled eggs it's almost impossible to fish it out.",
		"My evenings are pretty booked up.",
		"Follow my live tweets bro?",
		"I've got to make my kilt payment.",
		"Oh, doing some lunges here? I'll do a couple.",
		"I do it on the train all the time.",
		"Aw man, I do not like underwear.",
		"#artstudent",
		"Holy shit! This dress is so expensive!",
		"What's the big deal with selfies?",
		"Everyone wants to talk about me.",
		"Scoops, y'all.",
		"Feel better, boo.",
		"Ruby FTW!",
		"I've had balls in my mouth before",
		"I lose mad 'el bees' when I cut gluten out of my diet.",
		"Yeah, I once dated a girl who was really into The Lost Boys.",
		"Is there a Mr. Virginia competition or something?",
		"I did zumba with my mom last week.",
		"When I was a kid, my dad used to wake me up and say, 'Wake up! Wake up! Today is opportunity day, and today you have the opportunity to succeed!'",
		"Ever since the accident of 2012, I don't arm wrestle any more.",
		"I'm actually quite the pie aficionado.",
		"I prefer to have my shirt off all the time [in public].",
		"The bro language is all about mouth-feel.",
		"I've got a bro injury–I messed up my hand with too many high fives.",
		"If I bring in breakfast for you guys, it won't be donuts—it'll be salad.",
		"Alright, I'll touch your wood.",
		"What are you, some kind of robit?",
		"Solar jokes are only second to lunar jokes.",
		"I guess I don't have weak wrists like you guys.",
		"I have that same shirt but without the sleeves.",
	}
)

func render(w http.ResponseWriter, i int) {
	tmpl, err := template.New("index.html").ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, thatsWhatHeSaid[i])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saying(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if id > len(thatsWhatHeSaid) || err != nil {
		http.Redirect(w, r, "http://whatdidalexsaytoday.com", http.StatusFound)
		return
	}

	render(w, id)
}

func index(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	render(w, rand.Intn(len(thatsWhatHeSaid)))
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", saying)
	r.HandleFunc("/", index)
	http.Handle("/", r)
}
