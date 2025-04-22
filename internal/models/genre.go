package models

import (
	"database/sql/driver"
)

type GenreBook string

const (
    GenreFiction         GenreBook = "fiction"
    GenreNonFiction      GenreBook = "non-fiction"
    GenreMystery         GenreBook = "mystery"
    GenreFantasy         GenreBook = "fantasy"
    GenreScienceFiction  GenreBook = "science fiction"
    GenreBiography       GenreBook = "biography"
    GenreHistory         GenreBook = "history"
    GenreRomance         GenreBook = "romance"
    GenreThriller        GenreBook = "thriller"
    GenreSelfHelp        GenreBook = "self-help"
    GenreChildren        GenreBook = "children"
    GenreYoungAdult      GenreBook = "young adult"
    GenreHorror          GenreBook = "horror"
    GenrePoetry          GenreBook = "poetry"
    GenreCookbook        GenreBook = "cookbook"
    GenreGraphicNovel    GenreBook = "graphic novel"
    GenreTravel          GenreBook = "travel"
    GenreHealth          GenreBook = "health"
    GenreBusiness        GenreBook = "business"
    GenreReligion        GenreBook = "religion"
    GenrePhilosophy      GenreBook = "philosophy"
    GenreArt             GenreBook = "art"
    GenreMusic           GenreBook = "music"
    GenreSports          GenreBook = "sports"
    GenreTechnology      GenreBook = "technology"
    GenreEducation       GenreBook = "education"
    GenreParenting       GenreBook = "parenting"
    GenreHomeGarden      GenreBook = "home and garden"
    GenreCraftsHobbies   GenreBook = "crafts and hobbies"
    GenreComputers       GenreBook = "computers"
    GenreInternet        GenreBook = "internet"
    GenreScience         GenreBook = "science"
    GenreMathematics     GenreBook = "mathematics"
    GenreEngineering     GenreBook = "engineering"
    GenreLaw             GenreBook = "law"
    GenrePolitics        GenreBook = "politics"
    GenreSocialSciences  GenreBook = "social sciences"
)

func (g *GenreBook) Scan(value interface{}) error {
    *g = GenreBook(value.(string))
    return nil
}

func (g GenreBook) Value() (driver.Value, error) {
    return string(g), nil
}

func (g GenreBook) IsValid() bool {
    switch g {
    case GenreFiction, GenreNonFiction, GenreMystery, GenreFantasy,
        GenreScienceFiction, GenreBiography, GenreHistory, GenreRomance,
        GenreThriller, GenreSelfHelp, GenreChildren, GenreYoungAdult,
        GenreHorror, GenrePoetry, GenreCookbook, GenreGraphicNovel,
        GenreTravel, GenreHealth, GenreBusiness, GenreReligion,
        GenrePhilosophy, GenreArt, GenreMusic, GenreSports,
        GenreTechnology, GenreEducation, GenreParenting, GenreHomeGarden,
        GenreCraftsHobbies, GenreComputers, GenreInternet, GenreScience,
        GenreMathematics, GenreEngineering, GenreLaw, GenrePolitics,
        GenreSocialSciences:
        return true
    }
    return false
}