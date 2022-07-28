package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	grpcServer "github.com/netweng/go-simple-server/internal/grpc"
	pb "github.com/netweng/go-simple-server/proto"
	"github.com/openlyinc/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	vm "github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	models "github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func path(p string) string {
	return fmt.Sprintf(p)
}

func setupRouter() *gin.Engine {

	cloudClient, err := setUpTowerServer()
	if err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		// fmt.Println('add')
		c.String(http.StatusOK, "ok")
	})

	router.POST("/get-vms", func(c *gin.Context) {

		// fmt.Println("the go server is listen on: 9001")
		getVmParams := vm.NewGetVmsParams()
		getVmParams.RequestBody = &models.GetVmsRequestBody{First: pointy.Int32(1)}
		getVmRes, err := cloudClient.VM.GetVms(getVmParams)

		if err != nil {
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, getVmRes)
		}
	})

	return router
}

func setUpTowerServer() (*apiclient.Cloudtower, error) {
	client, err := apiclient.NewWithUserConfig(
		apiclient.ClientConfig{Host: "localhost:8090", BasePath: "", Schemes: []string{"http"}},
		apiclient.UserConfig{Name: "xiaojun", Password: "smartx930521.", Source: models.UserSourceLDAP})
	return client, err
}

const PORT = "9001"

func main() {
	// r := setupRouter()
	lis, err := net.Listen("tcp", "127.0.0.1:"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterGoServerServer(gs, &grpcServer.Backend{})

	reflection.Register(gs)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
