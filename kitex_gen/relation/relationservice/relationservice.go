// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	relation "github.com/simple/douyin/kitex_gen/relation"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"RelationAction": kitex.NewMethodInfo(relationActionHandler, newRelationActionArgs, newRelationActionResult, false),
		"FollowList":     kitex.NewMethodInfo(followListHandler, newFollowListArgs, newFollowListResult, false),
		"FollowerList":   kitex.NewMethodInfo(followerListHandler, newFollowerListArgs, newFollowerListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.RelationActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).RelationAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RelationActionArgs:
		success, err := handler.(relation.RelationService).RelationAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RelationActionResult)
		realResult.Success = success
	}
	return nil
}
func newRelationActionArgs() interface{} {
	return &RelationActionArgs{}
}

func newRelationActionResult() interface{} {
	return &RelationActionResult{}
}

type RelationActionArgs struct {
	Req *relation.RelationActionRequest
}

func (p *RelationActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.RelationActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RelationActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RelationActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RelationActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RelationActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RelationActionArgs) Unmarshal(in []byte) error {
	msg := new(relation.RelationActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RelationActionArgs_Req_DEFAULT *relation.RelationActionRequest

func (p *RelationActionArgs) GetReq() *relation.RelationActionRequest {
	if !p.IsSetReq() {
		return RelationActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RelationActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type RelationActionResult struct {
	Success *relation.RelationActionResponse
}

var RelationActionResult_Success_DEFAULT *relation.RelationActionResponse

func (p *RelationActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.RelationActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RelationActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RelationActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RelationActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RelationActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RelationActionResult) Unmarshal(in []byte) error {
	msg := new(relation.RelationActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RelationActionResult) GetSuccess() *relation.RelationActionResponse {
	if !p.IsSetSuccess() {
		return RelationActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RelationActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.RelationActionResponse)
}

func (p *RelationActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.FollowListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).FollowList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowListArgs:
		success, err := handler.(relation.RelationService).FollowList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowListArgs() interface{} {
	return &FollowListArgs{}
}

func newFollowListResult() interface{} {
	return &FollowListResult{}
}

type FollowListArgs struct {
	Req *relation.FollowListRequest
}

func (p *FollowListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.FollowListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowListArgs) Unmarshal(in []byte) error {
	msg := new(relation.FollowListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowListArgs_Req_DEFAULT *relation.FollowListRequest

func (p *FollowListArgs) GetReq() *relation.FollowListRequest {
	if !p.IsSetReq() {
		return FollowListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowListResult struct {
	Success *relation.FollowListResponse
}

var FollowListResult_Success_DEFAULT *relation.FollowListResponse

func (p *FollowListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.FollowListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowListResult) Unmarshal(in []byte) error {
	msg := new(relation.FollowListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowListResult) GetSuccess() *relation.FollowListResponse {
	if !p.IsSetSuccess() {
		return FollowListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowListResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.FollowListResponse)
}

func (p *FollowListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.FollowerListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).FollowerList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowerListArgs:
		success, err := handler.(relation.RelationService).FollowerList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowerListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowerListArgs() interface{} {
	return &FollowerListArgs{}
}

func newFollowerListResult() interface{} {
	return &FollowerListResult{}
}

type FollowerListArgs struct {
	Req *relation.FollowerListRequest
}

func (p *FollowerListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.FollowerListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowerListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowerListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowerListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowerListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowerListArgs) Unmarshal(in []byte) error {
	msg := new(relation.FollowerListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowerListArgs_Req_DEFAULT *relation.FollowerListRequest

func (p *FollowerListArgs) GetReq() *relation.FollowerListRequest {
	if !p.IsSetReq() {
		return FollowerListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowerListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowerListResult struct {
	Success *relation.FollowerListResponse
}

var FollowerListResult_Success_DEFAULT *relation.FollowerListResponse

func (p *FollowerListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.FollowerListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowerListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowerListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowerListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowerListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowerListResult) Unmarshal(in []byte) error {
	msg := new(relation.FollowerListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowerListResult) GetSuccess() *relation.FollowerListResponse {
	if !p.IsSetSuccess() {
		return FollowerListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowerListResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.FollowerListResponse)
}

func (p *FollowerListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) RelationAction(ctx context.Context, Req *relation.RelationActionRequest) (r *relation.RelationActionResponse, err error) {
	var _args RelationActionArgs
	_args.Req = Req
	var _result RelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, Req *relation.FollowListRequest) (r *relation.FollowListResponse, err error) {
	var _args FollowListArgs
	_args.Req = Req
	var _result FollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, Req *relation.FollowerListRequest) (r *relation.FollowerListResponse, err error) {
	var _args FollowerListArgs
	_args.Req = Req
	var _result FollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}