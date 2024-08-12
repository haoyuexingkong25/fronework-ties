package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterGrpc(port int, handle func(s *grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen----: %v", err)
	}
	s := grpc.NewServer()
	handle(s)

	log.Printf("server listening at ----:%v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve----: %v", err)
	}
}
