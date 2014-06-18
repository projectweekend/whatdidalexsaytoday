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
