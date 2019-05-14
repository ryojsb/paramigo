package paramigo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func push(c *cli.Context) {
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}

	// check private key or password.
	var passwd string
	auth := []ssh.AuthMethod{}
	if c.String("key") != "" {
		key, err := ioutil.ReadFile(c.String("key"))
		ce(err, "private key")

		signer, err := ssh.ParsePrivateKey(key)
		ce(err, "signer")

		auth = append(auth, ssh.PublicKeys(signer))
	} else {
		// check password.
		if c.String("password") != "" {
			passwd = c.String("password")
		} else {
			fmt.Print("Password: ")
			inPasswd, err := terminal.ReadPassword(int(syscall.Stdin))
			ce(err, "password")
			passwd = string(inPasswd)
		}
		auth = append(auth, ssh.Password(passwd))
	}

	// set ssh config.
	sshConfig := &ssh.ClientConfig{
		User:            c.String("user"),
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH connect.
	client, err := ssh.Dial("tcp", c.String("host")+":"+c.String("port"), sshConfig)
	ce(err, "dial")

	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	// Finally, run the command
	err = session.Run(c.String("cmd"))
	fmt.Println(b.String())
}

func SendCommand() {
	app := cli.NewApp()
	app.Name = "ssh command runner"
	app.Usage = "SSh command sending Tool"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
			Usage: "SSH connect host",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "22",
			Usage: "SSH connect port",
		},
		cli.StringFlag{
			Name:  "user, u",
			Value: "root",
			Usage: "SSH login user",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "SSH login user password",
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "SSH private key",
		},
		cli.StringFlag{
			Name:  "cmd",
			Usage: "command",
		},
	}

	app.Action = push
	app.Run(os.Args)
}
