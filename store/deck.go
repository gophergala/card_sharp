package store

func (ds DeckScope) AvailableDecks() DeckScope {
	return ds.Full().Eq(true).Private().Eq(false)
}
