package netdriver

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"
)

// Device implements tinygo.org/x/drivers/net.Adapter interface.
type Driver struct {
	conn          net.Conn
	Debug         bool
	MaxPacketSize int
}

// ConnectToAccessPoint connects to an access point.
func (d *Driver) ConnectToAccessPoint(ssid, pass string, timeout time.Duration) error {
	if d.Debug {
		fmt.Printf("ConnectToAccessPoint\n")
	}
	return nil
}

// Disconnect disconnects the connection.
func (d *Driver) Disconnect() error {
	if d.Debug {
		fmt.Printf("Disconnect\n")
	}
	return nil
}

// GetClientIP get th client IP.
func (d *Driver) GetClientIP() (string, error) {
	if d.Debug {
		fmt.Printf("GetClientIP()\n")
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}

	return "", errors.New("error GetClientIP()")
}

// GetDNS get host by name.
func (d *Driver) GetDNS(domain string) (string, error) {
	if d.Debug {
		fmt.Printf("GetDNS\n")
	}
	return "", nil
}

// ConnectTCPSocket connects TCP socket.
func (d *Driver) ConnectTCPSocket(addr, port string) error {
	if d.Debug {
		fmt.Printf("ConnectTCPSocket(%q, %q)\n", addr, port)
	}

	serverAddr, err := net.ResolveTCPAddr("tcp", addr+":"+port)
	if err != nil {
		fmt.Printf("err2: %#v\n", err.Error())
		return err
	}

	conn, err := net.Dial("tcp", serverAddr.AddrPort().String())
	if err != nil {
		fmt.Printf("err3: %#v\n", err.Error())
		return err
	}
	d.conn = conn
	return nil
}

// ConnectSSLSocket connects SSL socket.
func (d *Driver) ConnectSSLSocket(addr, port string) error {
	if d.Debug {
		fmt.Printf("ConnectSSLSocket\n")
	}

	err := d.ConnectTCPSocket(addr, port)
	if err != nil {
		return err
	}

	conn := tls.Client(d.conn, &tls.Config{
		InsecureSkipVerify: true,
	})

	err = conn.Handshake()
	if err != nil {
		return err
	}
	d.conn = conn

	return nil
}

// ConnectUDPSocket connects UDP socket.
func (d *Driver) ConnectUDPSocket(addr, sendport, listenport string) error {
	if d.Debug {
		fmt.Printf("ConnectUDPSocket\n")
	}
	return nil
}

// DisconnectSocket disconnects the socket
func (d *Driver) DisconnectSocket() error {
	if d.Debug {
		fmt.Printf("DisconnectSocket\n")
	}
	if d == nil {
		return nil
	}
	if d.conn == nil {
		return nil
	}
	return d.conn.Close()
}

// StartSocketSend ...
func (d *Driver) StartSocketSend(size int) error {
	if d.Debug {
		fmt.Printf("StartSocketSend(%d)\n", size)
	}
	return nil
}

// Write writes data.
func (d *Driver) Write(b []byte) (n int, err error) {
	if d.Debug {
		fmt.Printf("Write(%q)\n", string(b))
	}
	d.conn.Write(b)
	return 0, nil
}

// ReadSocket read data.
func (d *Driver) ReadSocket(b []byte) (n int, err error) {
	if d.Debug {
		fmt.Printf("ReadSocket()\n")
	}
	max := len(b)
	if d.MaxPacketSize > 0 {
		max = d.MaxPacketSize
	}
	n, err = d.conn.Read(b[:max])
	if false {
		fmt.Printf("  %d %q\n", n, string(b[:n]))
	}
	return n, err
}

// IsSocketDataAvailable ...
func (d *Driver) IsSocketDataAvailable() bool {
	if d.Debug {
		fmt.Printf("IsSocketDataAvailable\n")
	}
	return false
}

// Response ...
func (d *Driver) Response(timeout int) ([]byte, error) {
	if d.Debug {
		fmt.Printf("Response(%d)\n", timeout)
	}
	return nil, nil
}
