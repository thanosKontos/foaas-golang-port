package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var routesOneParamToMessage = map[string]string{
	"/awesome/{first}":     "This is Fucking Awesome.",
	"/asshole/{first}":     "Fuck you, asshole.",
	"/bag/{first}":         "Eat a bag of fucking dicks.",
	"/because/{first}":     "Why? Because fuck you, that's why.",
	"/bucket/{first}":      "Please choke on a bucket of cocks.",
	"/bye/{first}":         "Fuckity bye!",
	"/cool/{first}":        "Cool story, bro.",
	"/cup/{first}":         "How about a nice cup of shut the fuck up?",
	"/diabetes/{first}":    "I'd love to stop and chat to you but I'd rather have type 2 diabetes.",
	"/everyone/{first}":    "Everyone can go and fuck off.",
	"/everything/{first}":  "Fuck everything.",
	"/family/{first}":      "Fuck you, your whole family, your pets, and your feces.",
	"/fascinating/{first}": "Fascinating story, in what chapter do you shut the fuck up?",
	"/flying/{first}":      "I don't give a flying fuck.",
	"/fyyff/{first}":       "Fuck you, you fucking fuck.",
	"/give/{first}":        "I give zero fucks.",
	"/horse/{first}":       "Fuck you and the horse you rode in on.",
	"/immensity/{first}":   "You can not imagine the immensity of the FUCK I do not give.",
	"/life/{first}":        "Fuck my life.",
	"/looking/{first}":     "Looking for a fuck to give.",
	"/maybe/{first}":       "Maybe. Maybe not. Maybe fuck yourself.",
	"/me/{first}":          "Fuck me.",
	"/mornin/{first}":      "Happy fuckin' mornin'!",
	"/no/{first}":          "No fucks given.",
	"/pink/{first}":        "Well, fuck me pink.",
	"/programmer/{first}":  "Fuck you, I'm a programmer, bitch!",
	"/retard/{first}":      "You Fucktard!",
	"/ridiculous/{first}":  "That's fucking ridiculous",
	"/rtfm/{first}":        "Read the fucking manual!",
	"/sake/{first}":        "For Fuck's sake!",
	"/shit/{first}":        "Fuck this shit!",
	"/single/{first}":      "Not a single fuck was given.",
	"/thanks/{first}":      "Fuck you very much.",
	"/that/{first}":        "Fuck that.",
	"/this/{first}":        "Fuck this.",
	"/too/{first}":         "Thanks, fuck you too.",
	"/tucker/{first}":      "Come the fuck in or fuck the fuck off.",
	"/what/{first}":        "What the fuck‽",
	"/zayn/{first}":        "Ask me if I give a motherfuck ?!!",
	"/zero/{first}":        "Zero, thats the number of fucks I give.",
}

var routesTwoParamsToMessage = map[string]string{
	"/anyway/{first}/{second}": "Who the fuck are you anyway, %s, why are you stirring up so much trouble, and, who pays you?",
	"/back/{first}/{second}": "%s, back the fuck off.",
	"/bday/{first}/{second}": "Happy Fucking Birthday, %s.",
	"/blackadder/{first}/{second}": "%s, your head is as empty as a eunuch’s underpants. Fuck off!",
	"/bus/{first}/{second}": "Christ on a bendy-bus, %s, don't be such a fucking faff-arse.",
	"/caniuse/{first}/{second}": "Can you use %s? Fuck no!",
	"/chainsaw/{first}/{second}": "Fuck me gently with a chainsaw, %s. Do I look like Mother Teresa?",
	"/cocksplat/{first}/{second}": "Fuck off %s, you worthless cocksplat",
	"/dalton/{first}/{second}": "A fucking problem solving super-hero.",
	"/deraadt/{first}/{second}": "%s you are being the usual slimy hypocritical asshole... You may have had value ten years ago, but people will see that you don't anymore.",
	"/donut/{first}/{second}": "%s, go and take a flying fuck at a rolling donut.",
	"/gfy/{first}/{second}": "Golf foxtrot yankee, %s.",
	"/ing/{first}/{second}": "Fucking fuck off, %s.",
	"/keep/{first}/{second}": "%s: Fuck off. And when you get there, fuck off from there too. Then fuck off some more. Keep fucking off until you get back here. Then fuck off again.",
	"/keepcalm/{first}/{second}": "Keep the fuck calm and %s",
	"/king/{first}/{second}": "Oh fuck off, just really fuck off you total dickface. Christ, %s, you are fucking thick.",
	"/linus/{first}/{second}": "%s, there aren't enough swear-words in the English language, so now I'll have to call you perkeleen vittupää just to express my disgust and frustration with this crap.",
	"/look/{first}/{second}": "%s, do I look like I give a fuck?",
	"/nugget/{first}/{second}": "Well %s, aren't you a shining example of a rancid fuck-nugget.",
	"/off/{first}/{second}": "Fuck off, %s.",
	"/off-with/{first}/{second}": "Fuck off with %s",
	"/outside/{first}/{second}": "%s, why don't you go outside and play hide-and-go-fuck-yourself?",
	"/particular/{first}/{second}": "Fuck this %s in particular.",
	"/problem/{first}/{second}": "What the fuck is your problem %s?",
	"/pulp/{first}/{second}": ":language, motherfucker, do you speak it?",
	"/shakespeare/{first}/{second}": "%s, Thou clay-brained guts, thou knotty-pated fool, thou whoreson obscene greasy tallow-catch!",
	"/shutup/{first}/{second}": "%s, shut the fuck up.",
	"/thinking/{first}/{second}": "%s, what the fuck were you actually thinking?",
	"/thumbs/{first}/{second}": "Who has two thumbs and doesn't give a fuck? %s.",
	"/xmas/{first}/{second}": "Merry Fucking Christmas, %s.",
	"/yoda/{first}/{second}": "Fuck off, you must, %s.",
	"/you/{first}/{second}": "Fuck you, %s.",
}

var routesThreeParamsToMessage = map[string]string{
	"/dosomething/{first}/{second}/{third}": "%s the fucking %s",
	"/ballmer/{first}/{second}/{third}": "Fucking %s is a fucking pussy. I'm going to fucking bury that guy, I have done it before, and I will do it again. I'm going to fucking kill %s.",
}

// OneParameterHandler is a handler for all routes with one param
func OneParameterHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	apiResponse := buildAPIResponseOneParameter(r)

	if contentType == "application/json" {
		json.NewEncoder(w).Encode(apiResponse)
		return
	}

	t, _ := template.New("api template").ParseFiles("api/tmpl/api.html")
	t.ExecuteTemplate(w, "api.html", apiResponse)
}

// TwoParametersHandler is a handler for all routes with 2 params
func TwoParametersHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	apiResponse := buildAPIResponseTwoParameters(r)

	if contentType == "application/json" {
		json.NewEncoder(w).Encode(apiResponse)
		return
	}

	t, _ := template.New("api template").ParseFiles("api/tmpl/api.html")
	t.ExecuteTemplate(w, "api.html", apiResponse)
}

// ThreeParametersHandler is a handler for all routes with 3 params
func ThreeParametersHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	apiResponse := buildAPIResponseThreeParameters(r)

	if contentType == "application/json" {
		json.NewEncoder(w).Encode(apiResponse)
		return
	}

	t, _ := template.New("api template").ParseFiles("api/tmpl/api.html")
	t.ExecuteTemplate(w, "api.html", apiResponse)
}

func buildAPIResponseOneParameter(r *http.Request) apiResponse {
	matchedRoute, _ := mux.CurrentRoute(r).GetPathTemplate()
	params := mux.Vars(r)

	return apiResponse{
		routesOneParamToMessage[matchedRoute],
		fmt.Sprintf("- %s", params["first"]),
	}
}

func buildAPIResponseTwoParameters(r *http.Request) apiResponse {
	matchedRoute, _ := mux.CurrentRoute(r).GetPathTemplate()
	params := mux.Vars(r)

	return apiResponse{
		fmt.Sprintf(routesTwoParamsToMessage[matchedRoute], params["first"]),
		fmt.Sprintf("- %s", params["second"]),
	}
}

func buildAPIResponseThreeParameters(r *http.Request) apiResponse {
	matchedRoute, _ := mux.CurrentRoute(r).GetPathTemplate()
	params := mux.Vars(r)

	return apiResponse{
		fmt.Sprintf(routesThreeParamsToMessage[matchedRoute], params["first"], params["second"]),
		fmt.Sprintf("- %s", params["third"]),
	}
}
