package tclient

import (
	"fmt"
	"net"
	"time"
)

type TelnetClient struct {
	// Timeout of read/write operations
	Timeout		int
	// TimeoutGlobal is global operation timeout; i.e. like stucked in DLink-like refreshing pagination
	TimeoutGlobal	int
	Prompt		string
	conn		net.Conn
	closed		bool
	Options		[]int

	loginPrompt			string
	passwordPrompt		string

	reader		bool
	writer		bool
}

func New(tout int, prompt string) *TelnetClient {
	if tout < 1 {
		tout = 1
	}
	c := TelnetClient{
		Timeout: tout,
		Prompt: `(?msi:[\$%#>]$)`,
		loginPrompt: `[Uu]ser[Nn]ame\:$`,
		passwordPrompt: `[Pp]ass[Ww]ord\:$`,
		closed:false,
		Options: make([]int,0),
	}

	if prompt != "" {
		c.Prompt = prompt
	}

	// Global timeout defaults to 3 * rw timeout
	c.TimeoutGlobal = c.Timeout * 2

	// set default options
	// we will accept an offer from remote sidie for it to echo and suppress goaheads
	c.SetOpts([]int{TELOPT_ECHO, TELOPT_SGA})

	return &c
}

func (c *TelnetClient) GlobalTimeout(t int) {
	c.TimeoutGlobal = t
}

// SetLoginPrompt sets login prompt
func (c *TelnetClient) SetLoginPrompt(s string) {
	c.loginPrompt = s
}

func (c *TelnetClient) SetPasswordPrompt(s string) {
	c.passwordPrompt = s
}

func (c *TelnetClient) Close() {
	c.conn.Close()
}

// FlushOpts flushes default options set
func (c *TelnetClient) FlushOpts() {
	c.Options = make([]int, 0)
}

// SetOpts prepares default options set
func (c *TelnetClient) SetOpts(opts []int) error {
	for _, opt := range opts {
		if opt > 255 {
			return fmt.Errorf("Bad telnet option %d: option > 255", opt)
		}

		c.Options = append(c.Options, opt)
	}

	return nil
}

// Open tcp connection
func (c *TelnetClient) Open(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)

	var err error
	c.conn, err = net.DialTimeout("tcp", addr, time.Second * time.Duration(c.Timeout))
	if err != nil {
		c.closed = true
		return err
	}

	return nil
}

