package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// GenericErrorHandler is the handler for 404s and 405s
func GenericErrorHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(apiResponse{
		Message:  "622 - All The Fucks",
		Subtitle: "Server invites you to consider the truly monumental amount of fucks it couldn't give about your request.",
	})
}

// VersionHandler is the handler for version get requests
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	apiResponse := apiResponse{
		Message:  "Version 0.0.1",
		Subtitle: "FOAAS",
	}

	if contentType == "application/json" {
		json.NewEncoder(w).Encode(apiResponse)
		return
	}

	t, _ := template.New("api template").ParseFiles("api/tmpl/api.html")
	t.ExecuteTemplate(w, "api.html", apiResponse)
}

// IrregularRequestHandler is the handler for the following irregular requests
// /field/{first}/{second}/{third}
// /greed/{first}/{second}
// /madison/{first}/{second}
func IrregularRequestHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	params := mux.Vars(r)
	var response apiResponse

	switch matchedRoute, _ := mux.CurrentRoute(r).GetPathTemplate(); matchedRoute {
	case "/field/{first}/{second}/{third}":
		response = apiResponse{
			fmt.Sprintf("And %s said unto %s, 'Verily, cast thine eyes upon the field in which I grow my fucks', and %s gave witness unto the field, and saw that it was barren.", params["first"], params["second"], params["second"]),
			fmt.Sprintf("- %s", params["third"]),
		}
	case "/greed/{first}/{second}":
		response = apiResponse{
			fmt.Sprintf("The point is, ladies and gentleman, that %s -- for lack of a better word -- is good. %s is right. %s works. %s clarifies, cuts through, and captures the essence of the evolutionary spirit. %s, in all of its forms -- %s for life, for money, for love, knowledge -- has marked the upward surge of mankind.", params["first"], params["first"], params["first"], params["first"], params["first"], params["first"]),
			fmt.Sprintf("- %s", params["second"]),
		}
	case "/madison/{first}/{second}":
		response = apiResponse{
			fmt.Sprintf("What you've just said is one of the most insanely idiotic things I have ever heard, %s. At no point in your rambling, incoherent response were you even close to anything that could be considered a rational thought. Everyone in this room is now dumber for having listened to it. I award you no points %s, and may God have mercy on your soul.", params["first"], params["first"]),
			fmt.Sprintf("- %s", params["second"]),
		}
	}

	if contentType == "application/json" {
		json.NewEncoder(w).Encode(response)
		return
	}

	t, _ := template.New("api template").ParseFiles("api/tmpl/api.html")
	t.ExecuteTemplate(w, "api.html", response)
}

// OperationsHandler is the handler for operations get requests
func OperationsHandler(w http.ResponseWriter, r *http.Request) {
	var operations []operation

	operations = append(operations, operation{Name: "Who the fuck are you anyway", Url: "/anyway/:company/:from", Fields: []field{{Name: "Company", Field: "company"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Asshole", Url: "/asshole/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Awesome", Url: "/awesome/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Back the fuck off", Url: "/back/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Bag", Url: "/bag/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Ballmer", Url: "/ballmer/:name/:company/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "Company", Field: "company"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Bday", Url: "/bday/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Because", Url: "/because/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Blackadder", Url: "/blackadder/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Bravo Mike", Url: "/bm/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Bucket", Url: "/bucket/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Bus", Url: "/bus/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Bye", Url: "/bye/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Can I Use", Url: "/caniuse/:tool/:from", Fields: []field{{Name: "Tool", Field: "tool"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Chainsaw", Url: "/chainsaw/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Cocksplat", Url: "/cocksplat/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Cool Story", Url: "/cool/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Cup", Url: "/cup/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Dalton", Url: "/dalton/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "{name} you are being the usual slimy hypocritical asshole... You may have had value ten years ago, but people will see that you don't anymore.", Url: "/deraadt/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Diabetes", Url: "/diabetes/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Donut", Url: "/donut/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Do Something", Url: "/dosomething/:do/:something/:from", Fields: []field{{Name: "Do", Field: "do"}, {Name: "Something", Field: "something"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Everyone", Url: "/everyone/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Everything", Url: "/everything/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Family", Url: "/family/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fascinating", Url: "/fascinating/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Field of Fucks", Url: "/field/:name/:from/:reference", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}, {Name: "Reference", Field: "reference"}}})
	operations = append(operations, operation{Name: "Flying", Url: "/flying/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "FYYFF", Url: "/fyyff/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Golf Foxtrot Yankee", Url: "/gfy/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Give", Url: "/give/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Greed", Url: "/greed/:noun/:from", Fields: []field{{Name: "Noun", Field: "noun"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck you and the horse you rode in on", Url: "/horse/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Immensity", Url: "/immensity/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fucking", Url: "/ing/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Keep", Url: "/keep/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Keep Calm", Url: "/keepcalm/:reaction/:from", Fields: []field{{Name: "Reaction", Field: "reaction"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "King", Url: "/king/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Life", Url: "/life/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Linus", Url: "/linus/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Look", Url: "/look/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Looking", Url: "/looking/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Madison", Url: "/madison/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Maybe", Url: "/maybe/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck Me", Url: "/my/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "mornin", Url: "/mornin/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "No", Url: "/no/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Nugget", Url: "/nugget/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck Off", Url: "/off/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck Off With", Url: "/off-with/:behavior/:from", Fields: []field{{Name: "Behavior", Field: "behavior"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Outside", Url: "/outside/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "This Thing In Particular", Url: "/particular/:thing/:from", Fields: []field{{Name: "Thing", Field: "thing"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Pink", Url: "/pink/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Problem", Url: "/problem/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Programmer", Url: "/programmer/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Pulp", Url: "/pulp/:language/:from", Fields: []field{{Name: "Language", Field: "language"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Retard", Url: "/retard/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "That's fucking ridiculous", Url: "/ridiculous/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "RTFM", Url: "/rtfm/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "sake", Url: "/sake/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Shakespeare", Url: "/shakespeare/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck This Shit", Url: "/shit/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Shut Up", Url: "/shutup/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Single", Url: "/single/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Thanks", Url: "/thanks/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck That", Url: "/that/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "You Think", Url: "/think/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Thinking", Url: "/thinking/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck This", Url: "/this/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "This Guy", Url: "/thumbs/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Too", Url: "/too/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Tucker", Url: "/tucker/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "version", Url: "/version", Fields: []field{}})
	operations = append(operations, operation{Name: "What", Url: "/what/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Xmas", Url: "/xmas/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Yoda", Url: "/yoda/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Fuck You", Url: "/you/:name/:from", Fields: []field{{Name: "Name", Field: "name"}, {Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Zayn", Url: "/zayn/:from", Fields: []field{{Name: "From", Field: "from"}}})
	operations = append(operations, operation{Name: "Zero", Url: "/zero/:from", Fields: []field{{Name: "From", Field: "from"}}})

	json.NewEncoder(w).Encode(operations)
}
