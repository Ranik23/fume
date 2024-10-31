package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Request struct {
	ent.Schema
}

func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			Unique(),
		field.String("endpoint").
			MaxLen(255).
			NotEmpty(),
		field.String("method").
			MaxLen(10).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Request) Edges() []ent.Edge {
	return nil
}
