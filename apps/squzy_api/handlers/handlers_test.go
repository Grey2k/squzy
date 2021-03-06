package handlers

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	apiPb "github.com/squzy/squzy_generated/generated/proto/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

type mockAmOk struct {
}

func (m mockAmOk) GetApplicationListByAgentId(ctx context.Context, in *apiPb.AgentIdRequest, opts ...grpc.CallOption) (*apiPb.GetApplicationListResponse, error) {
	panic("implement me")
}

func (m mockAmOk) ArchiveApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return &apiPb.Application{}, nil
}

func (m mockAmOk) EnableApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return &apiPb.Application{}, nil
}

func (m mockAmOk) DisableApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return &apiPb.Application{}, nil
}

func (m mockAmOk) InitializeApplication(ctx context.Context, in *apiPb.ApplicationInfo, opts ...grpc.CallOption) (*apiPb.InitializeApplicationResponse, error) {
	return &apiPb.InitializeApplicationResponse{}, nil
}

func (m mockAmOk) SaveTransaction(ctx context.Context, in *apiPb.TransactionInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (m mockAmOk) GetApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return &apiPb.Application{}, nil
}

func (m mockAmOk) GetApplicationList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*apiPb.GetApplicationListResponse, error) {
	return &apiPb.GetApplicationListResponse{}, nil
}

type mockAmError struct {
}

func (m mockAmError) GetApplicationListByAgentId(ctx context.Context, in *apiPb.AgentIdRequest, opts ...grpc.CallOption) (*apiPb.GetApplicationListResponse, error) {
	panic("implement me")
}

func (m mockAmError) ArchiveApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return nil, errors.New("")
}

func (m mockAmError) EnableApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return nil, errors.New("")
}

func (m mockAmError) DisableApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return nil, errors.New("")
}

func (m mockAmError) InitializeApplication(ctx context.Context, in *apiPb.ApplicationInfo, opts ...grpc.CallOption) (*apiPb.InitializeApplicationResponse, error) {
	return nil, errors.New("")
}

func (m mockAmError) SaveTransaction(ctx context.Context, in *apiPb.TransactionInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	return nil, errors.New("")
}

func (m mockAmError) GetApplicationById(ctx context.Context, in *apiPb.ApplicationByIdReuqest, opts ...grpc.CallOption) (*apiPb.Application, error) {
	return nil, errors.New("")
}

func (m mockAmError) GetApplicationList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*apiPb.GetApplicationListResponse, error) {
	return nil, errors.New("")
}

type storageMockOk struct {
}

func (s storageMockOk) GetSchedulerUptime(ctx context.Context, in *apiPb.GetSchedulerUptimeRequest, opts ...grpc.CallOption) (*apiPb.GetSchedulerUptimeResponse, error) {
	return &apiPb.GetSchedulerUptimeResponse{}, nil
}

func (s storageMockOk) SaveTransaction(ctx context.Context, in *apiPb.TransactionInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s storageMockOk) GetTransactionsGroup(ctx context.Context, in *apiPb.GetTransactionGroupRequest, opts ...grpc.CallOption) (*apiPb.GetTransactionGroupResponse, error) {
	return &apiPb.GetTransactionGroupResponse{}, nil
}

func (s storageMockOk) GetTransactions(ctx context.Context, in *apiPb.GetTransactionsRequest, opts ...grpc.CallOption) (*apiPb.GetTransactionsResponse, error) {
	return &apiPb.GetTransactionsResponse{}, nil
}

func (s storageMockOk) GetTransactionById(ctx context.Context, in *apiPb.GetTransactionByIdRequest, opts ...grpc.CallOption) (*apiPb.GetTransactionByIdResponse, error) {
	return &apiPb.GetTransactionByIdResponse{}, nil
}

func (s storageMockOk) SaveResponseFromScheduler(ctx context.Context, in *apiPb.SchedulerResponse, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s storageMockOk) SaveResponseFromAgent(ctx context.Context, in *apiPb.Metric, opts ...grpc.CallOption) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s storageMockOk) GetSchedulerInformation(ctx context.Context, in *apiPb.GetSchedulerInformationRequest, opts ...grpc.CallOption) (*apiPb.GetSchedulerInformationResponse, error) {
	return &apiPb.GetSchedulerInformationResponse{}, nil
}

func (s storageMockOk) GetAgentInformation(ctx context.Context, in *apiPb.GetAgentInformationRequest, opts ...grpc.CallOption) (*apiPb.GetAgentInformationResponse, error) {
	return &apiPb.GetAgentInformationResponse{}, nil
}

type storageMockError struct {
}

func (s storageMockError) GetSchedulerUptime(ctx context.Context, in *apiPb.GetSchedulerUptimeRequest, opts ...grpc.CallOption) (*apiPb.GetSchedulerUptimeResponse, error) {
	return nil, errors.New("")
}

func (s storageMockError) SaveTransaction(ctx context.Context, in *apiPb.TransactionInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	return nil, errors.New("")
}

func (s storageMockError) GetTransactionsGroup(ctx context.Context, in *apiPb.GetTransactionGroupRequest, opts ...grpc.CallOption) (*apiPb.GetTransactionGroupResponse, error) {
	return nil, errors.New("")
}

func (s storageMockError) GetTransactions(ctx context.Context, in *apiPb.GetTransactionsRequest, opts ...grpc.CallOption) (*apiPb.GetTransactionsResponse, error) {
	return nil, errors.New("")
}

func (s storageMockError) GetTransactionById(ctx context.Context, in *apiPb.GetTransactionByIdRequest, opts ...grpc.CallOption) (*apiPb.GetTransactionByIdResponse, error) {
	return nil, errors.New("")
}

func (s storageMockError) SaveResponseFromScheduler(ctx context.Context, in *apiPb.SchedulerResponse, opts ...grpc.CallOption) (*empty.Empty, error) {
	return nil, errors.New("")
}

func (s storageMockError) SaveResponseFromAgent(ctx context.Context, in *apiPb.Metric, opts ...grpc.CallOption) (*empty.Empty, error) {
	return nil, errors.New("")
}

func (s storageMockError) GetSchedulerInformation(ctx context.Context, in *apiPb.GetSchedulerInformationRequest, opts ...grpc.CallOption) (*apiPb.GetSchedulerInformationResponse, error) {
	return nil, errors.New("")
}

func (s storageMockError) GetAgentInformation(ctx context.Context, in *apiPb.GetAgentInformationRequest, opts ...grpc.CallOption) (*apiPb.GetAgentInformationResponse, error) {
	return nil, errors.New("")
}

type agentMockOk struct {
}

func (a agentMockOk) Register(ctx context.Context, in *apiPb.RegisterRequest, opts ...grpc.CallOption) (*apiPb.RegisterResponse, error) {
	panic("implement me")
}

func (a agentMockOk) GetByAgentName(ctx context.Context, in *apiPb.GetByAgentNameRequest, opts ...grpc.CallOption) (*apiPb.GetAgentListResponse, error) {
	panic("implement me")
}

func (a agentMockOk) GetAgentById(ctx context.Context, in *apiPb.GetAgentByIdRequest, opts ...grpc.CallOption) (*apiPb.AgentItem, error) {
	return &apiPb.AgentItem{}, nil
}

func (a agentMockOk) UnRegister(ctx context.Context, in *apiPb.UnRegisterRequest, opts ...grpc.CallOption) (*apiPb.UnRegisterResponse, error) {
	panic("implement me")
}

func (a agentMockOk) GetAgentList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*apiPb.GetAgentListResponse, error) {
	return &apiPb.GetAgentListResponse{}, nil
}

func (a agentMockOk) SendMetrics(ctx context.Context, opts ...grpc.CallOption) (apiPb.AgentServer_SendMetricsClient, error) {
	panic("implement me")
}

type agentMockError struct {
}

func (a agentMockError) Register(ctx context.Context, in *apiPb.RegisterRequest, opts ...grpc.CallOption) (*apiPb.RegisterResponse, error) {
	panic("implement me")
}

func (a agentMockError) GetByAgentName(ctx context.Context, in *apiPb.GetByAgentNameRequest, opts ...grpc.CallOption) (*apiPb.GetAgentListResponse, error) {
	panic("implement me")
}

func (a agentMockError) GetAgentById(ctx context.Context, in *apiPb.GetAgentByIdRequest, opts ...grpc.CallOption) (*apiPb.AgentItem, error) {
	return nil, errors.New("")
}

func (a agentMockError) UnRegister(ctx context.Context, in *apiPb.UnRegisterRequest, opts ...grpc.CallOption) (*apiPb.UnRegisterResponse, error) {
	panic("implement me")
}

func (a agentMockError) GetAgentList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*apiPb.GetAgentListResponse, error) {
	return nil, errors.New("")
}

func (a agentMockError) SendMetrics(ctx context.Context, opts ...grpc.CallOption) (apiPb.AgentServer_SendMetricsClient, error) {
	panic("implement me")
}

type mockMonitoringError struct {
}

func (m mockMonitoringError) GetSchedulerList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*apiPb.GetSchedulerListResponse, error) {
	return nil, errors.New("")
}

func (m mockMonitoringError) GetSchedulerById(ctx context.Context, in *apiPb.GetSchedulerByIdRequest, opts ...grpc.CallOption) (*apiPb.Scheduler, error) {
	return nil, errors.New("")
}

func (m mockMonitoringError) Add(ctx context.Context, in *apiPb.AddRequest, opts ...grpc.CallOption) (*apiPb.AddResponse, error) {
	return nil, errors.New("")
}

func (m mockMonitoringError) Remove(ctx context.Context, in *apiPb.RemoveRequest, opts ...grpc.CallOption) (*apiPb.RemoveResponse, error) {
	return nil, errors.New("")
}

func (m mockMonitoringError) Run(ctx context.Context, in *apiPb.RunRequest, opts ...grpc.CallOption) (*apiPb.RunResponse, error) {
	return nil, errors.New("")
}

func (m mockMonitoringError) Stop(ctx context.Context, in *apiPb.StopRequest, opts ...grpc.CallOption) (*apiPb.StopResponse, error) {
	return nil, errors.New("")
}

type mockMonitoringOk struct {
}

func (m mockMonitoringOk) GetSchedulerList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*apiPb.GetSchedulerListResponse, error) {
	return &apiPb.GetSchedulerListResponse{}, nil
}

func (m mockMonitoringOk) GetSchedulerById(ctx context.Context, in *apiPb.GetSchedulerByIdRequest, opts ...grpc.CallOption) (*apiPb.Scheduler, error) {
	return &apiPb.Scheduler{}, nil
}

func (m mockMonitoringOk) Add(ctx context.Context, in *apiPb.AddRequest, opts ...grpc.CallOption) (*apiPb.AddResponse, error) {
	return &apiPb.AddResponse{}, nil
}

func (m mockMonitoringOk) Remove(ctx context.Context, in *apiPb.RemoveRequest, opts ...grpc.CallOption) (*apiPb.RemoveResponse, error) {
	return &apiPb.RemoveResponse{}, nil
}

func (m mockMonitoringOk) Run(ctx context.Context, in *apiPb.RunRequest, opts ...grpc.CallOption) (*apiPb.RunResponse, error) {
	return &apiPb.RunResponse{}, nil
}

func (m mockMonitoringOk) Stop(ctx context.Context, in *apiPb.StopRequest, opts ...grpc.CallOption) (*apiPb.StopResponse, error) {
	return &apiPb.StopResponse{}, nil
}

func TestNew(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, nil)
		assert.NotNil(t, s)
	})
}

func TestHandlers_AddScheduler(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringOk{}, nil, nil)
		_, err := s.AddScheduler(context.Background(), &apiPb.AddRequest{})
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringError{}, nil, nil)
		_, err := s.AddScheduler(context.Background(), &apiPb.AddRequest{})
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetAgentByID(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(&agentMockOk{}, nil, nil, nil)
		_, err := s.GetAgentByID(context.Background(), "")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(&agentMockError{}, nil, nil, nil)
		_, err := s.GetAgentByID(context.Background(), "")
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetAgentList(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(&agentMockOk{}, nil, nil, nil)
		_, err := s.GetAgentList(context.Background())
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(&agentMockError{}, nil, nil, nil)
		_, err := s.GetAgentList(context.Background())
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetAgentHistoryByID(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockOk{}, nil)
		_, err := s.GetAgentHistoryByID(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockError{}, nil)
		_, err := s.GetAgentHistoryByID(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetSchedulerHistoryByID(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockOk{}, nil)
		_, err := s.GetSchedulerHistoryByID(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockError{}, nil)
		_, err := s.GetSchedulerHistoryByID(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetSchedulerByID(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringOk{}, nil, nil)
		_, err := s.GetSchedulerByID(context.Background(), "")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringError{}, nil, nil)
		_, err := s.GetSchedulerByID(context.Background(), "")
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetSchedulerList(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringOk{}, nil, nil)
		_, err := s.GetSchedulerList(context.Background())
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringError{}, nil, nil)
		_, err := s.GetSchedulerList(context.Background())
		assert.NotNil(t, err)
	})
}

func TestHandlers_RemoveScheduler(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringOk{}, nil, nil)
		err := s.RemoveScheduler(context.Background(), "nil")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringError{}, nil, nil)
		err := s.RemoveScheduler(context.Background(), "nil")
		assert.NotNil(t, err)
	})
}

func TestHandlers_RunScheduler(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringOk{}, nil, nil)
		err := s.RunScheduler(context.Background(), "nil")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringError{}, nil, nil)
		err := s.RunScheduler(context.Background(), "nil")
		assert.NotNil(t, err)
	})
}

func TestHandlers_StopScheduler(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringOk{}, nil, nil)
		err := s.StopScheduler(context.Background(), "nil")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, &mockMonitoringError{}, nil, nil)
		err := s.StopScheduler(context.Background(), "nil")
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetApplicationById(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.GetApplicationById(context.Background(), "nil")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.GetApplicationById(context.Background(), "nil")
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetApplicationList(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.GetApplicationList(context.Background())
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.GetApplicationList(context.Background())
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetSchedulerUptime(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockOk{}, nil)
		_, err := s.GetSchedulerUptime(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockError{}, nil)
		_, err := s.GetSchedulerUptime(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetTransactionById(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockOk{}, nil)
		_, err := s.GetTransactionById(context.Background(), "")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockError{}, nil)
		_, err := s.GetTransactionById(context.Background(), "nil")
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetTransactionGroups(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockOk{}, nil)
		_, err := s.GetTransactionGroups(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockError{}, nil)
		_, err := s.GetTransactionGroups(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_GetTransactionsList(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockOk{}, nil)
		_, err := s.GetTransactionsList(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, &storageMockError{}, nil)
		_, err := s.GetTransactionsList(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_RegisterApplication(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.RegisterApplication(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.RegisterApplication(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_SaveTransaction(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.SaveTransaction(context.Background(), nil)
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.SaveTransaction(context.Background(), nil)
		assert.NotNil(t, err)
	})
}

func TestHandlers_ArchivedApplicationById(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.ArchivedApplicationById(context.Background(), "")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.ArchivedApplicationById(context.Background(), "")
		assert.NotNil(t, err)
	})
}

func TestHandlers_DisabledApplicationById(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.DisabledApplicationById(context.Background(), "")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.DisabledApplicationById(context.Background(), "")
		assert.NotNil(t, err)
	})
}

func TestHandlers_EnabledApplicationById(t *testing.T) {
	t.Run("Should: not return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmOk{})
		_, err := s.EnabledApplicationById(context.Background(), "")
		assert.Nil(t, err)
	})
	t.Run("Should: return error", func(t *testing.T) {
		s := New(nil, nil, nil, &mockAmError{})
		_, err := s.EnabledApplicationById(context.Background(), "")
		assert.NotNil(t, err)
	})
}
