package models

type CalendarAnime struct {
	Name      *string `json:"name"`
	Comments  *string `json:"comments"`
	Day       int     `json:"day"`
	WasFound  bool    `json:"wasFound"`
}

type Anime struct {
	IDSerie        int          `json:"id_serie"`
	Pgad          int           `json:"pgad"`
	Dub           int           `json:"dub"`
	Generos       string        `json:"generos"`
	Score         float64       `json:"score"`
	Titulo        string        `json:"titulo"`
	OriginalTitulo string       `json:"originaltitulo"`
	SlugSerie     string        `json:"slug_serie"`
	Ano           int           `json:"ano"`
	Diretor       string        `json:"diretor"`
	Elenco        string        `json:"elenco"`
	Duracao       string        `json:"duracao"`
	Origem        string        `json:"origem"`
	Censura       int           `json:"censura"`
	Sinopse       string        `json:"sinopse"`
	Aviso         string        `json:"aviso"`
	GenerateID    string        `json:"generate_id"`
	DataRegistro  string        `json:"data_registro"`
	Episodes      int           `json:"episodes"`
	FavAnimes     int           `json:"favAnimes"`
	CalendarAnime  CalendarAnime `json:"calendar_anime"`
}

type AnimeDataResponse struct {
	Props struct {
		PageProps struct {
			Data struct {
				Anime Anime `json:"data"` // This should correctly match the JSON structure.
			} `json:"data"`
		} `json:"pageProps"`
	} `json:"props"`
}
