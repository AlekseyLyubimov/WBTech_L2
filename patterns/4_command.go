package patterns

type Command interface {
    execute()
}

type Switch struct {
    command Command
}

func (s *Switch) press() {
    s.command.execute()
}

type OnCommand struct {
    device Device
}

func (c *OnCommand) execute() {
    c.device.on()
}

type OffCommand struct {
    device Device
}

func (c *OffCommand) execute() {
    c.device.off()
}

type Device interface {
    on()
    off()
}

type Tv struct {
    isOn bool
}

func (t *Tv) on() {
    t.isOn = true
    println("Turning tv on")
}

func (t *Tv) off() {
    t.isOn = false
    println("Turning tv off")
}
