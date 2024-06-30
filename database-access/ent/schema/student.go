package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
    ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
    }
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("department", Department.Type).
            Ref("students").
            Unique().
            Required(),
        edge.To("enrollments", Enrollment.Type),
    }
}

