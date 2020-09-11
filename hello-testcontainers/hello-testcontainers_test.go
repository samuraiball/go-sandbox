package hello_testcontainers__test

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHelloTestContainerWithCompose(t *testing.T) {

	composeFilePaths := []string{"docker-compose.yml"}
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, strings.ToLower(uuid.New().String()))
	execErr := compose.Invoke()
	checkErr(t, execErr.Error)
}

func TestHelloTestContainer(t *testing.T) {
	cxt := context.Background()

	request := testcontainers.ContainerRequest{
		Image:        "nginx",
		ExposedPorts: []string{"80/tcp"},
		WaitingFor:   wait.ForHTTP("/"),
		AutoRemove:   false,
	}

	container, err := testcontainers.GenericContainer(cxt, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
	checkErr(t, err)

	defer container.Terminate(cxt)

	host, err := container.Host(cxt)
	checkErr(t, err)

	port, err := container.MappedPort(cxt, "80")
	checkErr(t, err)

	resp, err := http.Get(fmt.Sprintf("http://%s:%s", host, port.Port()))
	checkErr(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(t, err)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("get wrong status code")
	}
	fmt.Println(string(body))
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
