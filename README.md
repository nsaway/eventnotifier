# eventnotifier
NSAway System Event Notifier

Make sure sources/etc/dbus-1/system.d/com.NSAway.EventNotifier.conf is in /etc/dbus-1/system.d.

forked from https://github.com/subgraph/eventnotifier

### Dependencies
To build  eventnotifier you need to install dependencies:

`$ go get github.com/godbus/dbus`

`$ go get github.com/TheCreeper/go-notify`

### Clone and building
`$ git clone https://github.com/nsaway/eventnotifier && cd eventnotifier`

`$ go build`

`# cp source/dbus-1/systemd/com.NSAway.EventNotifier.conf /etc/dbus-1/system.d/`

### Usage
See the exemple [here](https://github.com/nsaway/eventnotifier/tree/master/example)



