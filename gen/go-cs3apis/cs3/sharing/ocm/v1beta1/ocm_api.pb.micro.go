// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cs3/sharing/ocm/v1beta1/ocm_api.proto

package ocmv1beta1

import (
	fmt "fmt"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/identity/user/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/ocm/provider/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/rpc/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/storage/provider/v1beta1"
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

// Client API for OcmAPI service

type OcmAPIService interface {
	// Creates a new ocm share.
	// MUST return CODE_NOT_FOUND if the resource reference does not exist.
	// MUST return CODE_ALREADY_EXISTS if the share already exists for the 4-tuple consisting of
	// (owner, shared_resource, grantee).
	// New shares MUST be created in the state SHARE_STATE_PENDING.
	CreateOCMShare(ctx context.Context, in *CreateOCMShareRequest, opts ...client.CallOption) (*CreateOCMShareResponse, error)
	// Removes a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	RemoveOCMShare(ctx context.Context, in *RemoveOCMShareRequest, opts ...client.CallOption) (*RemoveOCMShareResponse, error)
	// Gets share information for a single share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	GetOCMShare(ctx context.Context, in *GetOCMShareRequest, opts ...client.CallOption) (*GetOCMShareResponse, error)
	// List the shares the authenticated principal has created,
	// both as owner and creator. If a filter is specified, only
	// shares satisfying the filter MUST be returned.
	ListOCMShares(ctx context.Context, in *ListOCMSharesRequest, opts ...client.CallOption) (*ListOCMSharesResponse, error)
	// Updates a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateOCMShare(ctx context.Context, in *UpdateOCMShareRequest, opts ...client.CallOption) (*UpdateOCMShareResponse, error)
	// List all shares the authenticated principal has received.
	ListReceivedOCMShares(ctx context.Context, in *ListReceivedOCMSharesRequest, opts ...client.CallOption) (*ListReceivedOCMSharesResponse, error)
	// Update the received share to change the share state or the display name.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateReceivedOCMShare(ctx context.Context, in *UpdateReceivedOCMShareRequest, opts ...client.CallOption) (*UpdateReceivedOCMShareResponse, error)
	// Get the information for the given received share reference.
	// MUST return CODE_NOT_FOUND if the received share reference does not exist.
	GetReceivedOCMShare(ctx context.Context, in *GetReceivedOCMShareRequest, opts ...client.CallOption) (*GetReceivedOCMShareResponse, error)
}

type ocmAPIService struct {
	c    client.Client
	name string
}

func NewOcmAPIService(name string, c client.Client) OcmAPIService {
	return &ocmAPIService{
		c:    c,
		name: name,
	}
}

func (c *ocmAPIService) CreateOCMShare(ctx context.Context, in *CreateOCMShareRequest, opts ...client.CallOption) (*CreateOCMShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.CreateOCMShare", in)
	out := new(CreateOCMShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) RemoveOCMShare(ctx context.Context, in *RemoveOCMShareRequest, opts ...client.CallOption) (*RemoveOCMShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.RemoveOCMShare", in)
	out := new(RemoveOCMShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) GetOCMShare(ctx context.Context, in *GetOCMShareRequest, opts ...client.CallOption) (*GetOCMShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.GetOCMShare", in)
	out := new(GetOCMShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) ListOCMShares(ctx context.Context, in *ListOCMSharesRequest, opts ...client.CallOption) (*ListOCMSharesResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.ListOCMShares", in)
	out := new(ListOCMSharesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) UpdateOCMShare(ctx context.Context, in *UpdateOCMShareRequest, opts ...client.CallOption) (*UpdateOCMShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.UpdateOCMShare", in)
	out := new(UpdateOCMShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) ListReceivedOCMShares(ctx context.Context, in *ListReceivedOCMSharesRequest, opts ...client.CallOption) (*ListReceivedOCMSharesResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.ListReceivedOCMShares", in)
	out := new(ListReceivedOCMSharesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) UpdateReceivedOCMShare(ctx context.Context, in *UpdateReceivedOCMShareRequest, opts ...client.CallOption) (*UpdateReceivedOCMShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.UpdateReceivedOCMShare", in)
	out := new(UpdateReceivedOCMShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocmAPIService) GetReceivedOCMShare(ctx context.Context, in *GetReceivedOCMShareRequest, opts ...client.CallOption) (*GetReceivedOCMShareResponse, error) {
	req := c.c.NewRequest(c.name, "OcmAPI.GetReceivedOCMShare", in)
	out := new(GetReceivedOCMShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OcmAPI service

type OcmAPIHandler interface {
	// Creates a new ocm share.
	// MUST return CODE_NOT_FOUND if the resource reference does not exist.
	// MUST return CODE_ALREADY_EXISTS if the share already exists for the 4-tuple consisting of
	// (owner, shared_resource, grantee).
	// New shares MUST be created in the state SHARE_STATE_PENDING.
	CreateOCMShare(context.Context, *CreateOCMShareRequest, *CreateOCMShareResponse) error
	// Removes a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	RemoveOCMShare(context.Context, *RemoveOCMShareRequest, *RemoveOCMShareResponse) error
	// Gets share information for a single share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	GetOCMShare(context.Context, *GetOCMShareRequest, *GetOCMShareResponse) error
	// List the shares the authenticated principal has created,
	// both as owner and creator. If a filter is specified, only
	// shares satisfying the filter MUST be returned.
	ListOCMShares(context.Context, *ListOCMSharesRequest, *ListOCMSharesResponse) error
	// Updates a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateOCMShare(context.Context, *UpdateOCMShareRequest, *UpdateOCMShareResponse) error
	// List all shares the authenticated principal has received.
	ListReceivedOCMShares(context.Context, *ListReceivedOCMSharesRequest, *ListReceivedOCMSharesResponse) error
	// Update the received share to change the share state or the display name.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateReceivedOCMShare(context.Context, *UpdateReceivedOCMShareRequest, *UpdateReceivedOCMShareResponse) error
	// Get the information for the given received share reference.
	// MUST return CODE_NOT_FOUND if the received share reference does not exist.
	GetReceivedOCMShare(context.Context, *GetReceivedOCMShareRequest, *GetReceivedOCMShareResponse) error
}

func RegisterOcmAPIHandler(s server.Server, hdlr OcmAPIHandler, opts ...server.HandlerOption) error {
	type ocmAPI interface {
		CreateOCMShare(ctx context.Context, in *CreateOCMShareRequest, out *CreateOCMShareResponse) error
		RemoveOCMShare(ctx context.Context, in *RemoveOCMShareRequest, out *RemoveOCMShareResponse) error
		GetOCMShare(ctx context.Context, in *GetOCMShareRequest, out *GetOCMShareResponse) error
		ListOCMShares(ctx context.Context, in *ListOCMSharesRequest, out *ListOCMSharesResponse) error
		UpdateOCMShare(ctx context.Context, in *UpdateOCMShareRequest, out *UpdateOCMShareResponse) error
		ListReceivedOCMShares(ctx context.Context, in *ListReceivedOCMSharesRequest, out *ListReceivedOCMSharesResponse) error
		UpdateReceivedOCMShare(ctx context.Context, in *UpdateReceivedOCMShareRequest, out *UpdateReceivedOCMShareResponse) error
		GetReceivedOCMShare(ctx context.Context, in *GetReceivedOCMShareRequest, out *GetReceivedOCMShareResponse) error
	}
	type OcmAPI struct {
		ocmAPI
	}
	h := &ocmAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&OcmAPI{h}, opts...))
}

type ocmAPIHandler struct {
	OcmAPIHandler
}

func (h *ocmAPIHandler) CreateOCMShare(ctx context.Context, in *CreateOCMShareRequest, out *CreateOCMShareResponse) error {
	return h.OcmAPIHandler.CreateOCMShare(ctx, in, out)
}

func (h *ocmAPIHandler) RemoveOCMShare(ctx context.Context, in *RemoveOCMShareRequest, out *RemoveOCMShareResponse) error {
	return h.OcmAPIHandler.RemoveOCMShare(ctx, in, out)
}

func (h *ocmAPIHandler) GetOCMShare(ctx context.Context, in *GetOCMShareRequest, out *GetOCMShareResponse) error {
	return h.OcmAPIHandler.GetOCMShare(ctx, in, out)
}

func (h *ocmAPIHandler) ListOCMShares(ctx context.Context, in *ListOCMSharesRequest, out *ListOCMSharesResponse) error {
	return h.OcmAPIHandler.ListOCMShares(ctx, in, out)
}

func (h *ocmAPIHandler) UpdateOCMShare(ctx context.Context, in *UpdateOCMShareRequest, out *UpdateOCMShareResponse) error {
	return h.OcmAPIHandler.UpdateOCMShare(ctx, in, out)
}

func (h *ocmAPIHandler) ListReceivedOCMShares(ctx context.Context, in *ListReceivedOCMSharesRequest, out *ListReceivedOCMSharesResponse) error {
	return h.OcmAPIHandler.ListReceivedOCMShares(ctx, in, out)
}

func (h *ocmAPIHandler) UpdateReceivedOCMShare(ctx context.Context, in *UpdateReceivedOCMShareRequest, out *UpdateReceivedOCMShareResponse) error {
	return h.OcmAPIHandler.UpdateReceivedOCMShare(ctx, in, out)
}

func (h *ocmAPIHandler) GetReceivedOCMShare(ctx context.Context, in *GetReceivedOCMShareRequest, out *GetReceivedOCMShareResponse) error {
	return h.OcmAPIHandler.GetReceivedOCMShare(ctx, in, out)
}
