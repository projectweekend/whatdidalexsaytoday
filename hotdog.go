package hotdog

import (
	"fmt"
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
		"I've got to make my kilt payment"
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(thatsWhatHeSaid))

	fmt.Fprintf(w, fmt.Sprintf(`
		<html>
			<style type="text/css">
				div {
					text-align: center;
					font-family: Arial,sans-serif;
					font-weight: bold;
					font-size: 52pt
				}
			</style>
			<div>%s</div>
		</html>`, thatsWhatHeSaid[index]))
}

func init() {
	http.HandleFunc("/", handler)
}
