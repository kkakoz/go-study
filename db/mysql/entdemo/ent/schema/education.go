package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Education holds the schema definition for the Education entity.
type Education struct {
	ent.Schema
}

func (Education) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "education"},
	}
}


// Fields of the Education.
func (Education) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
	}
}

// Edges of the Education.
func (Education) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("institution", Institution.Type),
	}
}
