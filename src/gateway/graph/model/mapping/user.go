package mapping

import (
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/model"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/utils"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
)

func UserEntity(user *pb.UserEntity) *model.User {
	return &model.User{
		ID:              utils.Uint64ToString(user.Id),
		Username:        user.Username,
		Avatar:          user.Avatar,
		RemindMe:        user.RemindMe,
		WakeUpTime:      user.WakeUpTime,
		SleepTime:       user.SleepTime,
		Gender:          Gender(user.Gender),
		Weight:          utils.Float32ToFloat64(user.Weight),
		DailyIntake:     utils.Int32ToInt(user.DailyIntake),
		ContainerImage:  user.ContainerImage,
		ContainerVolume: utils.Int32ToInt(user.ContainerVolume),
		DrinkAtATime:    utils.Int32ToInt(user.DrinkAtATime),
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}

func Gender(gender pb.Gender) model.Gender {
	switch gender {
	case pb.Gender_FEMALE:
		return model.GenderFemale
	default:
		return model.GenderMale
	}
}