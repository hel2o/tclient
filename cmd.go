package tclient

// Cmd sends command and returns output
func (c *TelnetClient) Cmd(cmd string) (string, error) {
	err := c.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	return c.ReadUntil(c.prompt)
}

// Cmd sends command then wait special character and returns output
func (c *TelnetClient) CmdUntil(cmd, until string) (out string, err error) {
	err = c.Write([]byte(cmd))
	if err != nil {
		return
	}

	out, err = c.ReadUntil(until)
	if err != nil {
		return
	}
	return
}
