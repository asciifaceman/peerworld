package server

import (
	"fmt"

	"github.com/asciifaceman/peerworld/pkg/player"
)

func NewServer(host string, port int, leader bool) *Server {
	p := player.NewPlayer(leader)

	return &Server{
		player: p,
		host:   host,
		port:   port,
	}
}

func NewLeader(host string, port int) *Server {
	s := NewServer(host, port, true)
	s.leader = true
	return s
}

func NewFollower(host string, port int) *Server {
	s := NewServer(host, port, false)
	s.leader = false
	return s
}

// Server is the central dispatch that handles intercommunications between players
type Server struct {
	leader bool
	player *player.Player
	host   string
	port   int
}

// Initialize is the initial standup of services
func (s *Server) Initialize() error {
	fmt.Printf("%s has been conceived!\n", s.player.FmtFullName())

	return nil
}

func (s *Server) fmtHostPort() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}
