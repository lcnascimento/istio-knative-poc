package domain

import (
	pb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"
	exportsPb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
)

// AudienceProvider ...
type AudienceProvider string

var (
	// GoogleAudienceProvider ...
	GoogleAudienceProvider AudienceProvider = "google"
	// FacebookAudienceProvider ...
	FacebookAudienceProvider AudienceProvider = "facebook"
)

// Audience ...
type Audience struct {
	ID           string           `json:"id"`
	AppKey       string           `json:"app_key"`
	SegmentID    string           `json:"segment_id"`
	Name         string           `json:"name"`
	LastExportID string           `json:"last_export_id"`
	Provider     AudienceProvider `json:"provider"`
}

var dtoToDomainProvider = map[pb.AudienceProvider]AudienceProvider{
	pb.AudienceProvider_GOOGLE:   GoogleAudienceProvider,
	pb.AudienceProvider_FACEBOOK: FacebookAudienceProvider,
}

var domainToDTOProvider = map[AudienceProvider]pb.AudienceProvider{
	GoogleAudienceProvider:   pb.AudienceProvider_GOOGLE,
	FacebookAudienceProvider: pb.AudienceProvider_FACEBOOK,
}

// ToGRPCDTO ...
func (a Audience) ToGRPCDTO() *pb.Audience {
	return &pb.Audience{
		Id:           a.ID,
		AppKey:       a.AppKey,
		SegmentId:    a.SegmentID,
		LastExportId: a.LastExportID,
		Name:         a.Name,
		Provider:     domainToDTOProvider[a.Provider],
	}
}

// FillByGRPCDTO ...
func (a *Audience) FillByGRPCDTO(dto *pb.Audience) {
	a.ID = dto.Id
	a.AppKey = dto.AppKey
	a.SegmentID = dto.SegmentId
	a.LastExportID = dto.LastExportId
	a.Name = dto.Name
	a.Provider = dtoToDomainProvider[dto.Provider]
}

// Export ...
type Export struct {
	ID        string       `json:"id"`
	AppKey    string       `json:"app_key"`
	SegmentID string       `json:"segment_id"`
	Name      string       `json:"name"`
	Module    ExportModule `json:"module"`
}

// ExportModule ...
type ExportModule string

// AdsExportModule ...
var AdsExportModule ExportModule = "ads"

var dtoToDomainExportModule = map[exportsPb.ExportModule]ExportModule{
	exportsPb.ExportModule_ADS: AdsExportModule,
}

var domainToDTOExportModule = map[ExportModule]exportsPb.ExportModule{
	AdsExportModule: exportsPb.ExportModule_ADS,
}

// ToGRPCDTO ...
func (e Export) ToGRPCDTO() *exportsPb.Export {
	return &exportsPb.Export{
		Id:        e.ID,
		AppKey:    e.AppKey,
		SegmentId: e.SegmentID,
		Name:      e.Name,
		Module:    domainToDTOExportModule[e.Module],
	}
}

// FillByGRPCDTO ...
func (e *Export) FillByGRPCDTO(dto *exportsPb.Export) {
	e.ID = dto.Id
	e.AppKey = dto.AppKey
	e.SegmentID = dto.SegmentId
	e.Name = dto.Name
	e.Module = dtoToDomainExportModule[dto.Module]
}
