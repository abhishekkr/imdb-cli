package imdb

import (
	"github.com/PuerkitoBio/goquery"

	"github.com/abhishekkr/imdb-cli/config"
)

type Credits struct {
	Directors           []Artist
	Writers             []Artist
	Cast                []Artist
	Producers           []Artist
	Composers           []Artist
	Cinematographers    []Artist
	Editors             []Artist
	CastingDirectors    []Artist
	ProductionDesigner  []Artist
	ArtDirector         []Artist
	SetDirector         []Artist
	CostumeDesigner     []Artist
	ProductionManager   []Artist
	AssistantDirector   []Artist
	ArtDepartment       []Artist
	SoundDepartment     []Artist
	SpecialEffects      []Artist
	VisualEffects       []Artist
	Stunts              []Artist
	CameraDepartment    []Artist
	Animators           []Artist
	EditorialDepartment []Artist
	MusicDepartment     []Artist
	ScriptDepartment    []Artist
}

func (credit *Credits) get(doc *goquery.Document) error {
	creditSections := doc.Find(config.CreditSectionSelector)
	creditSections.Each(func(i int, s *goquery.Selection) {
		sectionLabel := s.Find(config.CreditLabelSelector).Text()
		switch getCreditCategory(sectionLabel) {
		case ctDirectors:
			credit.Directors = queryArtists(s)
		case ctWriters:
			credit.Writers = queryArtists(s)
		case ctCast:
			credit.Cast = queryArtists(s)
		case ctProducers:
			credit.Producers = queryArtists(s)
		case ctComposers:
			credit.Composers = queryArtists(s)
		case ctCinematographers:
			credit.Cinematographers = queryArtists(s)
		case ctEditors:
			credit.Editors = queryArtists(s)
		case ctCastingDirectors:
			credit.CastingDirectors = queryArtists(s)
		case ctProductionDesigner:
			credit.ProductionDesigner = queryArtists(s)
		case ctArtDirector:
			credit.ArtDirector = queryArtists(s)
		case ctSetDirector:
			credit.SetDirector = queryArtists(s)
		case ctCostumeDesigner:
			credit.CostumeDesigner = queryArtists(s)
		case ctProductionManager:
			credit.ProductionManager = queryArtists(s)
		case ctAssistantDirector:
			credit.AssistantDirector = queryArtists(s)
		case ctArtDepartment:
			credit.ArtDepartment = queryArtists(s)
		case ctSoundDepartment:
			credit.SoundDepartment = queryArtists(s)
		case ctSpecialEffects:
			credit.SpecialEffects = queryArtists(s)
		case ctVisualEffects:
			credit.VisualEffects = queryArtists(s)
		case ctStunts:
			credit.Stunts = queryArtists(s)
		case ctCameraDepartment:
			credit.CameraDepartment = queryArtists(s)
		case ctAnimators:
			credit.Animators = queryArtists(s)
		case ctEditorialDepartment:
			credit.EditorialDepartment = queryArtists(s)
		case ctMusicDepartment:
			credit.MusicDepartment = queryArtists(s)
		case ctScriptDepartment:
			credit.ScriptDepartment = queryArtists(s)
		}
	})
	return nil
}
