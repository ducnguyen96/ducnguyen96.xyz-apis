package utils

import (
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/schema"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
)

func MapUserRecordToUserProto(u *ent.User) *pb.UserEntity {
	return &pb.UserEntity{
		Id:              u.ID,
		Username:        u.Username,
		Avatar:          u.Avatar,
		RemindMe:        u.RemindMe,
		WakeUpTime:      u.WakeUpTime.String(),
		SleepTime:       u.SleepTime.String(),
		Gender:          EntGenderToProtoGender(u.Gender),
		Weight:          u.Weight,
		DailyIntake:     u.DailyIntake,
		ContainerImage:  u.ContainerImage,
		ContainerVolume: u.ContainerVolume,
		DrinkAtATime:    u.DrinkAtATime,
		CreatedAt:       u.CreatedAt.String(),
		UpdatedAt:       u.UpdatedAt.String(),
	}
}

func EntGenderToProtoGender(g schema.Gender) pb.Gender {
	switch g {
	case schema.Male:
		return pb.Gender_MALE
	default:
		return pb.Gender_FEMALE
	}
}

func IntToOptionalInt32(number int) *int32 {
	num := int32(number)
	return &num
}
