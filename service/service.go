package service

import (
	"context"
	"fmt"

	"github.com/konjoot/proto-errors/proto"
	"github.com/micro/go-micro/errors"
)

type Service struct{}

func (s *Service) CreateThingOneOf(ctx context.Context, req *proto.CreateThingOneOfRequest, resp *proto.CreateThingOneOfResponse) error {
	switch req.GetName() {
	case "success":
		resp.Result = &proto.CreateThingOneOfResponse_Thing{
			Thing: &proto.Thing{
				ID:   "id",
				Name: "name",
			},
		}
		return nil
	case "business-error":
		resp.Result = &proto.CreateThingOneOfResponse_Error{
			Error: &proto.Error{
				Code:    proto.ErrorCode_One,
				Message: "can't validate user's data",
				Service: "service-name",
				Details: map[string]string{
					"user":       "user-1",
					"thing_name": req.GetName(),
				},
			},
		}
		return nil
	case "transport-error":
		return &errors.Error{
			Id: "err-id",
			Detail: fmt.Sprintf(
				`"user_id": %q, "thing_name": %q, "service_name": %q`,
				"user-1", req.GetName(), "service-name",
			),
			Code:   int32(proto.ErrorCode_Three),
			Status: "NotBad",
		}
	}

	return errors.NotFound("err-id", "user_id: %s", "user-1")
}

func (s *Service) CreateThingAny(context.Context, *proto.CreateThingAnyRequest, *proto.CreateThingAnyResponse) error {
	return nil
}
