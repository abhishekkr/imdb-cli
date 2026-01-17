package imdb

import (
	"github.com/PuerkitoBio/goquery"

	"github.com/abhishekkr/imdb-cli/config"
)

type Credits struct {
	Directors                []Artist
	Writers                  []Artist
	Cast                     []Artist
	Producers                []Artist
	Composers                []Artist
	Cinematographers         []Artist
	Editors                  []Artist
	CastingDirectors         []Artist
	ProductionDesigner       []Artist
	ArtDirector              []Artist
	SetDirector              []Artist
	CostumeDesigner          []Artist
	MakeUpDepartment         []Artist
	ProductionManager        []Artist
	AssistantDirector        []Artist
	ArtDepartment            []Artist
	SoundDepartment          []Artist
	SpecialEffects           []Artist
	VisualEffects            []Artist
	Stunts                   []Artist
	CameraDepartment         []Artist
	Animators                []Artist
	CastingDepartment        []Artist
	CostumeDepartment        []Artist
	EditorialDepartment      []Artist
	LocationManagement       []Artist
	MusicDepartment          []Artist
	ScriptDepartment         []Artist
	TransportationDepartment []Artist
	Miscellaneous            []Artist
	Thanks                   []Artist
}

func (credit *Credits) get(doc *goquery.Document) {
	creditSummaryLabels := doc.Find(config.CreditSummaryLabelsSelector)
	creditSummaryValues := doc.Find(config.CreditSummaryValuesSelector)
	value := creditSummaryValues.First()
	creditSummaryLabels.Each(func(i int, s *goquery.Selection) {
		switch s.AttrOr("id", "") {
		case "director":
			credit.Directors = queryArtists(value)
		case "writer":
			credit.Writers = queryArtists(value)
		case "cast":
			credit.Cast = queryArtists(value)
		case "producer":
			credit.Producers = queryArtists(value)
		case "composer":
			credit.Composers = queryArtists(value)
		case "cinematographer":
			credit.Cinematographers = queryArtists(value)
		case "editor":
			credit.Editors = queryArtists(value)
		case "casting_director":
			credit.CastingDirectors = queryArtists(value)
		case "production_designer":
			credit.ProductionDesigner = queryArtists(value)
		case "art_director":
			credit.ArtDirector = queryArtists(value)
		case "set_decorator":
			credit.SetDirector = queryArtists(value)
		case "costume_designer":
			credit.CostumeDesigner = queryArtists(value)
		case "make_up_department":
			credit.MakeUpDepartment = queryArtists(value)
		case "production_manager":
			credit.ProductionManager = queryArtists(value)
		case "assistant_director":
			credit.AssistantDirector = queryArtists(value)
		case "art_department":
			credit.ArtDepartment = queryArtists(value)
		case "sound_department":
			credit.SoundDepartment = queryArtists(value)
		case "special_effects":
			credit.SpecialEffects = queryArtists(value)
		case "visual_effects":
			credit.VisualEffects = queryArtists(value)
		case "stunts":
			credit.Stunts = queryArtists(value)
		case "camera_department":
			credit.CameraDepartment = queryArtists(value)
		case "animation_department":
			credit.Animators = queryArtists(value)
		case "casting_department":
			credit.CastingDepartment = queryArtists(value)
		case "costume_department":
			credit.CostumeDepartment = queryArtists(value)
		case "editorial_department":
			credit.EditorialDepartment = queryArtists(value)
		case "location_management":
			credit.LocationManagement = queryArtists(value)
		case "music_department":
			credit.MusicDepartment = queryArtists(value)
		case "script_department":
			credit.ScriptDepartment = queryArtists(value)
		case "transportation_department":
			credit.TransportationDepartment = queryArtists(value)
		case "miscellaneous":
			credit.Miscellaneous = queryArtists(value)
		case "thanks":
			credit.Thanks = queryArtists(value)
		}
		value = value.Next()
	})
}
