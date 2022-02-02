package service

import (
	"context"
	"github.com/OpenCal-FYDP/AsyncCalendarOptimizer/rpc"
)

type CalOptimizerService struct{}

func (c *CalOptimizerService) OptimizeCal(ctx context.Context, req *rpc.EventReq) (*rpc.EventRes, error) {
	panic("implement me")
}

func New() *CalOptimizerService {
	return &CalOptimizerService{}
}
