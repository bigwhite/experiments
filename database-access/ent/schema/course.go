package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
    ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
    return []ent.Field{
        field.String("title").NotEmpty(),
    }
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("department", Department.Type).
            Ref("courses").
            Unique().
            Required(),
        edge.To("enrollments", Enrollment.Type),
    }
}

