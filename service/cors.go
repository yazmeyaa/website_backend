package service

import (
	"context"
	"yazmeyaa_projects/cors"
)

type CorsService interface {
	AddOrigin(ctx context.Context, origin string, methods []string, headers []string) (*cors.CORSRecord, error)
	DisableOrigin(ctx context.Context, origin string) error
	EnableOrigin(ctx context.Context, origin string) error
	ModifyOrigin(ctx context.Context, origin string, methods []string, headers []string) error
	RemoveOrigin(ctx context.Context, origin string) error
	GetAllRecords(ctx context.Context) ([]cors.CORSRecord, error)
	GetRecord(ctx context.Context, origin string) (cors.CORSRecord, error)
}
