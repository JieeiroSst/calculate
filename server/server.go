package main

import (
	"context"
	"github.com/JIeeiroSst/go-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	proto.UnimplementedServiceServer
}

func (s *Server) Add(ctx context.Context,req *proto.Request)(*proto.Response,error){
	a,b:=req.GetA(),req.GetB()
	result:=a+b
	return &proto.Response{Result:result}, nil
}

func (s *Server) Multiply(ctx context.Context,req *proto.Request)(*proto.Response,error){
	a,b:=req.GetA(),req.GetB()
	result:=a*b
	return &proto.Response{Result:result}, nil
}

func (s *Server) Minus(ctx context.Context,req *proto.Request)(*proto.Response,error){
	a,b:=req.GetA(),req.GetB()
	result:=a-b
	return &proto.Response{Result:result}, nil
}

func (s *Server) Division(ctx context.Context,req *proto.Request)(*proto.Response,error){
	a,b:=req.GetA(),req.GetB()
	result:=a/b
	return &proto.Response{Result:result}, nil
}

func main(){
	lis,err:=net.Listen("tcp",":1234")
	if err!=nil{
		log.Println("server can't listen grpc",err)
	}else{
		log.Println("server can listen grpc")
	}
	grpcServer:=grpc.NewServer()
	proto.RegisterServiceServer(grpcServer,&Server{})
	reflection.Register(grpcServer)
	if err:=grpcServer.Serve(lis);err!=nil{
		log.Println("server can't serve",err)
	}else {
		log.Println("server running")
	}
}