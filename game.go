package main

// Game describes the game state JSON
type Game struct {
	// The small blind in the current round. The big blind is twice
	// the small blind
	SmallBlind int `json:"small_blind"`

	// The amount of the largest current bet from any one player
	CurrentBuyIn int `json:"current_buy_in"`

	// The size of the pot (sum of the player bets)
	Pot int `json:"pot"`

	// Minimum raise amount. To raise you have to return at least:
	//     current_buy_in - players[in_action][bet] + minimum_raise
	MinimumRaise int `json:"minimum_raise"`

	// The index of the player on the dealer button in this round
	// The first player is (dealer+1)%(players.length)
	Dealer int `json:"dealer"`

	// Number of orbits completed. (The number of times the dealer button
	// returned to the same player.)
	Orbits int `json:"orbits"`

	// The index of your player, in the players array
	InAction int `json:"in_action"`

	// An array of the players. The order stays the same during the
	// entire tournament
	Players []Player `json:"players"`

	// Finally the array of community cards.
	CommunityCards []Card `json:"community_cards"`
}
