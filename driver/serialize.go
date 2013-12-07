package driver

import (
	"encoding/binary"
	"github.com/tarm/goserial"
	"io"
)

const (
	Baud = 9600
)

type Connection struct {
	rwc io.ReadWriteCloser
}

func Connect(device string) (c *Connection, err error) {
	c = &Connection{}
	c0 := &serial.Config{
		Name: device,
		Baud: Baud,
	}

	c.rwc, err = serial.OpenPort(c0)
	if err != nil {
		return nil, err
	}

	return
}

// MoveRelativeSingle serializes instructions to the controller to
// move an arbitrary motor the given distance in meters. It does not
// block.
func (c *Connection) MoveRelativeSingle(motorIndex int, distance float64) {
	c.stepSingle(byte(motorIndex), int16(Steps(TurnsIn(distance))))
}

// MoveRelativeDouble serializes instructions to the controller to
// move the first and second motors the first and second distances in
// meters, respectively. It does not block.
func (c *Connection) MoveRelativeDouble(first, second float64) {
	firstSteps := Steps(TurnsIn(first))
	secondSteps := Steps(TurnsIn(second))
	c.stepDouble(int16(firstSteps), int16(secondSteps))
}

func (c *Connection) stepSingle(motorIndex byte, stepCount int16) {
	c.rwc.Write([]byte("t"))
	c.rwc.Write([]byte{motorIndex})
	binary.Write(c.rwc, binary.LittleEndian, stepCount)
}

func (c *Connection) stepDouble(firstSteps int16, secondSteps int16) {
	c.rwc.Write([]byte("i"))
	binary.Write(c.rwc, binary.LittleEndian, firstSteps)
	binary.Write(c.rwc, binary.LittleEndian, secondSteps)
}


func (c *Connection) PenDown(){
	c.rwc.Write([]byte("d"))
}

func (c *Connection) PenUp(){
	c.rwc.Write([]byte("u"))
}
