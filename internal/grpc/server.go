package grpc

import (
	"context"
	"fmt"

	pb "github.com/netweng/go-simple-server/proto"
)

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
	pb.UnimplementedGoServerServer
}

func (b Backend) GenAuditLog(ctx context.Context, req *pb.GenAuditLogRequest) (*pb.AuditLog, error) {
	return &pb.AuditLog{
		En: fmt.Sprint(req.A + req.B),
		Zh: "动物园",
	}, nil
}

func (b Backend) GetCat(ctx context.Context, req *pb.GetCatRequest) (*pb.Cat, error) {
	return &pb.Cat{
		Name: "kitty",
		Age:  "11",
	}, nil
}

func (b Backend) RegisterFuncs(ctx context.Context, req *pb.RegisterFuncsRequest) (*pb.Funcs, error) {
	funcs := []string{
		"GenAuditLog",
		"GetCat",
	}
	return &pb.Funcs{
		Funcs: funcs,
	}, nil
}
