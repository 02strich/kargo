// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: service/v1alpha1/service.proto

package svcv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/akuity/kargo/pkg/api/service/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// KargoServiceName is the fully-qualified name of the KargoService service.
	KargoServiceName = "akuity.io.kargo.service.v1alpha1.KargoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// KargoServiceGetVersionInfoProcedure is the fully-qualified name of the KargoService's
	// GetVersionInfo RPC.
	KargoServiceGetVersionInfoProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/GetVersionInfo"
	// KargoServiceGetPublicConfigProcedure is the fully-qualified name of the KargoService's
	// GetPublicConfig RPC.
	KargoServiceGetPublicConfigProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/GetPublicConfig"
	// KargoServiceAdminLoginProcedure is the fully-qualified name of the KargoService's AdminLogin RPC.
	KargoServiceAdminLoginProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/AdminLogin"
	// KargoServiceCreateStageProcedure is the fully-qualified name of the KargoService's CreateStage
	// RPC.
	KargoServiceCreateStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/CreateStage"
	// KargoServiceListStagesProcedure is the fully-qualified name of the KargoService's ListStages RPC.
	KargoServiceListStagesProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/ListStages"
	// KargoServiceGetStageProcedure is the fully-qualified name of the KargoService's GetStage RPC.
	KargoServiceGetStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/GetStage"
	// KargoServiceWatchStageProcedure is the fully-qualified name of the KargoService's WatchStage RPC.
	KargoServiceWatchStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/WatchStage"
	// KargoServiceUpdateStageProcedure is the fully-qualified name of the KargoService's UpdateStage
	// RPC.
	KargoServiceUpdateStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/UpdateStage"
	// KargoServiceDeleteStageProcedure is the fully-qualified name of the KargoService's DeleteStage
	// RPC.
	KargoServiceDeleteStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/DeleteStage"
	// KargoServicePromoteStageProcedure is the fully-qualified name of the KargoService's PromoteStage
	// RPC.
	KargoServicePromoteStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/PromoteStage"
	// KargoServiceSetAutoPromotionForStageProcedure is the fully-qualified name of the KargoService's
	// SetAutoPromotionForStage RPC.
	KargoServiceSetAutoPromotionForStageProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/SetAutoPromotionForStage"
	// KargoServiceCreatePromotionPolicyProcedure is the fully-qualified name of the KargoService's
	// CreatePromotionPolicy RPC.
	KargoServiceCreatePromotionPolicyProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/CreatePromotionPolicy"
	// KargoServiceListPromotionPoliciesProcedure is the fully-qualified name of the KargoService's
	// ListPromotionPolicies RPC.
	KargoServiceListPromotionPoliciesProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/ListPromotionPolicies"
	// KargoServiceGetPromotionPolicyProcedure is the fully-qualified name of the KargoService's
	// GetPromotionPolicy RPC.
	KargoServiceGetPromotionPolicyProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/GetPromotionPolicy"
	// KargoServiceUpdatePromotionPolicyProcedure is the fully-qualified name of the KargoService's
	// UpdatePromotionPolicy RPC.
	KargoServiceUpdatePromotionPolicyProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/UpdatePromotionPolicy"
	// KargoServiceDeletePromotionPolicyProcedure is the fully-qualified name of the KargoService's
	// DeletePromotionPolicy RPC.
	KargoServiceDeletePromotionPolicyProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/DeletePromotionPolicy"
	// KargoServiceCreateProjectProcedure is the fully-qualified name of the KargoService's
	// CreateProject RPC.
	KargoServiceCreateProjectProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/CreateProject"
	// KargoServiceListProjectsProcedure is the fully-qualified name of the KargoService's ListProjects
	// RPC.
	KargoServiceListProjectsProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/ListProjects"
	// KargoServiceDeleteProjectProcedure is the fully-qualified name of the KargoService's
	// DeleteProject RPC.
	KargoServiceDeleteProjectProcedure = "/akuity.io.kargo.service.v1alpha1.KargoService/DeleteProject"
)

// KargoServiceClient is a client for the akuity.io.kargo.service.v1alpha1.KargoService service.
type KargoServiceClient interface {
	GetVersionInfo(context.Context, *connect_go.Request[v1alpha1.GetVersionInfoRequest]) (*connect_go.Response[v1alpha1.GetVersionInfoResponse], error)
	GetPublicConfig(context.Context, *connect_go.Request[v1alpha1.GetPublicConfigRequest]) (*connect_go.Response[v1alpha1.GetPublicConfigResponse], error)
	AdminLogin(context.Context, *connect_go.Request[v1alpha1.AdminLoginRequest]) (*connect_go.Response[v1alpha1.AdminLoginResponse], error)
	CreateStage(context.Context, *connect_go.Request[v1alpha1.CreateStageRequest]) (*connect_go.Response[v1alpha1.CreateStageResponse], error)
	ListStages(context.Context, *connect_go.Request[v1alpha1.ListStagesRequest]) (*connect_go.Response[v1alpha1.ListStagesResponse], error)
	GetStage(context.Context, *connect_go.Request[v1alpha1.GetStageRequest]) (*connect_go.Response[v1alpha1.GetStageResponse], error)
	WatchStage(context.Context, *connect_go.Request[v1alpha1.WatchStageRequest]) (*connect_go.ServerStreamForClient[v1alpha1.WatchStageResponse], error)
	UpdateStage(context.Context, *connect_go.Request[v1alpha1.UpdateStageRequest]) (*connect_go.Response[v1alpha1.UpdateStageResponse], error)
	DeleteStage(context.Context, *connect_go.Request[v1alpha1.DeleteStageRequest]) (*connect_go.Response[v1alpha1.DeleteStageResponse], error)
	PromoteStage(context.Context, *connect_go.Request[v1alpha1.PromoteStageRequest]) (*connect_go.Response[v1alpha1.PromoteStageResponse], error)
	SetAutoPromotionForStage(context.Context, *connect_go.Request[v1alpha1.SetAutoPromotionForStageRequest]) (*connect_go.Response[v1alpha1.SetAutoPromotionForStageResponse], error)
	CreatePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.CreatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.CreatePromotionPolicyResponse], error)
	ListPromotionPolicies(context.Context, *connect_go.Request[v1alpha1.ListPromotionPoliciesRequest]) (*connect_go.Response[v1alpha1.ListPromotionPoliciesResponse], error)
	GetPromotionPolicy(context.Context, *connect_go.Request[v1alpha1.GetPromotionPolicyRequest]) (*connect_go.Response[v1alpha1.GetPromotionPolicyResponse], error)
	UpdatePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.UpdatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.UpdatePromotionPolicyResponse], error)
	DeletePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.DeletePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.DeletePromotionPolicyResponse], error)
	CreateProject(context.Context, *connect_go.Request[v1alpha1.CreateProjectRequest]) (*connect_go.Response[v1alpha1.CreateProjectResponse], error)
	ListProjects(context.Context, *connect_go.Request[v1alpha1.ListProjectsRequest]) (*connect_go.Response[v1alpha1.ListProjectsResponse], error)
	DeleteProject(context.Context, *connect_go.Request[v1alpha1.DeleteProjectRequest]) (*connect_go.Response[v1alpha1.DeleteProjectResponse], error)
}

// NewKargoServiceClient constructs a client for the akuity.io.kargo.service.v1alpha1.KargoService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewKargoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) KargoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &kargoServiceClient{
		getVersionInfo: connect_go.NewClient[v1alpha1.GetVersionInfoRequest, v1alpha1.GetVersionInfoResponse](
			httpClient,
			baseURL+KargoServiceGetVersionInfoProcedure,
			opts...,
		),
		getPublicConfig: connect_go.NewClient[v1alpha1.GetPublicConfigRequest, v1alpha1.GetPublicConfigResponse](
			httpClient,
			baseURL+KargoServiceGetPublicConfigProcedure,
			opts...,
		),
		adminLogin: connect_go.NewClient[v1alpha1.AdminLoginRequest, v1alpha1.AdminLoginResponse](
			httpClient,
			baseURL+KargoServiceAdminLoginProcedure,
			opts...,
		),
		createStage: connect_go.NewClient[v1alpha1.CreateStageRequest, v1alpha1.CreateStageResponse](
			httpClient,
			baseURL+KargoServiceCreateStageProcedure,
			opts...,
		),
		listStages: connect_go.NewClient[v1alpha1.ListStagesRequest, v1alpha1.ListStagesResponse](
			httpClient,
			baseURL+KargoServiceListStagesProcedure,
			opts...,
		),
		getStage: connect_go.NewClient[v1alpha1.GetStageRequest, v1alpha1.GetStageResponse](
			httpClient,
			baseURL+KargoServiceGetStageProcedure,
			opts...,
		),
		watchStage: connect_go.NewClient[v1alpha1.WatchStageRequest, v1alpha1.WatchStageResponse](
			httpClient,
			baseURL+KargoServiceWatchStageProcedure,
			opts...,
		),
		updateStage: connect_go.NewClient[v1alpha1.UpdateStageRequest, v1alpha1.UpdateStageResponse](
			httpClient,
			baseURL+KargoServiceUpdateStageProcedure,
			opts...,
		),
		deleteStage: connect_go.NewClient[v1alpha1.DeleteStageRequest, v1alpha1.DeleteStageResponse](
			httpClient,
			baseURL+KargoServiceDeleteStageProcedure,
			opts...,
		),
		promoteStage: connect_go.NewClient[v1alpha1.PromoteStageRequest, v1alpha1.PromoteStageResponse](
			httpClient,
			baseURL+KargoServicePromoteStageProcedure,
			opts...,
		),
		setAutoPromotionForStage: connect_go.NewClient[v1alpha1.SetAutoPromotionForStageRequest, v1alpha1.SetAutoPromotionForStageResponse](
			httpClient,
			baseURL+KargoServiceSetAutoPromotionForStageProcedure,
			opts...,
		),
		createPromotionPolicy: connect_go.NewClient[v1alpha1.CreatePromotionPolicyRequest, v1alpha1.CreatePromotionPolicyResponse](
			httpClient,
			baseURL+KargoServiceCreatePromotionPolicyProcedure,
			opts...,
		),
		listPromotionPolicies: connect_go.NewClient[v1alpha1.ListPromotionPoliciesRequest, v1alpha1.ListPromotionPoliciesResponse](
			httpClient,
			baseURL+KargoServiceListPromotionPoliciesProcedure,
			opts...,
		),
		getPromotionPolicy: connect_go.NewClient[v1alpha1.GetPromotionPolicyRequest, v1alpha1.GetPromotionPolicyResponse](
			httpClient,
			baseURL+KargoServiceGetPromotionPolicyProcedure,
			opts...,
		),
		updatePromotionPolicy: connect_go.NewClient[v1alpha1.UpdatePromotionPolicyRequest, v1alpha1.UpdatePromotionPolicyResponse](
			httpClient,
			baseURL+KargoServiceUpdatePromotionPolicyProcedure,
			opts...,
		),
		deletePromotionPolicy: connect_go.NewClient[v1alpha1.DeletePromotionPolicyRequest, v1alpha1.DeletePromotionPolicyResponse](
			httpClient,
			baseURL+KargoServiceDeletePromotionPolicyProcedure,
			opts...,
		),
		createProject: connect_go.NewClient[v1alpha1.CreateProjectRequest, v1alpha1.CreateProjectResponse](
			httpClient,
			baseURL+KargoServiceCreateProjectProcedure,
			opts...,
		),
		listProjects: connect_go.NewClient[v1alpha1.ListProjectsRequest, v1alpha1.ListProjectsResponse](
			httpClient,
			baseURL+KargoServiceListProjectsProcedure,
			opts...,
		),
		deleteProject: connect_go.NewClient[v1alpha1.DeleteProjectRequest, v1alpha1.DeleteProjectResponse](
			httpClient,
			baseURL+KargoServiceDeleteProjectProcedure,
			opts...,
		),
	}
}

// kargoServiceClient implements KargoServiceClient.
type kargoServiceClient struct {
	getVersionInfo           *connect_go.Client[v1alpha1.GetVersionInfoRequest, v1alpha1.GetVersionInfoResponse]
	getPublicConfig          *connect_go.Client[v1alpha1.GetPublicConfigRequest, v1alpha1.GetPublicConfigResponse]
	adminLogin               *connect_go.Client[v1alpha1.AdminLoginRequest, v1alpha1.AdminLoginResponse]
	createStage              *connect_go.Client[v1alpha1.CreateStageRequest, v1alpha1.CreateStageResponse]
	listStages               *connect_go.Client[v1alpha1.ListStagesRequest, v1alpha1.ListStagesResponse]
	getStage                 *connect_go.Client[v1alpha1.GetStageRequest, v1alpha1.GetStageResponse]
	watchStage               *connect_go.Client[v1alpha1.WatchStageRequest, v1alpha1.WatchStageResponse]
	updateStage              *connect_go.Client[v1alpha1.UpdateStageRequest, v1alpha1.UpdateStageResponse]
	deleteStage              *connect_go.Client[v1alpha1.DeleteStageRequest, v1alpha1.DeleteStageResponse]
	promoteStage             *connect_go.Client[v1alpha1.PromoteStageRequest, v1alpha1.PromoteStageResponse]
	setAutoPromotionForStage *connect_go.Client[v1alpha1.SetAutoPromotionForStageRequest, v1alpha1.SetAutoPromotionForStageResponse]
	createPromotionPolicy    *connect_go.Client[v1alpha1.CreatePromotionPolicyRequest, v1alpha1.CreatePromotionPolicyResponse]
	listPromotionPolicies    *connect_go.Client[v1alpha1.ListPromotionPoliciesRequest, v1alpha1.ListPromotionPoliciesResponse]
	getPromotionPolicy       *connect_go.Client[v1alpha1.GetPromotionPolicyRequest, v1alpha1.GetPromotionPolicyResponse]
	updatePromotionPolicy    *connect_go.Client[v1alpha1.UpdatePromotionPolicyRequest, v1alpha1.UpdatePromotionPolicyResponse]
	deletePromotionPolicy    *connect_go.Client[v1alpha1.DeletePromotionPolicyRequest, v1alpha1.DeletePromotionPolicyResponse]
	createProject            *connect_go.Client[v1alpha1.CreateProjectRequest, v1alpha1.CreateProjectResponse]
	listProjects             *connect_go.Client[v1alpha1.ListProjectsRequest, v1alpha1.ListProjectsResponse]
	deleteProject            *connect_go.Client[v1alpha1.DeleteProjectRequest, v1alpha1.DeleteProjectResponse]
}

// GetVersionInfo calls akuity.io.kargo.service.v1alpha1.KargoService.GetVersionInfo.
func (c *kargoServiceClient) GetVersionInfo(ctx context.Context, req *connect_go.Request[v1alpha1.GetVersionInfoRequest]) (*connect_go.Response[v1alpha1.GetVersionInfoResponse], error) {
	return c.getVersionInfo.CallUnary(ctx, req)
}

// GetPublicConfig calls akuity.io.kargo.service.v1alpha1.KargoService.GetPublicConfig.
func (c *kargoServiceClient) GetPublicConfig(ctx context.Context, req *connect_go.Request[v1alpha1.GetPublicConfigRequest]) (*connect_go.Response[v1alpha1.GetPublicConfigResponse], error) {
	return c.getPublicConfig.CallUnary(ctx, req)
}

// AdminLogin calls akuity.io.kargo.service.v1alpha1.KargoService.AdminLogin.
func (c *kargoServiceClient) AdminLogin(ctx context.Context, req *connect_go.Request[v1alpha1.AdminLoginRequest]) (*connect_go.Response[v1alpha1.AdminLoginResponse], error) {
	return c.adminLogin.CallUnary(ctx, req)
}

// CreateStage calls akuity.io.kargo.service.v1alpha1.KargoService.CreateStage.
func (c *kargoServiceClient) CreateStage(ctx context.Context, req *connect_go.Request[v1alpha1.CreateStageRequest]) (*connect_go.Response[v1alpha1.CreateStageResponse], error) {
	return c.createStage.CallUnary(ctx, req)
}

// ListStages calls akuity.io.kargo.service.v1alpha1.KargoService.ListStages.
func (c *kargoServiceClient) ListStages(ctx context.Context, req *connect_go.Request[v1alpha1.ListStagesRequest]) (*connect_go.Response[v1alpha1.ListStagesResponse], error) {
	return c.listStages.CallUnary(ctx, req)
}

// GetStage calls akuity.io.kargo.service.v1alpha1.KargoService.GetStage.
func (c *kargoServiceClient) GetStage(ctx context.Context, req *connect_go.Request[v1alpha1.GetStageRequest]) (*connect_go.Response[v1alpha1.GetStageResponse], error) {
	return c.getStage.CallUnary(ctx, req)
}

// WatchStage calls akuity.io.kargo.service.v1alpha1.KargoService.WatchStage.
func (c *kargoServiceClient) WatchStage(ctx context.Context, req *connect_go.Request[v1alpha1.WatchStageRequest]) (*connect_go.ServerStreamForClient[v1alpha1.WatchStageResponse], error) {
	return c.watchStage.CallServerStream(ctx, req)
}

// UpdateStage calls akuity.io.kargo.service.v1alpha1.KargoService.UpdateStage.
func (c *kargoServiceClient) UpdateStage(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateStageRequest]) (*connect_go.Response[v1alpha1.UpdateStageResponse], error) {
	return c.updateStage.CallUnary(ctx, req)
}

// DeleteStage calls akuity.io.kargo.service.v1alpha1.KargoService.DeleteStage.
func (c *kargoServiceClient) DeleteStage(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteStageRequest]) (*connect_go.Response[v1alpha1.DeleteStageResponse], error) {
	return c.deleteStage.CallUnary(ctx, req)
}

// PromoteStage calls akuity.io.kargo.service.v1alpha1.KargoService.PromoteStage.
func (c *kargoServiceClient) PromoteStage(ctx context.Context, req *connect_go.Request[v1alpha1.PromoteStageRequest]) (*connect_go.Response[v1alpha1.PromoteStageResponse], error) {
	return c.promoteStage.CallUnary(ctx, req)
}

// SetAutoPromotionForStage calls
// akuity.io.kargo.service.v1alpha1.KargoService.SetAutoPromotionForStage.
func (c *kargoServiceClient) SetAutoPromotionForStage(ctx context.Context, req *connect_go.Request[v1alpha1.SetAutoPromotionForStageRequest]) (*connect_go.Response[v1alpha1.SetAutoPromotionForStageResponse], error) {
	return c.setAutoPromotionForStage.CallUnary(ctx, req)
}

// CreatePromotionPolicy calls akuity.io.kargo.service.v1alpha1.KargoService.CreatePromotionPolicy.
func (c *kargoServiceClient) CreatePromotionPolicy(ctx context.Context, req *connect_go.Request[v1alpha1.CreatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.CreatePromotionPolicyResponse], error) {
	return c.createPromotionPolicy.CallUnary(ctx, req)
}

// ListPromotionPolicies calls akuity.io.kargo.service.v1alpha1.KargoService.ListPromotionPolicies.
func (c *kargoServiceClient) ListPromotionPolicies(ctx context.Context, req *connect_go.Request[v1alpha1.ListPromotionPoliciesRequest]) (*connect_go.Response[v1alpha1.ListPromotionPoliciesResponse], error) {
	return c.listPromotionPolicies.CallUnary(ctx, req)
}

// GetPromotionPolicy calls akuity.io.kargo.service.v1alpha1.KargoService.GetPromotionPolicy.
func (c *kargoServiceClient) GetPromotionPolicy(ctx context.Context, req *connect_go.Request[v1alpha1.GetPromotionPolicyRequest]) (*connect_go.Response[v1alpha1.GetPromotionPolicyResponse], error) {
	return c.getPromotionPolicy.CallUnary(ctx, req)
}

// UpdatePromotionPolicy calls akuity.io.kargo.service.v1alpha1.KargoService.UpdatePromotionPolicy.
func (c *kargoServiceClient) UpdatePromotionPolicy(ctx context.Context, req *connect_go.Request[v1alpha1.UpdatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.UpdatePromotionPolicyResponse], error) {
	return c.updatePromotionPolicy.CallUnary(ctx, req)
}

// DeletePromotionPolicy calls akuity.io.kargo.service.v1alpha1.KargoService.DeletePromotionPolicy.
func (c *kargoServiceClient) DeletePromotionPolicy(ctx context.Context, req *connect_go.Request[v1alpha1.DeletePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.DeletePromotionPolicyResponse], error) {
	return c.deletePromotionPolicy.CallUnary(ctx, req)
}

// CreateProject calls akuity.io.kargo.service.v1alpha1.KargoService.CreateProject.
func (c *kargoServiceClient) CreateProject(ctx context.Context, req *connect_go.Request[v1alpha1.CreateProjectRequest]) (*connect_go.Response[v1alpha1.CreateProjectResponse], error) {
	return c.createProject.CallUnary(ctx, req)
}

// ListProjects calls akuity.io.kargo.service.v1alpha1.KargoService.ListProjects.
func (c *kargoServiceClient) ListProjects(ctx context.Context, req *connect_go.Request[v1alpha1.ListProjectsRequest]) (*connect_go.Response[v1alpha1.ListProjectsResponse], error) {
	return c.listProjects.CallUnary(ctx, req)
}

// DeleteProject calls akuity.io.kargo.service.v1alpha1.KargoService.DeleteProject.
func (c *kargoServiceClient) DeleteProject(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteProjectRequest]) (*connect_go.Response[v1alpha1.DeleteProjectResponse], error) {
	return c.deleteProject.CallUnary(ctx, req)
}

// KargoServiceHandler is an implementation of the akuity.io.kargo.service.v1alpha1.KargoService
// service.
type KargoServiceHandler interface {
	GetVersionInfo(context.Context, *connect_go.Request[v1alpha1.GetVersionInfoRequest]) (*connect_go.Response[v1alpha1.GetVersionInfoResponse], error)
	GetPublicConfig(context.Context, *connect_go.Request[v1alpha1.GetPublicConfigRequest]) (*connect_go.Response[v1alpha1.GetPublicConfigResponse], error)
	AdminLogin(context.Context, *connect_go.Request[v1alpha1.AdminLoginRequest]) (*connect_go.Response[v1alpha1.AdminLoginResponse], error)
	CreateStage(context.Context, *connect_go.Request[v1alpha1.CreateStageRequest]) (*connect_go.Response[v1alpha1.CreateStageResponse], error)
	ListStages(context.Context, *connect_go.Request[v1alpha1.ListStagesRequest]) (*connect_go.Response[v1alpha1.ListStagesResponse], error)
	GetStage(context.Context, *connect_go.Request[v1alpha1.GetStageRequest]) (*connect_go.Response[v1alpha1.GetStageResponse], error)
	WatchStage(context.Context, *connect_go.Request[v1alpha1.WatchStageRequest], *connect_go.ServerStream[v1alpha1.WatchStageResponse]) error
	UpdateStage(context.Context, *connect_go.Request[v1alpha1.UpdateStageRequest]) (*connect_go.Response[v1alpha1.UpdateStageResponse], error)
	DeleteStage(context.Context, *connect_go.Request[v1alpha1.DeleteStageRequest]) (*connect_go.Response[v1alpha1.DeleteStageResponse], error)
	PromoteStage(context.Context, *connect_go.Request[v1alpha1.PromoteStageRequest]) (*connect_go.Response[v1alpha1.PromoteStageResponse], error)
	SetAutoPromotionForStage(context.Context, *connect_go.Request[v1alpha1.SetAutoPromotionForStageRequest]) (*connect_go.Response[v1alpha1.SetAutoPromotionForStageResponse], error)
	CreatePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.CreatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.CreatePromotionPolicyResponse], error)
	ListPromotionPolicies(context.Context, *connect_go.Request[v1alpha1.ListPromotionPoliciesRequest]) (*connect_go.Response[v1alpha1.ListPromotionPoliciesResponse], error)
	GetPromotionPolicy(context.Context, *connect_go.Request[v1alpha1.GetPromotionPolicyRequest]) (*connect_go.Response[v1alpha1.GetPromotionPolicyResponse], error)
	UpdatePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.UpdatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.UpdatePromotionPolicyResponse], error)
	DeletePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.DeletePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.DeletePromotionPolicyResponse], error)
	CreateProject(context.Context, *connect_go.Request[v1alpha1.CreateProjectRequest]) (*connect_go.Response[v1alpha1.CreateProjectResponse], error)
	ListProjects(context.Context, *connect_go.Request[v1alpha1.ListProjectsRequest]) (*connect_go.Response[v1alpha1.ListProjectsResponse], error)
	DeleteProject(context.Context, *connect_go.Request[v1alpha1.DeleteProjectRequest]) (*connect_go.Response[v1alpha1.DeleteProjectResponse], error)
}

// NewKargoServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewKargoServiceHandler(svc KargoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(KargoServiceGetVersionInfoProcedure, connect_go.NewUnaryHandler(
		KargoServiceGetVersionInfoProcedure,
		svc.GetVersionInfo,
		opts...,
	))
	mux.Handle(KargoServiceGetPublicConfigProcedure, connect_go.NewUnaryHandler(
		KargoServiceGetPublicConfigProcedure,
		svc.GetPublicConfig,
		opts...,
	))
	mux.Handle(KargoServiceAdminLoginProcedure, connect_go.NewUnaryHandler(
		KargoServiceAdminLoginProcedure,
		svc.AdminLogin,
		opts...,
	))
	mux.Handle(KargoServiceCreateStageProcedure, connect_go.NewUnaryHandler(
		KargoServiceCreateStageProcedure,
		svc.CreateStage,
		opts...,
	))
	mux.Handle(KargoServiceListStagesProcedure, connect_go.NewUnaryHandler(
		KargoServiceListStagesProcedure,
		svc.ListStages,
		opts...,
	))
	mux.Handle(KargoServiceGetStageProcedure, connect_go.NewUnaryHandler(
		KargoServiceGetStageProcedure,
		svc.GetStage,
		opts...,
	))
	mux.Handle(KargoServiceWatchStageProcedure, connect_go.NewServerStreamHandler(
		KargoServiceWatchStageProcedure,
		svc.WatchStage,
		opts...,
	))
	mux.Handle(KargoServiceUpdateStageProcedure, connect_go.NewUnaryHandler(
		KargoServiceUpdateStageProcedure,
		svc.UpdateStage,
		opts...,
	))
	mux.Handle(KargoServiceDeleteStageProcedure, connect_go.NewUnaryHandler(
		KargoServiceDeleteStageProcedure,
		svc.DeleteStage,
		opts...,
	))
	mux.Handle(KargoServicePromoteStageProcedure, connect_go.NewUnaryHandler(
		KargoServicePromoteStageProcedure,
		svc.PromoteStage,
		opts...,
	))
	mux.Handle(KargoServiceSetAutoPromotionForStageProcedure, connect_go.NewUnaryHandler(
		KargoServiceSetAutoPromotionForStageProcedure,
		svc.SetAutoPromotionForStage,
		opts...,
	))
	mux.Handle(KargoServiceCreatePromotionPolicyProcedure, connect_go.NewUnaryHandler(
		KargoServiceCreatePromotionPolicyProcedure,
		svc.CreatePromotionPolicy,
		opts...,
	))
	mux.Handle(KargoServiceListPromotionPoliciesProcedure, connect_go.NewUnaryHandler(
		KargoServiceListPromotionPoliciesProcedure,
		svc.ListPromotionPolicies,
		opts...,
	))
	mux.Handle(KargoServiceGetPromotionPolicyProcedure, connect_go.NewUnaryHandler(
		KargoServiceGetPromotionPolicyProcedure,
		svc.GetPromotionPolicy,
		opts...,
	))
	mux.Handle(KargoServiceUpdatePromotionPolicyProcedure, connect_go.NewUnaryHandler(
		KargoServiceUpdatePromotionPolicyProcedure,
		svc.UpdatePromotionPolicy,
		opts...,
	))
	mux.Handle(KargoServiceDeletePromotionPolicyProcedure, connect_go.NewUnaryHandler(
		KargoServiceDeletePromotionPolicyProcedure,
		svc.DeletePromotionPolicy,
		opts...,
	))
	mux.Handle(KargoServiceCreateProjectProcedure, connect_go.NewUnaryHandler(
		KargoServiceCreateProjectProcedure,
		svc.CreateProject,
		opts...,
	))
	mux.Handle(KargoServiceListProjectsProcedure, connect_go.NewUnaryHandler(
		KargoServiceListProjectsProcedure,
		svc.ListProjects,
		opts...,
	))
	mux.Handle(KargoServiceDeleteProjectProcedure, connect_go.NewUnaryHandler(
		KargoServiceDeleteProjectProcedure,
		svc.DeleteProject,
		opts...,
	))
	return "/akuity.io.kargo.service.v1alpha1.KargoService/", mux
}

// UnimplementedKargoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedKargoServiceHandler struct{}

func (UnimplementedKargoServiceHandler) GetVersionInfo(context.Context, *connect_go.Request[v1alpha1.GetVersionInfoRequest]) (*connect_go.Response[v1alpha1.GetVersionInfoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.GetVersionInfo is not implemented"))
}

func (UnimplementedKargoServiceHandler) GetPublicConfig(context.Context, *connect_go.Request[v1alpha1.GetPublicConfigRequest]) (*connect_go.Response[v1alpha1.GetPublicConfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.GetPublicConfig is not implemented"))
}

func (UnimplementedKargoServiceHandler) AdminLogin(context.Context, *connect_go.Request[v1alpha1.AdminLoginRequest]) (*connect_go.Response[v1alpha1.AdminLoginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.AdminLogin is not implemented"))
}

func (UnimplementedKargoServiceHandler) CreateStage(context.Context, *connect_go.Request[v1alpha1.CreateStageRequest]) (*connect_go.Response[v1alpha1.CreateStageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.CreateStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) ListStages(context.Context, *connect_go.Request[v1alpha1.ListStagesRequest]) (*connect_go.Response[v1alpha1.ListStagesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.ListStages is not implemented"))
}

func (UnimplementedKargoServiceHandler) GetStage(context.Context, *connect_go.Request[v1alpha1.GetStageRequest]) (*connect_go.Response[v1alpha1.GetStageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.GetStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) WatchStage(context.Context, *connect_go.Request[v1alpha1.WatchStageRequest], *connect_go.ServerStream[v1alpha1.WatchStageResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.WatchStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) UpdateStage(context.Context, *connect_go.Request[v1alpha1.UpdateStageRequest]) (*connect_go.Response[v1alpha1.UpdateStageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.UpdateStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) DeleteStage(context.Context, *connect_go.Request[v1alpha1.DeleteStageRequest]) (*connect_go.Response[v1alpha1.DeleteStageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.DeleteStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) PromoteStage(context.Context, *connect_go.Request[v1alpha1.PromoteStageRequest]) (*connect_go.Response[v1alpha1.PromoteStageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.PromoteStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) SetAutoPromotionForStage(context.Context, *connect_go.Request[v1alpha1.SetAutoPromotionForStageRequest]) (*connect_go.Response[v1alpha1.SetAutoPromotionForStageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.SetAutoPromotionForStage is not implemented"))
}

func (UnimplementedKargoServiceHandler) CreatePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.CreatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.CreatePromotionPolicyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.CreatePromotionPolicy is not implemented"))
}

func (UnimplementedKargoServiceHandler) ListPromotionPolicies(context.Context, *connect_go.Request[v1alpha1.ListPromotionPoliciesRequest]) (*connect_go.Response[v1alpha1.ListPromotionPoliciesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.ListPromotionPolicies is not implemented"))
}

func (UnimplementedKargoServiceHandler) GetPromotionPolicy(context.Context, *connect_go.Request[v1alpha1.GetPromotionPolicyRequest]) (*connect_go.Response[v1alpha1.GetPromotionPolicyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.GetPromotionPolicy is not implemented"))
}

func (UnimplementedKargoServiceHandler) UpdatePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.UpdatePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.UpdatePromotionPolicyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.UpdatePromotionPolicy is not implemented"))
}

func (UnimplementedKargoServiceHandler) DeletePromotionPolicy(context.Context, *connect_go.Request[v1alpha1.DeletePromotionPolicyRequest]) (*connect_go.Response[v1alpha1.DeletePromotionPolicyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.DeletePromotionPolicy is not implemented"))
}

func (UnimplementedKargoServiceHandler) CreateProject(context.Context, *connect_go.Request[v1alpha1.CreateProjectRequest]) (*connect_go.Response[v1alpha1.CreateProjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.CreateProject is not implemented"))
}

func (UnimplementedKargoServiceHandler) ListProjects(context.Context, *connect_go.Request[v1alpha1.ListProjectsRequest]) (*connect_go.Response[v1alpha1.ListProjectsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.ListProjects is not implemented"))
}

func (UnimplementedKargoServiceHandler) DeleteProject(context.Context, *connect_go.Request[v1alpha1.DeleteProjectRequest]) (*connect_go.Response[v1alpha1.DeleteProjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("akuity.io.kargo.service.v1alpha1.KargoService.DeleteProject is not implemented"))
}
