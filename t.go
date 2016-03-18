package main

import (
	"log"
	"os"
	"time"
)

// func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

type service struct {
	exec string;
	args []string;
	proc *os.Process;
	tm <-chan time.Time;
}

func (srv *service) srv_h () {
	for {
		select {
		case _ = <- srv.tm:
			attr := new (os.ProcAttr);
			attr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
			cmd, err := os.StartProcess(srv.exec, srv.args, attr);
			if err != nil {
				log.Fatal(err)
			}
			if _, err := cmd.Wait(); err != nil {
				log.Fatal(err)
			}
			srv.tm = time.After (5 * time.Second);
		}
	}
}

func newService(exec string, args []string) *service {
	srv := &service{
		exec: exec,
		args: args,
		proc: nil,
		tm: time.After (time.Second),
	}
	go srv.srv_h ();
	return srv;
}

func main() {
	newService ("/usr/bin/echo", []string{"echo","a"});
	for {
		time.Sleep(3);
	}
}
