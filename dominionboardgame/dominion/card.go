package dominion

import (
	"fmt"
	"strings"
)

type Actioner interface {
	//Draw(p *Player)
	//AddBuy(p *Player)
	//AddAction(p *Player)
	InitCard()
	GetCardID() CardID
	DoAbility(p *Player)
	DoSpecialAbility(p *Player, g *GameMan)
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
	Curse
	Village
	Festival
	Smithy
	Market
	//Bandit
	Laboratory
	Artisan
	Cellar
	Chapel

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
	"Curse",
	// Original : Action Card
	"Village",
	"Festival",
	"Smithy",
	"Market",
	//"Bandit",
	"Laboratory",
	"Artisan",
	"Cellar",
	"Chapel",

	// Intrigue : Action Card
	"Upgrade",
}

func (r CardID) String() string {
	return CardIDString[r%MaxCardID]
}

func (r CardIDs) String() string {
	s := ""
	s = fmt.Sprintf(":%dn|", len(r))
	for i, v := range r {
		s += fmt.Sprintf("#%d:%s|", i, v)
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

type Cards []*Card

/*
func (r *Card) InitCard() {

}
*/

func (r *Card) GetCardID() CardID {
	return r.CardID
}

func (r *Card) GetAbilityCount(a AbilityType) (int, bool) {
	for _, v := range r.Ability {
		if v.abilityType == a {
			return v.count, true
		}
	}

	return 0, false
}

func (r *Card) DoAbility(p *Player) {
	//sline + "\n" + s + "\n"

	sbuy := r.AddBuy(p)
	saction := r.AddAction(p)
	scard := r.AddCard(p)
	scoin := r.AddCoin(p)

	s := fmt.Sprintf("<PlayCard> Player:%s(ID:%d) %s", p.name, p.ID, r.CardID)
	sline := strings.Repeat("~", len(s))

	log := sline
	log += "\n"
	log += s
	log += "\n"
	log += sline
	log += "\n"
	log += sbuy
	log += saction
	log += scard
	log += scoin
	log += sline

	fmt.Println(log)
}

func (r *Card) DoSpecialAbility(p *Player, g *GameMan) {
}

func (r *Card) AddBuy(p *Player) string {
	cnt, _ := r.GetAbilityCount(AbilityAddBuy)
	s := ""
	if cnt > 0 {
		s = fmt.Sprintf("\tAddBuy:buys=%d+%d\n", p.buys, cnt)
	}
	p.buys += cnt

	return s
}

func (r *Card) AddAction(p *Player) string {
	cnt, _ := r.GetAbilityCount(AbilityAddAction)

	s := ""
	if cnt > 0 {
		s = fmt.Sprintf("\tAddAction:actions=%d+%d\n", p.actions, cnt)
	}
	p.actions += cnt

	return s
}

func (r *Card) AddCard(p *Player) string {
	cnt, _ := r.GetAbilityCount(AbilityAddCard)

	cardIDs, _ := p.DrawCard(cnt)

	s := ""
	if cnt > 0 {
		s = fmt.Sprintf("\tAddCard%s", cardIDs)
	}
	return s
}
func (r *Card) AddCoin(p *Player) string {
	cnt, _ := r.GetAbilityCount(AbilityAddCoin)

	s := ""
	if cnt > 0 {
		s = fmt.Sprintf("\tAddCoin:coins=%d+%d\n", p.coins, cnt)
	}

	p.coins += cnt

	return s
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
