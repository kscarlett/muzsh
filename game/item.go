package game

type Item struct {
	Name        string
	NameArticle string
	Type        ItemType
	Durability  int
}

type ItemType int32

const (
	Useless ItemType = iota
	MeleeWeapon
	RangedWeapon
	Food
	Drink
	Medical
	Poison
)
