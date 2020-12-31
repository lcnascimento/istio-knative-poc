package domain

import (
	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
	segmentsPb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

// Export ...
type Export struct {
	ID        string `json:"id"`
	AppKey    string `json:"app_key"`
	SegmentID string `json:"segment_id"`
	Name      string `json:"json"`
	Module    string `json:"module"`
}

// ToGRPCDTO ...
func (e Export) ToGRPCDTO() *pb.Export {
	return &pb.Export{
		Id:        e.ID,
		AppKey:    e.AppKey,
		SegmentId: e.SegmentID,
		Name:      e.Name,
		Module:    e.Module,
	}
}

// FillByGRPCDTO ...
func (e *Export) FillByGRPCDTO(dto *pb.Export) {
	e.ID = dto.Id
	e.AppKey = dto.AppKey
	e.SegmentID = dto.SegmentId
	e.Name = dto.Name
	e.Module = dto.Module
}

// User ...
type User struct {
	Reference string `json:"reference"`
	AppKey    string `json:"app_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

// ToGRPCDTO ...
func (u User) ToGRPCDTO() *segmentsPb.User {
	return &segmentsPb.User{
		Reference: u.Reference,
		AppKey:    u.AppKey,
		Name:      u.Name,
		Email:     u.Email,
	}
}

// FillByGRPCDTO ...
func (u *User) FillByGRPCDTO(dto *segmentsPb.User) {
	u.Reference = dto.Reference
	u.AppKey = dto.AppKey
	u.Name = dto.Name
	u.Email = dto.Email
}

// Segment ...
type Segment struct {
	ID          string      `json:"id"`
	AppKey      string      `json:"app_key"`
	Name        string      `json:"json"`
	Description string      `json:"description"`
	Rules       interface{} `json:"rules"`
	Version     int         `json:"version"`
}

// ToGRPCDTO ...
func (s Segment) ToGRPCDTO() *segmentsPb.Segment {
	return &segmentsPb.Segment{
		Id:          s.ID,
		AppKey:      s.AppKey,
		Name:        s.Name,
		Description: s.Description,
		Version:     int32(s.Version),
	}
}

// FillByGRPCDTO ...
func (s *Segment) FillByGRPCDTO(dto *segmentsPb.Segment) {
	s.ID = dto.Id
	s.AppKey = dto.AppKey
	s.Name = dto.Name
	s.Description = dto.Description
	s.Version = int(dto.Version)
}
