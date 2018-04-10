package game

// Item defines an interactable item inside the game world
type Item struct {
	Name        string
	NameArticle string
	Type        ItemType
	Durability  int
	ExamineText string
}

// ItemType defines an enum which defines the shared properties of an item
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
