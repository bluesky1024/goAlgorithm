package raft

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"net"
	"os"
	"time"
)

func NewNode() {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID("127.0.0.1")
	raftConfig.Logger = hclog.New(nil)

	logStore, err := raftboltdb.NewBoltStore("raft-log.bolt")
	stableStore, err := raftboltdb.NewBoltStore("raft-stable.bolt")

	snapshotStore, err := raft.NewFileSnapshotStore("/tmp/raft", 1, os.Stderr)

	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1")
	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
}
