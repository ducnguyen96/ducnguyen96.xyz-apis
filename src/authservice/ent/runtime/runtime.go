// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/schema"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[1].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescRemindMe is the schema descriptor for remind_me field.
	userDescRemindMe := userFields[2].Descriptor()
	// user.DefaultRemindMe holds the default value on creation for the remind_me field.
	user.DefaultRemindMe = userDescRemindMe.Default.(bool)
	// userDescWeight is the schema descriptor for weight field.
	userDescWeight := userFields[6].Descriptor()
	// user.DefaultWeight holds the default value on creation for the weight field.
	user.DefaultWeight = userDescWeight.Default.(float32)
	// userDescDailyIntake is the schema descriptor for daily_intake field.
	userDescDailyIntake := userFields[7].Descriptor()
	// user.DefaultDailyIntake holds the default value on creation for the daily_intake field.
	user.DefaultDailyIntake = userDescDailyIntake.Default.(int32)
	// userDescContainerImage is the schema descriptor for container_image field.
	userDescContainerImage := userFields[8].Descriptor()
	// user.DefaultContainerImage holds the default value on creation for the container_image field.
	user.DefaultContainerImage = userDescContainerImage.Default.(string)
	// userDescContainerVolume is the schema descriptor for container_volume field.
	userDescContainerVolume := userFields[9].Descriptor()
	// user.DefaultContainerVolume holds the default value on creation for the container_volume field.
	user.DefaultContainerVolume = userDescContainerVolume.Default.(int32)
	// userDescDrinkAtATime is the schema descriptor for drink_at_a_time field.
	userDescDrinkAtATime := userFields[10].Descriptor()
	// user.DefaultDrinkAtATime holds the default value on creation for the drink_at_a_time field.
	user.DefaultDrinkAtATime = userDescDrinkAtATime.Default.(int32)
}

const (
	Version = "v0.9.1"                                          // Version of ent codegen.
	Sum     = "h1:IG8andyeD79GG24U8Q+1Y45hQXj6gY5evSBcva5gtBk=" // Sum of ent codegen.
)
