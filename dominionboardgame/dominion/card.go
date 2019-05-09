package dominion

import (
	"fmt"
)

type Actioner interface {
	Draw(p *Player)
	AddBuy(p *Player)
	AddAction(p *Player)
	String() string
	//DoSpecailACtion()
}

// CardType is	Action, Treasure, Victory
type CardType int

const (
	CardTypeAction CardType = 0 + iota
	CardTypeTreasure
	CardTypeVictory
	CardTypeCurse
)

const (
	InvaildCardID = -1
)

const (
	CardWidth = 20
)

var CardTypeString = [...]string{
	"Action",
	"Treasure",
	"Victory",
	"Curse",
}

func (r CardType) String() string {
	return CardTypeString[r%3]
}

type CardID int

type CardIDs []CardID

const (
	Copper CardID = 0 + iota
	Silver
	Gold
	Estate
	Duchy
	Province
	Village
	Festival
	Smithy
	Market
	Bandit
	Upgrade
	MaxCardID
)

var CardIDString = [...]string{
	// Original : Treasure Card
	"Copper",
	"Silver",
	"Gold",
	// Original : Victory Card
	"Estate",
	"Duchy",
	"Province",
	// Original : Action Card
	"Village",
	"Festival",
	"Smithy",
	"Market",
	"Bandit",
	// Intrigue : Action Card
	"Upgrade",
}

func (r CardID) String() string {
	return CardIDString[r%MaxCardID]
}

func (r CardIDs) String() string {
	s := ""
	s = fmt.Sprintf("#%d\n\t|", len(r))
	for _, v := range r {
		s += fmt.Sprintf("%s|", v)
		//s += fmt.Sprintf("%s(%d)|", v, v)
	}
	s += "\n"

	return s
}

type Card struct {
	name     string
	CardID   CardID
	cardType []CardType
	cost     int
	Ability  []Ability
}

func (r Card) Draw(p *Player) {
	fmt.Println("Card.name", r)
}

func (r Card) AddBuy(p *Player) {
	p.buys++
}

func (r Card) AddAction(p *Player) {
	p.actions++
}

func (r Card) String() string {
	return fmt.Sprintf("%s%s(ID:%d)\n\tcost(%d)\n\tType%s\n\tAbility%s\n", r.name, r.CardID, r.CardID, r.cost, r.cardType, r.Ability)
}

func (r Card) TermString() string {

	ct := fmt.Sprintf("%s", r.cardType)
	//ct = strings.TrimLeft(ct, "[")
	//ct = strings.TrimRight(ct, "]")

	return ConvertTermString(CardWidth, r.name) + "\n" + ConvertTermString(CardWidth, ct) + "\n"

}

func init() {
	fmt.Println("import dominon/card")
}

func NewCardIDGenerator() func() CardID {
	var next int
	return func() CardID {
		next++
		return CardID(next)
	}
}

type PlayCardAbilityer interface {
	Play() error
}

type AbilityType int

const (
	AbilityAddAction AbilityType = 0 + iota
	AbilityAddCard
	AbilityAddBuy
	AbilityAddCoin
	AbilityAddVictory
	AbilitySpecial
	MaxAbility
)

var AbilityTypeString = [...]string{
	"Action",
	"Card",
	"Buy",
	"Coin",
	"Victory",
	"Special",
}

func (r AbilityType) String() string {
	return AbilityTypeString[r%MaxAbility]
}

type Ability struct {
	abilityType AbilityType
	count       int
}

func (r Ability) String() string {
	return fmt.Sprintf("\n\t\t+%d %s", r.count, r.abilityType)
}

func (r Card) Play(palyer *Player) error {
	fmt.Println(fmt.Sprintf("Play:%s", r))

	for _, v := range r.Ability {
		fmt.Println("\tAbility->", v)
		/*
			switch v.abilityType
			{
			case :
			case :
			}
		*/
	}

	return nil
}

func CallPlayCard(p PlayCardAbilityer) {
	p.Play()
}