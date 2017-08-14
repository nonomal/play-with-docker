package docker

import (
	"io"
	"net"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) CreateNetwork(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *Mock) ConnectNetwork(container, network, ip string) (string, error) {
	args := m.Called(container, network, ip)
	return args.String(0), args.Error(1)
}

func (m *Mock) GetDaemonInfo() (types.Info, error) {
	args := m.Called()
	return args.Get(0).(types.Info), args.Error(1)
}

func (m *Mock) GetDaemonHost() string {
	args := m.Called()
	return args.String(0)
}

func (m *Mock) GetSwarmPorts() ([]string, []uint16, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Get(1).([]uint16), args.Error(2)
}

func (m *Mock) GetPorts() ([]uint16, error) {
	args := m.Called()
	return args.Get(0).([]uint16), args.Error(1)
}
func (m *Mock) GetContainerStats(name string) (io.ReadCloser, error) {
	args := m.Called(name)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}
func (m *Mock) ContainerResize(name string, rows, cols uint) error {
	args := m.Called(name, rows, cols)
	return args.Error(0)
}
func (m *Mock) ContainerRename(old, new string) error {
	args := m.Called(old, new)
	return args.Error(0)
}
func (m *Mock) CreateAttachConnection(name string) (net.Conn, error) {
	args := m.Called(name)
	return args.Get(0).(net.Conn), args.Error(1)
}
func (m *Mock) CopyToContainer(containerName, destination, fileName string, content io.Reader) error {
	args := m.Called(containerName, destination, fileName, content)
	return args.Error(0)
}
func (m *Mock) DeleteContainer(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *Mock) CreateContainer(opts CreateContainerOpts) (string, error) {
	args := m.Called(opts)
	return args.String(0), args.Error(1)
}
func (m *Mock) ExecAttach(instanceName string, command []string, out io.Writer) (int, error) {
	args := m.Called(instanceName, command, out)
	return args.Int(0), args.Error(1)
}
func (m *Mock) DisconnectNetwork(containerId, networkId string) error {
	args := m.Called(containerId, networkId)
	return args.Error(0)
}
func (m *Mock) DeleteNetwork(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *Mock) Exec(instanceName string, command []string) (int, error) {
	args := m.Called(instanceName, command)
	return args.Int(0), args.Error(1)
}
func (m *Mock) SwarmInit() (*SwarmTokens, error) {
	args := m.Called()
	return args.Get(0).(*SwarmTokens), args.Error(1)
}
func (m *Mock) SwarmJoin(addr, token string) error {
	args := m.Called(addr, token)
	return args.Error(0)
}

type MockConn struct {
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return len(b), nil
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return &net.IPAddr{}
}

func (m *MockConn) RemoteAddr() net.Addr {
	return &net.IPAddr{}
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
