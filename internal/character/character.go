package character

const (
	MELEE = iota
)

type Character struct {
	IsMine   bool
	Type     int
	Mobility int
}

func GetCharacter(t int, mobility int, isMine bool) *Character {
	ch := &Character{
		Type:     t,
		Mobility: mobility,
		IsMine:   isMine,
	}
	return ch
}
