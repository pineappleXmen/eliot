package raft

import (
	"eliot/logs"
	"eliot/ticker"
	"sync"
)

type State int

const (
	Follower State = iota
	Leader
	Candidate
	Dead
)

type Raft struct {
	mu    sync.Mutex // Lock to protect shared access to this peer's state
	peers []*Raft    // RPC end points of all peers
	me    int        // this peer's index into peers[]
	dead  int32      // set by Kill()

	//applyCh   chan ApplyMsg
	//applyCond *sync.Cond

	state  State
	ticker ticker.Ticker

	//persistent state
	term     int64
	votedFor int64
	log      logs.Log

	lastIncludeIndex int64
	lastIncludeTerm  int64

	//volatile state on all servers
	commitIndex int64
	lastApplied int64

	//Snapshot state
	snapshot      []byte
	snapshotIndex int64
	snapshotTerm  int64

	waitingSnapshot []byte
	waitingIndex    int64
	waitingTerm     int64

	//volatile state on leaders
	nextIndex  []int64
	matchIndex []int64
}

func (rf *Raft) currentTerm() int64 {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.term
}

func (rf *Raft) currentState() State {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.state
}
