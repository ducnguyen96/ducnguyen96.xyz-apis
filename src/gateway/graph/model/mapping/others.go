package mapping

import (
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/model"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/utils"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
)

func TokenPayloadDto(payload *pb.TokenPayloadDto) *model.TokenPayloadDto {
	return &model.TokenPayloadDto{
		ExpiresIn:   utils.OptionalInt32ToOptionalInt(payload.ExpiresIn),
		AccessToken: payload.AccessToken,
	}
}