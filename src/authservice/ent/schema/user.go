package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("avatar").Default("https://robohash.org/honey?set=set1"),
		field.Bool("remind_me").Default(true),
		field.Time("wake_up_time"),
		field.Time("sleep_time"),
		field.Enum("gender").GoType(Gender(0)),
		field.Float32("weight").Default(50),
		field.Int32("daily_intake").Default(1500),
		field.String("container_image").Default("images/glass-of-water.png"),
		field.Int32("container_volume").Default(300),
		field.Int32("drink_at_a_time").Default(300),
		field.String("password"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}
