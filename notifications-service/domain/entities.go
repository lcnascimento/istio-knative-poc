package domain

import (
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	segments_pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

var dtoToDomainChannelMap = map[pb.NotificationChannel]Channel{
	pb.NotificationChannel_EMAIL:   EmailChannel,
	pb.NotificationChannel_SMS:     SMSChannel,
	pb.NotificationChannel_BROWSER: BrowserPushChannel,
}

var domainToDTOChannelMap = map[Channel]pb.NotificationChannel{
	EmailChannel:       pb.NotificationChannel_EMAIL,
	SMSChannel:         pb.NotificationChannel_SMS,
	BrowserPushChannel: pb.NotificationChannel_BROWSER,
}

// Notification ...
type Notification struct {
	ID        string  `json:"id"`
	AppKey    string  `json:"app_key"`
	Name      string  `json:"name"`
	Channel   Channel `json:"channel"`
	SegmentID string  `json:"segment_id"`
}

// ToGRPCDTO ...
func (n Notification) ToGRPCDTO() *pb.Notification {
	return &pb.Notification{
		Id:        n.ID,
		AppKey:    n.AppKey,
		Name:      n.Name,
		Channel:   domainToDTOChannelMap[n.Channel],
		SegmentId: n.SegmentID,
	}
}

// FillByGRPCDTO ...
func (n *Notification) FillByGRPCDTO(dto *pb.Notification) {
	n.ID = dto.Id
	n.AppKey = dto.AppKey
	n.Name = dto.Name
	n.Channel = dtoToDomainChannelMap[dto.Channel]
	n.SegmentID = dto.SegmentId
}

// User ...
type User struct {
	Reference string `json:"reference"`
	AppKey    string `json:"app_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

// ToGRPCDTO ...
func (u User) ToGRPCDTO() *segments_pb.User {
	return &segments_pb.User{
		Reference: u.Reference,
		AppKey:    u.AppKey,
		Name:      u.Name,
		Email:     u.Email,
	}
}

// FillByGRPCDTO ...
func (u *User) FillByGRPCDTO(dto *segments_pb.User) {
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
func (s Segment) ToGRPCDTO() *segments_pb.Segment {
	return &segments_pb.Segment{
		Id:          s.ID,
		AppKey:      s.AppKey,
		Name:        s.Name,
		Description: s.Description,
		Version:     int32(s.Version),
	}
}

// FillByGRPCDTO ...
func (s *Segment) FillByGRPCDTO(dto *segments_pb.Segment) {
	s.ID = dto.Id
	s.AppKey = dto.AppKey
	s.Name = dto.Name
	s.Description = dto.Description
	s.Version = int(dto.Version)
}
