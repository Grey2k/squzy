package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	apiPb "github.com/squzy/squzy_generated/generated/proto/v1"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"squzy/internal/database"
)

type server struct {
	database database.Database
}

func NewServer(db database.Database) apiPb.StorageServer {
	return &server{
		database: db,
	}
}

func (s *server) SaveResponseFromScheduler(ctx context.Context, request *apiPb.SchedulerResponse) (*empty.Empty, error) {
	err := s.database.InsertSnapshot(request)
	if err != nil {
		return nil, grpcStatus.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s *server) SaveResponseFromAgent(ctx context.Context, request *apiPb.Metric) (*empty.Empty, error) {
	err := s.database.InsertStatRequest(request)
	if err != nil {
		return nil, grpcStatus.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s *server) GetSchedulerInformation(ctx context.Context, request *apiPb.GetSchedulerInformationRequest) (*apiPb.GetSchedulerInformationResponse, error) {
	snapshots, count, err := s.database.GetSnapshots(request)
	return &apiPb.GetSchedulerInformationResponse{
		Snapshots: snapshots,
		Count:     count,
	}, wrapError(err)
}

func (s *server) GetSchedulerUptime(ctx context.Context, request *apiPb.GetSchedulerUptimeRequest) (*apiPb.GetSchedulerUptimeResponse, error) {
	response, err := s.database.GetSnapshotsUptime(request)
	return response, wrapError(err)
}

func (s *server) GetAgentInformation(ctx context.Context, request *apiPb.GetAgentInformationRequest) (*apiPb.GetAgentInformationResponse, error) {
	var res []*apiPb.GetAgentInformationResponse_Statistic
	var count int32
	var err error
	fmt.Println("HERE: " + request.GetAgentId())
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

func (s *server) SaveTransaction(ctx context.Context, info *apiPb.TransactionInfo) (*empty.Empty, error) {
	return &empty.Empty{}, wrapError(s.database.InsertTransactionInfo(info))
}

func (s *server) GetTransactions(ctx context.Context, request *apiPb.GetTransactionsRequest) (*apiPb.GetTransactionsResponse, error) {
	transactions, count, err := s.database.GetTransactionInfo(request)
	return &apiPb.GetTransactionsResponse{
		Count:        count,
		Transactions: transactions,
	}, wrapError(err)
}

func (s *server) GetTransactionById(ctx context.Context, request *apiPb.GetTransactionByIdRequest) (*apiPb.GetTransactionByIdResponse, error) {
	transaction, children, err := s.database.GetTransactionByID(request)
	return &apiPb.GetTransactionByIdResponse{
		Transaction: transaction,
		Children:    children,
	}, wrapError(err)
}

func (s *server) GetTransactionsGroup(ctx context.Context, request *apiPb.GetTransactionGroupRequest) (*apiPb.GetTransactionGroupResponse, error) {
	res, err := s.database.GetTransactionGroup(request)
	return &apiPb.GetTransactionGroupResponse{
		Transactions: res,
	}, wrapError(err)
}

func wrapError(err error) error {
	if err != nil {
		return grpcStatus.Errorf(codes.Internal, err.Error())
	}
	return nil
}
