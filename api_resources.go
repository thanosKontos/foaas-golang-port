package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiResponse struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

var routesFromToMessage = map[string]string{
	"/awesome/{from}":     "This is Fucking Awesome.",
	"/asshole/{from}":     "Fuck you, asshole.",
	"/bag/{from}":         "Eat a bag of fucking dicks.",
	"/because/{from}":     "Why? Because fuck you, that's why.",
	"/bucket/{from}":      "Please choke on a bucket of cocks.",
	"/bye/{from}":         "Fuckity bye!",
	"/cool/{from}":        "Cool story, bro.",
	"/cup/{from}":         "How about a nice cup of shut the fuck up?",
	"/diabetes/{from}":    "I'd love to stop and chat to you but I'd rather have type 2 diabetes.",
	"/everyone/{from}":    "Everyone can go and fuck off.",
	"/everything/{from}":  "Fuck everything.",
	"/family/{from}":      "Fuck you, your whole family, your pets, and your feces.",
	"/fascinating/{from}": "Fascinating story, in what chapter do you shut the fuck up?",
	"/flying/{from}":      "I don't give a flying fuck.",
	"/fyyff/{from}":       "Fuck you, you fucking fuck.",
	"/give/{from}":        "I give zero fucks.",
	"/horse/{from}":       "Fuck you and the horse you rode in on.",
	"/immensity/{from}":   "You can not imagine the immensity of the FUCK I do not give.",
	"/life/{from}":        "Fuck my life.",
	"/looking/{from}":     "Looking for a fuck to give.",
	"/maybe/{from}":       "Maybe. Maybe not. Maybe fuck yourself.",
	"/me/{from}":          "Fuck me.",
	"/mornin/{from}":      "Happy fuckin' mornin'!",
	"/no/{from}":          "No fucks given.",
	"/pink/{from}":        "Well, fuck me pink.",
	"/programmer/{from}":  "Fuck you, I'm a programmer, bitch!",
	"/retard/{from}":      "You Fucktard!",
	"/ridiculous/{from}":  "That's fucking ridiculous",
	"/rtfm/{from}":        "Read the fucking manual!",
	"/sake/{from}":        "For Fuck's sake!",
	"/shit/{from}":        "Fuck this shit!",
	"/single/{from}":      "Not a single fuck was given.",
	"/thanks/{from}":      "Fuck you very much.",
	"/that/{from}":        "Fuck that.",
	"/this/{from}":        "Fuck this.",
	"/too/{from}":         "Thanks, fuck you too.",
	"/tucker/{from}":      "Come the fuck in or fuck the fuck off.",
	"/what/{from}":        "What the fuck‽",
	"/zayn/{from}":        "Ask me if I give a motherfuck ?!!",
	"/zero/{from}":        "Zero, thats the number of fucks I give.",
}

var routesFromCompanyToMessage = map[string]string{
	"/anyway/{company}/{from}": "Who the fuck are you anyway, %s, why are you stirring up so much trouble, and, who pays you?",
}

func GetCompanyFromHandler(w http.ResponseWriter, r *http.Request) {
	matchedRoute, _ := mux.CurrentRoute(r).GetPathTemplate()
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(ApiResponse{
		fmt.Sprintf(routesFromCompanyToMessage[matchedRoute], params["company"]),
		fmt.Sprintf("- %s", params["from"]),
	})
}

func GetFromHandler(w http.ResponseWriter, r *http.Request) {
	matchedRoute, _ := mux.CurrentRoute(r).GetPathTemplate()
	fmt.Println(matchedRoute)
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(ApiResponse{
		routesFromToMessage[matchedRoute],
		fmt.Sprintf("- %s", params["from"]),
	})
}
