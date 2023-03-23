package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Car struct {
	frontLeftDoor  bool
	frontRightDoor bool
	rearLeftDoor   bool
	rearRightDoor  bool
	engineLock     bool
	trunkLock      bool
	hoodLock       bool
	alarm          bool
}

func (c *Car) LockAllDoors() {
	c.frontLeftDoor = true
	c.frontRightDoor = true
	c.rearLeftDoor = true
	c.rearRightDoor = true
}

func (c *Car) LockTrunk() {
	c.trunkLock = true
}
func (c *Car) LockHood() {
	c.hoodLock = true
}
func (c *Car) LockEngine() {
	c.engineLock = true
}
func (c *Car) SetUpAlarmOn() {
	c.alarm = true
}

func (c *Car) UnlockHood() {
	c.hoodLock = false
}

func (c *Car) UnlockEngine() {
	c.engineLock = false
}

func (c *Car) UnlockTrunk() {
	c.trunkLock = false
}

func (c *Car) UnlockAllDoors() {
	c.frontLeftDoor = false
	c.frontRightDoor = false
	c.rearLeftDoor = false
	c.rearRightDoor = false
}

func (c *Car) SetUpAlarmOff() {
	c.alarm = false
}

func FacadeCarLock(c *Car) {
	c.LockHood()
	c.LockEngine()
	c.LockTrunk()
	c.LockAllDoors()
	c.SetUpAlarmOn()
}

func FacadeCarUnlock(c *Car) {
	c.UnlockHood()
	c.UnlockEngine()
	c.UnlockTrunk()
	c.UnlockAllDoors()
	c.SetUpAlarmOff()
}
