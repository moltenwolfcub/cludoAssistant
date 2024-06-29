package cluedo

import "slices"

type Card struct {
	name string

	isMurderItem bool

	found         bool
	possessor     *Player
	nonPossessors []*Player

	links    []Link
	trilinks []TriLink
}

func NewCard(name string) *Card {
	return &Card{
		name: name,
	}
}

func (c *Card) SetFound(possessor *Player, destroyLinks bool) {
	c.found = true
	c.possessor = possessor
	c.isMurderItem = false

	if destroyLinks {
		for i, l := range c.links {
			if l.player != possessor {
				//set the other half of the link to found
				l.other.SetFound(l.player, false)
			}
			//destroy the link as it's now redundant
			for j := range l.other.links {
				if l.other.links[j].other == c && l.other.links[j].player == l.player {
					l.other.links = slices.Delete(l.other.links, j, j+1)
					break
				}
			}
			c.links = slices.Delete(c.links, i, i+1)
		}
	}

	// resolve trilinks
	for i, t := range c.trilinks {
		if t.player != possessor {
			//shrink the trilink to a normal link
			t.other1.AddLink(t.player, t.other2)
			t.other2.AddLink(t.player, t.other1)
		}
		//destroy the link as it's now redundant
		for j := range t.other1.trilinks {
			if t.other1.trilinks[j].Equals(t) {
				t.other1.trilinks = slices.Delete(t.other1.trilinks, j, j+1)
				break
			}
		}
		for j := range t.other2.trilinks {
			if t.other2.trilinks[j].Equals(t) {
				t.other2.trilinks = slices.Delete(t.other2.trilinks, j, j+1)
				break
			}
		}
		c.trilinks = slices.Delete(c.trilinks, i, i+1)
	}
}

func (c Card) IsFound() bool {
	return c.found
}

func (c *Card) AddNonPossessor(player *Player) {
	if slices.Contains(c.nonPossessors, player) {
		return
	}

	c.nonPossessors = append(c.nonPossessors, player)
}

func (c *Card) AddLink(player *Player, other *Card) {
	newLink := Link{
		player: player,
		other:  other,
	}

	if !slices.Contains(c.links, newLink) {
		c.links = append(c.links, newLink)
	}
}

func (c *Card) AddTriLink(player *Player, one *Card, two *Card) {
	newLink := TriLink{
		this: c,

		player: player,
		other1: one,
		other2: two,
	}

	if !slices.Contains(c.trilinks, newLink) {
		c.trilinks = append(c.trilinks, newLink)
	}
}

type Link struct {
	player *Player
	other  *Card
}

type TriLink struct {
	this *Card

	player *Player
	other1 *Card
	other2 *Card
}

func (t TriLink) Equals(other TriLink) bool {
	if t.player != other.player {
		return false
	}
	thisCards := []*Card{
		t.this,
		t.other1,
		t.other2,
	}
	otherCards := []*Card{
		other.this,
		other.other1,
		other.other2,
	}

	for _, c := range thisCards {
		if !slices.Contains(otherCards, c) {
			return false
		}
	}

	return true
}

type CardCategory struct {
	Cards []*Card
}

func NewCardCategory(cards ...*Card) CardCategory {
	return CardCategory{
		Cards: cards,
	}
}

func (c CardCategory) UpdateMurderKnowledge() {
	foundCards := 0

	var potentialMurderPart *Card
	for _, card := range c.Cards {
		if card.isMurderItem {
			return
		}
		if card.IsFound() {
			foundCards++
		} else {
			potentialMurderPart = card
		}
	}
	if len(c.Cards)-foundCards == 1 {
		potentialMurderPart.isMurderItem = true
	}
}

func (c CardCategory) HasKnownSolution() bool {
	return c.GetKnownSolution() != nil
}
func (c CardCategory) GetKnownSolution() *Card {
	c.UpdateMurderKnowledge()
	for _, card := range c.Cards {
		if card.isMurderItem {
			return card
		}
	}
	return nil
}

func (c *CardCategory) FoundCard(foundCard *Card, possessor *Player) (success bool) {
	for _, card := range c.Cards {
		if foundCard.name == card.name {
			card.SetFound(possessor, true)
			return true
		}
	}
	return false
}
