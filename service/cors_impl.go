package service

import (
	"context"
	"fmt"
	"yazmeyaa_projects/cors"

	"github.com/redis/go-redis/v9"
)

type corsServiceImpl struct {
	rc *redis.Client
}

const (
	origins_set = "origins_set"
)

// AddOrigin implements CorsService.
func (c *corsServiceImpl) AddOrigin(ctx context.Context, origin string, methods []string, headers []string) (*cors.CORSRecord, error) {
	record := cors.CORSRecord{
		Origin:         origin,
		OriginAllowed:  true,
		AllowedMethods: methods,
		AllowedHeaders: headers,
	}

	_, err := c.rc.Pipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.Set(ctx, fmt.Sprintf("origin:%s", origin), cors.ORIGIN_ALLOWED, 0).Err(); err != nil {
			return err
		}

		if err := p.SAdd(ctx, origins_set, origin).Err(); err != nil {
			return err
		}

		if err := p.Del(ctx, fmt.Sprintf("headers:%s", origin), fmt.Sprintf("methods:%s", origin)).Err(); err != nil {
			return err
		}

		if err := p.RPush(ctx, fmt.Sprintf("headers:%s", origin), headers).Err(); err != nil {
			return err
		}
		if err := p.RPush(ctx, fmt.Sprintf("methods:%s", origin), methods).Err(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &record, nil
}

// DisableOrigin implements CorsService.
func (c *corsServiceImpl) DisableOrigin(ctx context.Context, origin string) error {
	if err := c.rc.Get(ctx, fmt.Sprintf("origin:%s", origin)).Err(); err != nil {
		return cors.ErrNotFoundRecord
	}

	if err := c.rc.Set(ctx, fmt.Sprintf("origin:%s", origin), cors.ORIGIN_NOT_ALLOWED, 0).Err(); err != nil {
		return nil
	}

	return nil
}

// EnableOrigin implements CorsService.
func (c *corsServiceImpl) EnableOrigin(ctx context.Context, origin string) error {
	if err := c.rc.Get(ctx, fmt.Sprintf("origin:%s", origin)).Err(); err != nil {
		return cors.ErrNotFoundRecord
	}

	if err := c.rc.Set(ctx, fmt.Sprintf("origin:%s", origin), cors.ORIGIN_ALLOWED, 0).Err(); err != nil {
		return nil
	}

	return nil
}

// GetAllRecords implements CorsService.
func (c *corsServiceImpl) GetAllRecords(ctx context.Context) ([]cors.CORSRecord, error) {
	var records []cors.CORSRecord

	origins, err := c.rc.SMembers(ctx, origins_set).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	for _, origin := range origins {
		record, err := c.GetRecord(ctx, origin)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// GetRecord implements CorsService.
func (c *corsServiceImpl) GetRecord(ctx context.Context, origin string) (cors.CORSRecord, error) {
	var record cors.CORSRecord

	results, err := c.rc.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Get(ctx, fmt.Sprintf("origin:%s", origin))
		pipe.LRange(ctx, fmt.Sprintf("methods:%s", origin), 0, -1)
		pipe.LRange(ctx, fmt.Sprintf("headers:%s", origin), 0, -1)
		return nil
	})
	if err != nil && err != redis.Nil {
		return record, err
	}

	originAllowed, err := results[0].(*redis.StringCmd).Result()
	if err != nil && err != redis.Nil {
		return record, err
	}
	record.OriginAllowed = originAllowed == cors.ORIGIN_ALLOWED

	methods, err := results[1].(*redis.StringSliceCmd).Result()
	if err != nil && err != redis.Nil {
		return record, err
	}
	record.AllowedMethods = methods

	headers, err := results[2].(*redis.StringSliceCmd).Result()
	if err != nil && err != redis.Nil {
		return record, err
	}
	record.AllowedHeaders = headers
	record.Origin = origin

	return record, nil
}

// ModifyOrigin implements CorsService.
func (c *corsServiceImpl) ModifyOrigin(ctx context.Context, origin string, methods []string, headers []string) error {
	exists, err := c.rc.Exists(ctx, fmt.Sprintf("origin:%s", origin)).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return cors.ErrNotFoundRecord
	}

	_, err = c.rc.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, fmt.Sprintf("headers:%s", origin), fmt.Sprintf("methods:%s", origin))

		if len(headers) > 0 {
			p.RPush(ctx, fmt.Sprintf("headers:%s", origin), headers)
		}
		if len(methods) > 0 {
			p.RPush(ctx, fmt.Sprintf("methods:%s", origin), methods)
		}

		return nil
	})

	return err
}

// RemoveOrigin implements CorsService.
func (c *corsServiceImpl) RemoveOrigin(ctx context.Context, origin string) error {
	exists, err := c.rc.Exists(ctx, fmt.Sprintf("origin:%s", origin)).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return cors.ErrNotFoundRecord
	}

	_, err = c.rc.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(
			ctx,
			fmt.Sprintf("origin:%s", origin),
			fmt.Sprintf("headers:%s", origin),
			fmt.Sprintf("methods:%s", origin),
		)
		p.SRem(ctx, origins_set, origin)
		return nil
	})

	return err
}

func NewCorsService(rc *redis.Client) CorsService {
	return &corsServiceImpl{
		rc: rc,
	}
}
