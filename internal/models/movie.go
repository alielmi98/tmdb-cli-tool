package models

type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}
type Dates struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}
type MovieRes struct {
	Dates         Dates   `json:"dates"`
	Page          int     `json:"page"`
	Results       []Movie `json:"results"`
	Total_Pages   int     `json:"total_pages"`
	Total_Results int     `json:"total_results"`
}

type MovieResError struct {
	Success        bool   `json:"success"`
	Status_Code    int    `json:"status_code"`
	Status_Message string `json:"status_message"`
}
