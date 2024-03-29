// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sso.proto

package go_micro_srv_sso

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Sso service

type SsoService interface {
	Token(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	CurrentUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*Userinfo, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	SendEmail(ctx context.Context, in *EmailRequest, opts ...client.CallOption) (*EmailResponse, error)
	ForgetPassword(ctx context.Context, in *PasswordRequest, opts ...client.CallOption) (*PasswordResponse, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Sso_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Sso_PingPongService, error)
}

type ssoService struct {
	c    client.Client
	name string
}

func NewSsoService(name string, c client.Client) SsoService {
	return &ssoService{
		c:    c,
		name: name,
	}
}

func (c *ssoService) Token(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Sso.Token", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoService) CurrentUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*Userinfo, error) {
	req := c.c.NewRequest(c.name, "Sso.CurrentUser", in)
	out := new(Userinfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Sso.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoService) SendEmail(ctx context.Context, in *EmailRequest, opts ...client.CallOption) (*EmailResponse, error) {
	req := c.c.NewRequest(c.name, "Sso.SendEmail", in)
	out := new(EmailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoService) ForgetPassword(ctx context.Context, in *PasswordRequest, opts ...client.CallOption) (*PasswordResponse, error) {
	req := c.c.NewRequest(c.name, "Sso.ForgetPassword", in)
	out := new(PasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Sso_StreamService, error) {
	req := c.c.NewRequest(c.name, "Sso.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &ssoServiceStream{stream}, nil
}

type Sso_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type ssoServiceStream struct {
	stream client.Stream
}

func (x *ssoServiceStream) Close() error {
	return x.stream.Close()
}

func (x *ssoServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *ssoServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *ssoServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *ssoServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ssoService) PingPong(ctx context.Context, opts ...client.CallOption) (Sso_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Sso.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &ssoServicePingPong{stream}, nil
}

type Sso_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type ssoServicePingPong struct {
	stream client.Stream
}

func (x *ssoServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *ssoServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *ssoServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *ssoServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *ssoServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *ssoServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Sso service

type SsoHandler interface {
	Token(context.Context, *AuthRequest, *AuthResponse) error
	CurrentUser(context.Context, *UserRequest, *Userinfo) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	SendEmail(context.Context, *EmailRequest, *EmailResponse) error
	ForgetPassword(context.Context, *PasswordRequest, *PasswordResponse) error
	Stream(context.Context, *StreamingRequest, Sso_StreamStream) error
	PingPong(context.Context, Sso_PingPongStream) error
}

func RegisterSsoHandler(s server.Server, hdlr SsoHandler, opts ...server.HandlerOption) error {
	type sso interface {
		Token(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		CurrentUser(ctx context.Context, in *UserRequest, out *Userinfo) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		SendEmail(ctx context.Context, in *EmailRequest, out *EmailResponse) error
		ForgetPassword(ctx context.Context, in *PasswordRequest, out *PasswordResponse) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Sso struct {
		sso
	}
	h := &ssoHandler{hdlr}
	return s.Handle(s.NewHandler(&Sso{h}, opts...))
}

type ssoHandler struct {
	SsoHandler
}

func (h *ssoHandler) Token(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.SsoHandler.Token(ctx, in, out)
}

func (h *ssoHandler) CurrentUser(ctx context.Context, in *UserRequest, out *Userinfo) error {
	return h.SsoHandler.CurrentUser(ctx, in, out)
}

func (h *ssoHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.SsoHandler.Register(ctx, in, out)
}

func (h *ssoHandler) SendEmail(ctx context.Context, in *EmailRequest, out *EmailResponse) error {
	return h.SsoHandler.SendEmail(ctx, in, out)
}

func (h *ssoHandler) ForgetPassword(ctx context.Context, in *PasswordRequest, out *PasswordResponse) error {
	return h.SsoHandler.ForgetPassword(ctx, in, out)
}

func (h *ssoHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SsoHandler.Stream(ctx, m, &ssoStreamStream{stream})
}

type Sso_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type ssoStreamStream struct {
	stream server.Stream
}

func (x *ssoStreamStream) Close() error {
	return x.stream.Close()
}

func (x *ssoStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *ssoStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *ssoStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *ssoStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *ssoHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.SsoHandler.PingPong(ctx, &ssoPingPongStream{stream})
}

type Sso_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type ssoPingPongStream struct {
	stream server.Stream
}

func (x *ssoPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *ssoPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *ssoPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *ssoPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *ssoPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *ssoPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
