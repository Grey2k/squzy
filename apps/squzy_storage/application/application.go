package application

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	apiPb "github.com/squzy/squzy_generated/generated/proto/v1"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"squzy/internal/database"
)

type service struct {
	database database.Database
}

func NewService(db database.Database) apiPb.StorageServer {
	return &service{
		database: db,
	}
}

func (s *service) SendResponseFromScheduler(ctx context.Context, request *apiPb.SchedulerResponse) (*empty.Empty, error) {
	err := s.database.InsertSnapshot(request)
	if err != nil {
		return nil, grpcStatus.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s *service) SendResponseFromAgent(ctx context.Context, request *apiPb.Metric) (*empty.Empty, error) {
	err := s.database.InsertStatRequest(request)
	if err != nil {
		return nil, grpcStatus.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s *service) GetSchedulerInformation(ctx context.Context, request *apiPb.GetSchedulerInformationRequest) (*apiPb.GetSchedulerInformationResponse, error) {
	snapshots, err := s.database.GetSnapshots(request.SchedulerId)
	return &apiPb.GetSchedulerInformationResponse{
		Snapshots: snapshots,
	}, wrapError(err)
}

func (s *service) GetAgentInformation(ctx context.Context, request *apiPb.GetAgentInformationRequest) (*apiPb.GetAgentInformationResponse, error) {
	var res []*apiPb.GetAgentInformationResponse_Statistic
	var count int32
	var err error
	switch request.GetType() {
	case apiPb.TypeAgentStat_ALL:
		res, count, err = s.database.GetStatRequest(request.GetAgentId(), request.GetPagination(), request.GetTimeRange())
	case apiPb.TypeAgentStat_CPU:
		res, count, err = s.database.GetCPUInfo(request.GetAgentId(), request.GetPagination(), request.GetTimeRange())
	case apiPb.TypeAgentStat_MEMORY:
		res, count, err = s.database.GetMemoryInfo(request.GetAgentId(), request.GetPagination(), request.GetTimeRange())
	case apiPb.TypeAgentStat_DISK:
		res, count, err = s.database.GetDiskInfo(request.GetAgentId(), request.GetPagination(), request.GetTimeRange())
	case apiPb.TypeAgentStat_NET:
		res, count, err = s.database.GetNetInfo(request.GetAgentId(), request.GetPagination(), request.GetTimeRange())
	default:
		err = errors.New("invalid type")
	}
	return &apiPb.GetAgentInformationResponse{
		Stats: res,
		Count: count,
	}, wrapError(err)
}

func wrapError(err error) error {
	if err != nil {
		return grpcStatus.Errorf(codes.Internal, err.Error())
	}
	return nil
}
