package schema

import "entgo.io/ent"

// Mqtt holds the schema definition for the Mqtt entity.
type Mqtt struct {
	ent.Schema
}

// Fields of the Mqtt.
func (Mqtt) Fields() []ent.Field {
	return nil
}

// Edges of the Mqtt.
func (Mqtt) Edges() []ent.Edge {
	return nil
}
