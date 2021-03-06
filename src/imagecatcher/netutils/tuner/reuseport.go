// +build linux darwin dragonfly freebsd netbsd openbsd

// Package reuseport provides TCP net.Listener with SO_REUSEPORT support.
//
// SO_REUSEPORT allows linear scaling server performance on multi-CPU servers.
// See https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ for more details :)
//
// The package is based on https://github.com/kavu/go_reuseport .
package tuner

import (
	"fmt"
	"syscall"

	"imagecatcher/logger"
)

// ErrNoReusePort is returned if the OS doesn't support SO_REUSEPORT.
type ErrNoReusePort struct {
	err error
}

// Error implements error interface.
func (e *ErrNoReusePort) Error() string {
	return fmt.Sprintf("The OS doesn't support SO_REUSEPORT: %s", e.err)
}

func Reuseport(sd int) error {
	if err := syscall.SetsockoptInt(sd, syscall.SOL_SOCKET, soReusePort, 1); err != nil {
		noReuseErr := &ErrNoReusePort{err}
		logger.Printf("%s", noReuseErr)
		return noReuseErr
	}
	return nil
}
