package store

func (ds DeckScope) AvailableDecks() DeckScope {
	return ds.FullGame().Eq(true).Private().Eq(false)
}
