package imdb

import "regexp"

type CreditCategory int

const (
	ctNil CreditCategory = iota
	ctDirectors
	ctWriters
	ctCast
	ctProducers
	ctComposers
	ctCinematographers
	ctEditors
	ctCastingDirectors
	ctProductionDesigner
	ctArtDirector
	ctSetDirector
	ctCostumeDesigner
	ctProductionManager
	ctAssistantDirector
	ctArtDepartment
	ctSoundDepartment
	ctSpecialEffects
	ctVisualEffects
	ctStunts
	ctCameraDepartment
	ctAnimators
	ctEditorialDepartment
	ctMusicDepartment
	ctScriptDepartment
)

var (
	CreditsRe = map[CreditCategory]*regexp.Regexp{
		ctDirectors:           regexp.MustCompile(`^Directors*$`),
		ctWriters:             regexp.MustCompile(`^Writers*$`),
		ctCast:                regexp.MustCompile(`^Cast$`),
		ctProducers:           regexp.MustCompile(`^Producers*$`),
		ctComposers:           regexp.MustCompile(`^Composers*`),
		ctCinematographers:    regexp.MustCompile(`^Cinematographers*`),
		ctEditors:             regexp.MustCompile(`^Editors*`),
		ctCastingDirectors:    regexp.MustCompile(`^Casting.*`),
		ctProductionDesigner:  regexp.MustCompile(`^Production\sDesigners*`),
		ctArtDirector:         regexp.MustCompile(`^Art\sDirector.*`),
		ctSetDirector:         regexp.MustCompile(`^Set\sDecorator.*`),
		ctCostumeDesigner:     regexp.MustCompile(`^Costume\sDesigner.*`),
		ctProductionManager:   regexp.MustCompile(`^Production\sManager.*`),
		ctAssistantDirector:   regexp.MustCompile(`^Assistant\sDirector.*`),
		ctArtDepartment:       regexp.MustCompile(`^Art\sDepartment.*`),
		ctSoundDepartment:     regexp.MustCompile(`^Sound\sDepartment.*`),
		ctSpecialEffects:      regexp.MustCompile(`^Special\sEffects.*`),
		ctVisualEffects:       regexp.MustCompile(`^Visual\sEffects.*`),
		ctStunts:              regexp.MustCompile(`^Stunts.*`),
		ctCameraDepartment:    regexp.MustCompile(`^Camera\sDepartment.*`),
		ctAnimators:           regexp.MustCompile(`^Animation\sDepartment.*`),
		ctEditorialDepartment: regexp.MustCompile(`^Editorial\sDepartment.*`),
		ctMusicDepartment:     regexp.MustCompile(`^Music\sDepartment.*`),
		ctScriptDepartment:    regexp.MustCompile(`^Script\sDepartment.*`),
	}
)

func getCreditCategory(label string) CreditCategory {
	for k, v := range CreditsRe {
		if v.MatchString(label) {
			return k
		}
	}
	return ctNil
}
