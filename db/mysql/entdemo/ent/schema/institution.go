package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Institution holds the schema definition for the Institution entity.
type Institution struct {
	ent.Schema
}

func (Institution) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "institution"},
	}
}


// Fields of the Institution.
func (Institution) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
	}
}

// Edges of the Institution.
func (Institution) Edges() []ent.Edge {
	return nil
}
