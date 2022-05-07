package constant

import "time"

const (
	ServiceName           = "project-info-manager.npool.top" //nolint
	GrpcTimeout           = time.Second * 10
	DefaultPageSize int32 = 10
)
