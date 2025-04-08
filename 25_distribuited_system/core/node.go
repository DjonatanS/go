package core

import (
	"context"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// NodeServiceGrpcServer implementation
type NodeServiceGrpcServer struct {
	UnimplementedNodeServiceServer
	CmdChannel chan string
}

func (s *NodeServiceGrpcServer) ReportStatus(ctx context.Context, request *Request) (*Response, error) {
	return &Response{Data: "ok"}, nil
}

func (s *NodeServiceGrpcServer) AssignTask(request *Request, server NodeService_AssignTaskServer) error {
	for {
		select {
		case cmd := <-s.CmdChannel:
			if err := server.Send(&Response{Data: cmd}); err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

var nodeServiceGrpcServer *NodeServiceGrpcServer

func GetNodeServiceGrpcServer() *NodeServiceGrpcServer {
	if nodeServiceGrpcServer == nil {
		nodeServiceGrpcServer = &NodeServiceGrpcServer{
			CmdChannel: make(chan string),
		}
	}
	return nodeServiceGrpcServer
}

// MasterNode implementation
type MasterNode struct {
	api     *gin.Engine
	ln      net.Listener
	svr     *grpc.Server
	nodeSvr *NodeServiceGrpcServer
}

func (n *MasterNode) Init() (err error) {
	n.ln, err = net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	n.svr = grpc.NewServer()
	n.nodeSvr = GetNodeServiceGrpcServer()
	RegisterNodeServiceServer(n.svr, n.nodeSvr)
	n.api = gin.Default()
	n.api.POST("/tasks", func(c *gin.Context) {
		var payload struct {
			Cmd string `json:"cmd"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		n.nodeSvr.CmdChannel <- payload.Cmd
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	return nil
}

func (n *MasterNode) Start() {
	go n.svr.Serve(n.ln)
	_ = n.api.Run(":9092")
	n.svr.Stop()
}

var node *MasterNode

func GetMasterNode() *MasterNode {
	if node == nil {
		node = &MasterNode{}
		if err := node.Init(); err != nil {
			panic(err)
		}
	}
	return node
}
