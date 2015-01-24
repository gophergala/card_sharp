/*
  This code was generated by the Doctor ORM Generator and isn't meant to be edited.
	If at all possible, please regenerate this file from your gp files instead of
	attempting to edit it to add changes.
*/

package store

import "github.com/acsellers/dr/schema"

var Schema = schema.Schema{
	Tables: map[string]*schema.Table{

		"Account": &schema.Table{
			Name: "Account",
			Columns: []schema.Column{

				DefaultInt("ID"),

				DefaultString("Email"),

				DefaultBool("Management"),

				schema.Column{
					Name:   "CryptPassword",
					Type:   "blob",
					Length: 0,
				},
			},
			Index: []schema.Index{
				schema.Index{
					Columns: []string{"ID"},
				},

				schema.Index{
					Columns: []string{
						"Email",
					},
				},
			},
		},

		"Deck": &schema.Table{
			Name: "Deck",
			Columns: []schema.Column{

				DefaultInt("ID"),

				DefaultString("Name"),

				DefaultBool("Private"),

				DefaultBool("Full"),

				DefaultString("GameType"),

				DefaultInt("AccountID"),
			},
			Index: []schema.Index{
				schema.Index{
					Columns: []string{"ID"},
				},

				schema.Index{
					Columns: []string{
						"AccountID",
					},
				},
			},
		},

		"Card": &schema.Table{
			Name: "Card",
			Columns: []schema.Column{

				DefaultInt("ID"),

				DefaultString("Name"),

				DefaultString("Type"),

				schema.Column{
					Name:   "Data",
					Type:   "text",
					Length: 0,
				},

				DefaultInt("DeckID"),
			},
			Index: []schema.Index{
				schema.Index{
					Columns: []string{"ID"},
				},

				schema.Index{
					Columns: []string{
						"DeckID",
					},
				},
			},
		},
	},
}

func init() {

	Schema.Tables["Account"].HasMany = []schema.ManyRelationship{

		schema.ManyRelationship{
			Parent:      Schema.Tables["Account"],
			Child:       Schema.Tables["Deck"],
			ChildColumn: Schema.Tables["Deck"].FindColumn("AccountID"),
			Alias:       "",
		},
	}

	Schema.Tables["Deck"].HasMany = []schema.ManyRelationship{

		schema.ManyRelationship{
			Parent:      Schema.Tables["Deck"],
			Child:       Schema.Tables["Card"],
			ChildColumn: Schema.Tables["Card"].FindColumn("DeckID"),
			Alias:       "",
		},
	}

	Schema.Tables["Deck"].ChildOf = []schema.ManyRelationship{

		schema.ManyRelationship{
			Parent:      Schema.Tables["Account"],
			Child:       Schema.Tables["Deck"],
			ChildColumn: Schema.Tables["Deck"].FindColumn("AccountID"),
			Alias:       "",
		},
	}

	Schema.Tables["Card"].ChildOf = []schema.ManyRelationship{

		schema.ManyRelationship{
			Parent:      Schema.Tables["Deck"],
			Child:       Schema.Tables["Card"],
			ChildColumn: Schema.Tables["Card"].FindColumn("DeckID"),
			Alias:       "",
		},
	}

}

func (t Account) Scope() AccountScope {
	return t.cached_conn.Account.ID().Eq(t.ID)
}

func (t Account) ToScope(c *Conn) AccountScope {
	return c.Account.ID().Eq(t.ID)
}

func (t *Account) Save(c *Conn) error {

	// check the primary key vs the zero value, if they match then
	// we will assume we have a new record
	var pkz int
	if t.ID == pkz {
		return t.create(c)
	} else {
		return t.update(c)
	}
}

func (t *Account) simpleCols(c *Conn) []string {
	return []string{c.SQLColumn("Account", "Email"), c.SQLColumn("Account", "Management"), c.SQLColumn("Account", "CryptPassword")}
}

func (t *Account) simpleVals() []interface{} {
	return []interface{}{t.Email, t.Management, t.CryptPassword}
}

func (t *Account) create(c *Conn) error {
	cols := t.simpleCols(c)
	vals := t.simpleVals()

	pk, err := createRecord(c, cols, vals, "Account", "ID")
	if err == nil {
		t.ID = pk
		t.cached_conn = c
	}
	return err
}

func (t *Account) update(c *Conn) error {
	if c == nil {
		return updateRecord(t.cached_conn, t.simpleCols(t.cached_conn), append(t.simpleVals(), t.ID), "Account", "ID")
	} else {
		return updateRecord(c, t.simpleCols(c), append(t.simpleVals(), t.ID), "Account", "ID")
	}
}

func (t Account) Delete(c *Conn) error {
	return deleteRecord(c, t.ID, "Account", "ID")
}

func (t Deck) Scope() DeckScope {
	return t.cached_conn.Deck.ID().Eq(t.ID)
}

func (t Deck) ToScope(c *Conn) DeckScope {
	return c.Deck.ID().Eq(t.ID)
}

func (t *Deck) Save(c *Conn) error {

	// check the primary key vs the zero value, if they match then
	// we will assume we have a new record
	var pkz int
	if t.ID == pkz {
		return t.create(c)
	} else {
		return t.update(c)
	}
}

func (t *Deck) simpleCols(c *Conn) []string {
	return []string{c.SQLColumn("Deck", "Name"), c.SQLColumn("Deck", "Private"), c.SQLColumn("Deck", "Full"), c.SQLColumn("Deck", "GameType"), c.SQLColumn("Deck", "AccountID")}
}

func (t *Deck) simpleVals() []interface{} {
	return []interface{}{t.Name, t.Private, t.Full, t.GameType, t.AccountID}
}

func (t *Deck) create(c *Conn) error {
	cols := t.simpleCols(c)
	vals := t.simpleVals()

	pk, err := createRecord(c, cols, vals, "Deck", "ID")
	if err == nil {
		t.ID = pk
		t.cached_conn = c
	}
	return err
}

func (t *Deck) update(c *Conn) error {
	if c == nil {
		return updateRecord(t.cached_conn, t.simpleCols(t.cached_conn), append(t.simpleVals(), t.ID), "Deck", "ID")
	} else {
		return updateRecord(c, t.simpleCols(c), append(t.simpleVals(), t.ID), "Deck", "ID")
	}
}

func (t Deck) Delete(c *Conn) error {
	return deleteRecord(c, t.ID, "Deck", "ID")
}

func (t Card) Scope() CardScope {
	return t.cached_conn.Card.ID().Eq(t.ID)
}

func (t Card) ToScope(c *Conn) CardScope {
	return c.Card.ID().Eq(t.ID)
}

func (t *Card) Save(c *Conn) error {

	// check the primary key vs the zero value, if they match then
	// we will assume we have a new record
	var pkz int
	if t.ID == pkz {
		return t.create(c)
	} else {
		return t.update(c)
	}
}

func (t *Card) simpleCols(c *Conn) []string {
	return []string{c.SQLColumn("Card", "Name"), c.SQLColumn("Card", "Type"), c.SQLColumn("Card", "Data"), c.SQLColumn("Card", "DeckID")}
}

func (t *Card) simpleVals() []interface{} {
	return []interface{}{t.Name, t.Type, t.Data, t.DeckID}
}

func (t *Card) create(c *Conn) error {
	cols := t.simpleCols(c)
	vals := t.simpleVals()

	pk, err := createRecord(c, cols, vals, "Card", "ID")
	if err == nil {
		t.ID = pk
		t.cached_conn = c
	}
	return err
}

func (t *Card) update(c *Conn) error {
	if c == nil {
		return updateRecord(t.cached_conn, t.simpleCols(t.cached_conn), append(t.simpleVals(), t.ID), "Card", "ID")
	} else {
		return updateRecord(c, t.simpleCols(c), append(t.simpleVals(), t.ID), "Card", "ID")
	}
}

func (t Card) Delete(c *Conn) error {
	return deleteRecord(c, t.ID, "Card", "ID")
}
