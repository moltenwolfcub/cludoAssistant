package cluedo

import (
	"slices"
	"testing"
)

func lookupCard(t *testing.T, category CardCategory, cardName string) (found *Card) {
	for _, c := range category.Cards {
		if c.name == cardName {
			found = c
			return
		}
	}
	t.Errorf("Couldn't find card %s", cardName)
	return nil
}

func GenSampleGame() (game Game, a, b, c *Player) {
	a = NewPlayer("alice", 0)
	b = NewPlayer("bob", 0)
	c = NewPlayer("charlie", 0)

	game = NewDefaultGame(a, b, c)
	return
}

func TestTurnWhoAnswerUpdatesFound(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhoAnswer)

	game.DoTurn(question)

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if !whiteCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a who card but didn't update the who card")
	}

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if pistolCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a who card but the what card was set to found")
	}

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if bedroomCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a who card but the where card was set to found")
	}
}

func TestTurnWhatAnswerUpdatesFound(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhatAnswer)

	game.DoTurn(question)

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if !pistolCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a what card but didn't update the what card")
	}

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if whiteCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a what card but the who card was set to found")
	}

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if bedroomCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a what card but the where card was set to found")
	}
}

func TestTurnWhereAnswerUpdatesFound(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhereAnswer)

	game.DoTurn(question)

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if !bedroomCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a where card but didn't update the where card")
	}

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if whiteCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a where card but the who card was set to found")
	}

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if pistolCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a where card but the what card was set to found")
	}
}

func TestTurnWhoAnswerPosessor(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhoAnswer)

	game.DoTurn(question)

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if whiteCard.possessor != alice {
		t.Error("Game.DoTurn() Was shown a who card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestTurnWhatAnswerPosessor(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhatAnswer)

	game.DoTurn(question)

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if pistolCard.possessor != alice {
		t.Error("Game.DoTurn() Was shown a what card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestTurnWhereAnswerPosessor(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhereAnswer)

	game.DoTurn(question)

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if bedroomCard.possessor != alice {
		t.Error("Game.DoTurn() Was shown a where card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestStartingHandOneCard(t *testing.T) {
	game, _, _, _ := GenSampleGame()
	game.AddStartingHand(
		[]*Card{NewCard("lead pipe")},
	)

	pipeCard := lookupCard(t, game.whatCategory, "lead pipe")
	if !pipeCard.IsFound() {
		t.Error("Game.AddStartingHand() Started with 1 card but it wasn't marked as found.")
	}
	if pipeCard.possessor != game.Me {
		t.Error("Game.AddStartingHand() Started with 1 card but its owner wasn't THIS.")
	}
}

func TestStartingHandMultipleCards(t *testing.T) {
	game, _, _, _ := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("wrench"),
		NewCard("green"),
		NewCard("study"),
		NewCard("bathroom"),
	})

	if !lookupCard(t, game.whatCategory, "wrench").IsFound() {
		t.Error("Game.AddStartingHand() Wrench card wasn't marked as found when it was in the starting hand")
	}
	if !lookupCard(t, game.whoCategory, "green").IsFound() {
		t.Error("Game.AddStartingHand() Green card wasn't marked as found when it was in the starting hand")
	}
	if !lookupCard(t, game.whereCategory, "study").IsFound() {
		t.Error("Game.AddStartingHand() Study card wasn't marked as found when it was in the starting hand")
	}
	if !lookupCard(t, game.whereCategory, "bathroom").IsFound() {
		t.Error("Game.AddStartingHand() Bathroom card wasn't marked as found when it was in the starting hand")
	}
}

func TestTurnNonPossessor(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(NoAnswer)

	game.DoTurn(question)

	whoCard := lookupCard(t, game.whoCategory, "white")
	if !slices.Contains(whoCard.nonPossessors, alice) {
		t.Error("Game.DoTurn() Alice couldn't answer the question but she wasn't marked as not having the person")
	}
	whatCard := lookupCard(t, game.whatCategory, "pistol")
	if !slices.Contains(whatCard.nonPossessors, alice) {
		t.Error("Game.DoTurn() Alice couldn't answer the question but she wasn't marked as not having the weapon")
	}
	whereCard := lookupCard(t, game.whereCategory, "bedroom")
	if !slices.Contains(whereCard.nonPossessors, alice) {
		t.Error("Game.DoTurn() Alice couldn't answer the question but she wasn't marked as not having the location")
	}
}

func TestUnkownAnswerWith2SimpleSelfKnown(t *testing.T) {
	game, alice, bob, _ := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("green"),
		NewCard("dagger"),
	})

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		alice,
	)
	question.SetAnswer(UnknownAnswer)

	game.DoTurn(question)

	whereCard := lookupCard(t, game.whereCategory, "bedroom")
	if whereCard.possessor != alice {
		t.Error("Game.analyseUnknownAnswer() I had 2 cards and alice showed a card when asked about them but the 3rd wasn't marked as hers.")
	}
}

func TestUnkownAnswerWith2SimpleOthersKnown(t *testing.T) {
	game, alice, bob, charlie := GenSampleGame()

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhoAnswer)
	game.DoTurn(question)

	question = NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		game.Me,
		bob,
	)
	question.SetAnswer(WhatAnswer)
	game.DoTurn(question)

	question = NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	whereCard := lookupCard(t, game.whereCategory, "bedroom")
	if whereCard.possessor != charlie {
		t.Error("Game.analyseUnknownAnswer() 2 cards were in known locations and charlie showed a card when asked about them but the 3rd wasn't marked as his.")
	}
}

func TestUnkownAnswerWith2KnownsAssumptions(t *testing.T) {
	game, _, bob, charlie := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("green"),
	})

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		game.Me,
		charlie,
	)
	question.SetAnswer(WhatAnswer)
	game.DoTurn(question)

	question = NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	whereCard := lookupCard(t, game.whereCategory, "bedroom")
	if whereCard.possessor == charlie {
		t.Error("Game.analyseUnknownAnswer() 2 cards were in known locations but one was charlie's and charlie showed a card when asked about them and the 3rd was incorrectly assumed to have been shown.")
	}
}

func TestUnkownAnswerWith1Known(t *testing.T) {
	game, _, bob, charlie := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("green"),
	})

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	daggerCard := lookupCard(t, game.whatCategory, "dagger")
	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")

	daggerLink := Link{
		player: charlie,
		other:  bedroomCard,
	}
	bedroomLink := Link{
		player: charlie,
		other:  daggerCard,
	}

	if !slices.Contains(daggerCard.links, daggerLink) {
		t.Error("Game.analyseUnknownAnswer() 1 card was in a known location but Dagger didn't have Bedroom as a link")
	}
	if !slices.Contains(bedroomCard.links, bedroomLink) {
		t.Error("Game.analyseUnknownAnswer() 1 card was in a known location but Bedroom didn't have Dagger as a link")
	}
}

func TestUnkownAnswerWith0Knowns(t *testing.T) {
	game, _, bob, charlie := GenSampleGame()

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	greenCard := lookupCard(t, game.whoCategory, "green")
	daggerCard := lookupCard(t, game.whatCategory, "dagger")
	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")

	greenLink := TriLink{
		this:   greenCard,
		player: charlie,
		other1: bedroomCard,
		other2: daggerCard,
	}
	daggerLink := TriLink{
		this:   daggerCard,
		player: charlie,
		other1: bedroomCard,
		other2: greenCard,
	}
	bedroomLink := TriLink{
		this:   bedroomCard,
		player: charlie,
		other1: daggerCard,
		other2: greenCard,
	}

	if !slices.ContainsFunc(greenCard.trilinks, func(t TriLink) bool { return t.Equals(greenLink) }) {
		t.Error("Game.analyseUnknownAnswer() no card was in a known location but Green didn't have Dagger and Bedroom as a trilink")
	}
	if !slices.ContainsFunc(daggerCard.trilinks, func(t TriLink) bool { return t.Equals(daggerLink) }) {
		t.Error("Game.analyseUnknownAnswer() no card was in a known location but Dagger didn't have Bedroom and Green as a trilink")
	}
	if !slices.ContainsFunc(bedroomCard.trilinks, func(t TriLink) bool { return t.Equals(bedroomLink) }) {
		t.Error("Game.analyseUnknownAnswer() no card was in a known location but Bedroom didn't have Green and Dagger as a trilink")
	}
}

func TestLinkResolutionWithFind(t *testing.T) {
	game, alice, bob, charlie := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("green"),
	})

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	question = NewQuestion(
		NewCard("peacock"),
		NewCard("dagger"),
		NewCard("living room"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhatAnswer)
	game.DoTurn(question)

	daggerCard := lookupCard(t, game.whatCategory, "dagger")
	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")

	if !bedroomCard.IsFound() {
		t.Error("Charlie had either the bedroom or the dagger and we know alice has the dagger but the bedroom wasn't marked as found")
	}
	if bedroomCard.possessor != charlie {
		t.Error("Charlie had either the bedroom or the dagger and we know alice has the dagger but charlie wasn't the possessor of the bedroom")
	}

	daggerLink := Link{
		player: charlie,
		other:  bedroomCard,
	}
	bedroomLink := Link{
		player: charlie,
		other:  daggerCard,
	}
	if slices.Contains(daggerCard.links, daggerLink) {
		t.Error("The dagger's link has served it's purpose but it wasn't removed")
	}
	if slices.Contains(bedroomCard.links, bedroomLink) {
		t.Error("The bedroom's link has served it's purpose but it wasn't removed")
	}
}

func TestLinkResolutionWithoutFind(t *testing.T) {
	game, _, bob, charlie := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("green"),
	})

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	question = NewQuestion(
		NewCard("peacock"),
		NewCard("dagger"),
		NewCard("living room"),
		game.Me,
		charlie,
	)
	question.SetAnswer(WhatAnswer)
	game.DoTurn(question)

	daggerCard := lookupCard(t, game.whatCategory, "dagger")
	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")

	if bedroomCard.IsFound() {
		t.Error("Charlie had either the bedroom or the dagger and we now know charlie has the dagger but the bedroom was incorrectly marked as found")
	}

	daggerLink := Link{
		player: charlie,
		other:  bedroomCard,
	}
	bedroomLink := Link{
		player: charlie,
		other:  daggerCard,
	}
	if slices.Contains(daggerCard.links, daggerLink) {
		t.Error("The dagger's link has served it's purpose but it wasn't removed")
	}
	if slices.Contains(bedroomCard.links, bedroomLink) {
		t.Error("The bedroom's link has served it's purpose but it wasn't removed")
	}
}

func TestTriLinkResolutionWithFind(t *testing.T) {
	game, alice, bob, charlie := GenSampleGame()

	//create trilink
	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	//shrink to link
	question = NewQuestion(
		NewCard("peacock"),
		NewCard("dagger"),
		NewCard("living room"),
		game.Me,
		alice,
	)
	question.SetAnswer(WhatAnswer)
	game.DoTurn(question)

	//check
	greenCard := lookupCard(t, game.whoCategory, "green")
	daggerCard := lookupCard(t, game.whatCategory, "dagger")
	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")

	greenLink := Link{
		player: charlie,
		other:  bedroomCard,
	}
	bedroomLink := Link{
		player: charlie,
		other:  greenCard,
	}

	if !slices.Contains(greenCard.links, greenLink) {
		t.Error("Game.analyseUnknownAnswer() TriLink should have been resolved to a normal link but one wasn't created")
	}
	if !slices.Contains(bedroomCard.links, bedroomLink) {
		t.Error("Game.analyseUnknownAnswer() TriLink should have been resolved to a normal link but one wasn't created")
	}

	greenTriLink := TriLink{
		this:   greenCard,
		player: charlie,
		other1: daggerCard,
		other2: bedroomCard,
	}
	daggerTriLink := TriLink{
		this:   daggerCard,
		player: charlie,
		other1: greenCard,
		other2: bedroomCard,
	}
	bedroomTriLink := TriLink{
		this:   bedroomCard,
		player: charlie,
		other1: daggerCard,
		other2: greenCard,
	}

	if slices.ContainsFunc(greenCard.trilinks, func(t TriLink) bool { return greenTriLink.Equals(t) }) {
		t.Error("The bedroom's trilink has served it's purpose but it wasn't removed")
	}
	if slices.ContainsFunc(daggerCard.trilinks, func(t TriLink) bool { return daggerTriLink.Equals(t) }) {
		t.Error("The dagger's trilink has served it's purpose but it wasn't removed")
	}
	if slices.ContainsFunc(bedroomCard.trilinks, func(t TriLink) bool { return bedroomTriLink.Equals(t) }) {
		t.Error("The bedroom's trilink has served it's purpose but it wasn't removed")
	}
}

func TestTriLinkResolutionWithoutFind(t *testing.T) {
	game, _, bob, charlie := GenSampleGame()

	//create trilink
	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		bob,
		charlie,
	)
	question.SetAnswer(UnknownAnswer)
	game.DoTurn(question)

	//shrink to link
	question = NewQuestion(
		NewCard("peacock"),
		NewCard("dagger"),
		NewCard("living room"),
		game.Me,
		charlie,
	)
	question.SetAnswer(WhatAnswer)
	game.DoTurn(question)

	//check
	greenCard := lookupCard(t, game.whoCategory, "green")
	daggerCard := lookupCard(t, game.whatCategory, "dagger")
	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")

	greenLink := Link{
		player: charlie,
		other:  bedroomCard,
	}
	bedroomLink := Link{
		player: charlie,
		other:  greenCard,
	}

	if slices.Contains(greenCard.links, greenLink) {
		t.Error("Game.analyseUnknownAnswer() TriLink was resolved to a normal link but shouldn't have been")
	}
	if slices.Contains(bedroomCard.links, bedroomLink) {
		t.Error("Game.analyseUnknownAnswer() TriLink was resolved to a normal link but shouldn't have been")
	}

	greenTriLink := TriLink{
		this:   greenCard,
		player: charlie,
		other1: daggerCard,
		other2: bedroomCard,
	}
	daggerTriLink := TriLink{
		this:   daggerCard,
		player: charlie,
		other1: greenCard,
		other2: bedroomCard,
	}
	bedroomTriLink := TriLink{
		this:   bedroomCard,
		player: charlie,
		other1: daggerCard,
		other2: greenCard,
	}

	if slices.ContainsFunc(greenCard.trilinks, func(t TriLink) bool { return greenTriLink.Equals(t) }) {
		t.Error("The bedroom's trilink has served it's purpose but it wasn't removed")
	}
	if slices.ContainsFunc(daggerCard.trilinks, func(t TriLink) bool { return daggerTriLink.Equals(t) }) {
		t.Error("The dagger's trilink has served it's purpose but it wasn't removed")
	}
	if slices.ContainsFunc(bedroomCard.trilinks, func(t TriLink) bool { return bedroomTriLink.Equals(t) }) {
		t.Error("The bedroom's trilink has served it's purpose but it wasn't removed")
	}
}
