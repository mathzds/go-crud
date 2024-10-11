package models

type Release struct {
	Episode struct {
		N_Episodio string `json:"n_episodio"` 
		Anime struct {
			SlugSerie string `json:"slug_serie"`
		} `json:"anime"`
	} `json:"episode"`
}

type Data struct {
	Props struct {
		PageProps struct {
			Data struct {
				DataReleases []Release `json:"data_releases"` 
			} `json:"data"`
		} `json:"pageProps"`
	} `json:"props"`
}
