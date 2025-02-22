package actors

import "github.com/NChitty/archaeologist/pkg/characters"

type Actor interface {
	Do(character *characters.CharacterWrapper) error
}
