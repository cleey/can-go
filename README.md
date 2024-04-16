# :electric_plug: CAN Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/cleey/can-go)](https://pkg.go.dev/github.com/cleey/can-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/cleey/can-go)](https://goreportcard.com/report/github.com/cleey/can-go)
[![Codecov](https://codecov.io/gh/einride/can-go/branch/master/graph/badge.svg)](https://codecov.io/gh/einride/can-go)

CAN toolkit for Go programmers.

can-go makes use of the Linux SocketCAN abstraction for CAN communication. (See
the [SocketCAN](https://www.kernel.org/doc/Documentation/networking/can.txt)
documentation for more details).

## Examples

### Use DBC file to parse OBC CanFd data

```go
func main() {
	dbcFile = "/path/xxxx.dbc"
	obcCanFd, err := obc.NewObcCanFD(dbcFile)

	canID = uint32(298)
	canDataHexStr = "1A02A300C23290802"
	params := obcCanFd.ParseCanHexStr(canID, canDataHexStr)
	fmt.Printf("canID(%d), canData(%v)", canID, params)
}
```


### Setting up a CAN interface

```go
func main() {
	// Error handling omitted to keep example simple
	d, _ := candevice.New("can0")
	_ := d.SetBitrate(250000)
	_ := d.SetUp()
	defer d.SetDown()
}
```

### Receiving CAN frames

Receiving CAN frames from a socketcan interface.

```go
func main() {
	// Error handling omitted to keep example simple
	conn, _ := socketcan.DialContext(context.Background(), "can", "can0")

	recv := socketcan.NewReceiver(conn)
	for recv.Receive() {
		frame := recv.Frame()
		fmt.Println(frame.String())
	}
}
```

### Sending CAN frames/messages

Sending CAN frames to a socketcan interface.

```go
func main() {
	// Error handling omitted to keep example simple

	conn, _ := socketcan.DialContext(context.Background(), "can", "can0")

	frame := can.Frame{}
	tx := socketcan.NewTransmitter(conn)
	_ = tx.TransmitFrame(context.Background(), frame)
}
```

### Generating Go code from a DBC file

It is possible to generate Go code from a `.dbc` file.

```
$ go run github.com/cleey/can-go/cmd/cantool generate <dbc file root folder> <output folder>
```

In order to generate Go code that makes sense, we currently perform some
validations when parsing the DBC file so there may need to be some changes on
the DBC file to make it work

After generating Go code we can marshal a message to a frame:

```go
// import etruckcan "github.com/myproject/myrepo/gen"

auxMsg := etruckcan.NewAuxiliary().SetHeadLights(etruckcan.Auxiliary_HeadLights_LowBeam)
frame := auxMsg.Frame()
```

Or unmarshal a frame to a message:

```go
// import etruckcan "github.com/myproject/myrepo/gen"

// Error handling omitted for simplicity
_ := recv.Receive()
frame := recv.Frame()

var auxMsg *etruckcan.Auxiliary
_ = auxMsg.UnmarshalFrame(frame)

```

## Running integration tests

Building the tests:

```shell
$ make build-integration-tests
```

Built tests are placed in build/tests.

The candevice test requires access to physical HW, so run it using sudo.
Example:

```shell
$ sudo ./build/tests/candevice.test
> PASS
```
