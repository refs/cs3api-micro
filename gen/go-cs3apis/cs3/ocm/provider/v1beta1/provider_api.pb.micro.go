// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cs3/ocm/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	fmt "fmt"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/rpc/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProviderAPI service

type ProviderAPIService interface {
	// Check if a given system provider is registered in the mesh or not.
	// MUST return CODE_UNAUTHENTICATED if the system is not registered
	IsProviderAllowed(ctx context.Context, in *IsProviderAllowedRequest, opts ...client.CallOption) (*IsProviderAllowedResponse, error)
	// Get the information of the provider identified by a specific domain.
	// MUST return CODE_NOT_FOUND if the sync'n'share system provider does not exist.
	GetInfoByDomain(ctx context.Context, in *GetInfoByDomainRequest, opts ...client.CallOption) (*GetInfoByDomainResponse, error)
	// Get the information of all the providers registered in the mesh.
	ListAllProviders(ctx context.Context, in *ListAllProvidersRequest, opts ...client.CallOption) (*ListAllProvidersResponse, error)
}

type providerAPIService struct {
	c    client.Client
	name string
}

func NewProviderAPIService(name string, c client.Client) ProviderAPIService {
	return &providerAPIService{
		c:    c,
		name: name,
	}
}

func (c *providerAPIService) IsProviderAllowed(ctx context.Context, in *IsProviderAllowedRequest, opts ...client.CallOption) (*IsProviderAllowedResponse, error) {
	req := c.c.NewRequest(c.name, "ProviderAPI.IsProviderAllowed", in)
	out := new(IsProviderAllowedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerAPIService) GetInfoByDomain(ctx context.Context, in *GetInfoByDomainRequest, opts ...client.CallOption) (*GetInfoByDomainResponse, error) {
	req := c.c.NewRequest(c.name, "ProviderAPI.GetInfoByDomain", in)
	out := new(GetInfoByDomainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerAPIService) ListAllProviders(ctx context.Context, in *ListAllProvidersRequest, opts ...client.CallOption) (*ListAllProvidersResponse, error) {
	req := c.c.NewRequest(c.name, "ProviderAPI.ListAllProviders", in)
	out := new(ListAllProvidersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProviderAPI service

type ProviderAPIHandler interface {
	// Check if a given system provider is registered in the mesh or not.
	// MUST return CODE_UNAUTHENTICATED if the system is not registered
	IsProviderAllowed(context.Context, *IsProviderAllowedRequest, *IsProviderAllowedResponse) error
	// Get the information of the provider identified by a specific domain.
	// MUST return CODE_NOT_FOUND if the sync'n'share system provider does not exist.
	GetInfoByDomain(context.Context, *GetInfoByDomainRequest, *GetInfoByDomainResponse) error
	// Get the information of all the providers registered in the mesh.
	ListAllProviders(context.Context, *ListAllProvidersRequest, *ListAllProvidersResponse) error
}

func RegisterProviderAPIHandler(s server.Server, hdlr ProviderAPIHandler, opts ...server.HandlerOption) error {
	type providerAPI interface {
		IsProviderAllowed(ctx context.Context, in *IsProviderAllowedRequest, out *IsProviderAllowedResponse) error
		GetInfoByDomain(ctx context.Context, in *GetInfoByDomainRequest, out *GetInfoByDomainResponse) error
		ListAllProviders(ctx context.Context, in *ListAllProvidersRequest, out *ListAllProvidersResponse) error
	}
	type ProviderAPI struct {
		providerAPI
	}
	h := &providerAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&ProviderAPI{h}, opts...))
}

type providerAPIHandler struct {
	ProviderAPIHandler
}

func (h *providerAPIHandler) IsProviderAllowed(ctx context.Context, in *IsProviderAllowedRequest, out *IsProviderAllowedResponse) error {
	return h.ProviderAPIHandler.IsProviderAllowed(ctx, in, out)
}

func (h *providerAPIHandler) GetInfoByDomain(ctx context.Context, in *GetInfoByDomainRequest, out *GetInfoByDomainResponse) error {
	return h.ProviderAPIHandler.GetInfoByDomain(ctx, in, out)
}

func (h *providerAPIHandler) ListAllProviders(ctx context.Context, in *ListAllProvidersRequest, out *ListAllProvidersResponse) error {
	return h.ProviderAPIHandler.ListAllProviders(ctx, in, out)
}
