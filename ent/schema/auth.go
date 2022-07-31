package schema

import "entgo.io/ent"

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return nil
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return nil
}
