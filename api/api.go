package api

type operation struct {
	Name   string  `json:"name"`
	Url    string  `json:"url"`
	Fields []field `json:"fields"`
}

type field struct {
	Name  string `json:"name"`
	Field string `json:"field"`
}

type apiResponse struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}
