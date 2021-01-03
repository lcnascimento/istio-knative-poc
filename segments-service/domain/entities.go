package domain

import pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"

// User ...
type User struct {
	Reference string `json:"reference"`
	AppKey    string `json:"app_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

// ToGRPCDTO ...
func (u User) ToGRPCDTO() *pb.User {
	return &pb.User{
		Reference: u.Reference,
		AppKey:    u.AppKey,
		Name:      u.Name,
		Email:     u.Email,
	}
}

// FillByGRPCDTO ...
func (u *User) FillByGRPCDTO(dto *pb.User) {
	u.Reference = dto.Reference
	u.AppKey = dto.AppKey
	u.Name = dto.Name
	u.Email = dto.Email
}

// Segment ...
type Segment struct {
	ID          string      `json:"id"`
	AppKey      string      `json:"app_key"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Rules       interface{} `json:"rules"`
	Version     int         `json:"version"`
}

// ToGRPCDTO ...
func (s Segment) ToGRPCDTO() *pb.Segment {
	return &pb.Segment{
		Id:          s.ID,
		AppKey:      s.AppKey,
		Name:        s.Name,
		Description: s.Description,
		Version:     int32(s.Version),
	}
}

// FillByGRPCDTO ...
func (s *Segment) FillByGRPCDTO(dto *pb.Segment) {
	s.ID = dto.Id
	s.AppKey = dto.AppKey
	s.Name = dto.Name
	s.Description = dto.Description
	s.Version = int(dto.Version)
}
