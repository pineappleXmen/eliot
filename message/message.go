package message

import (
	"eliot/logs"
	"time"
)

type Type int8

const (
	HeartBeat Type = iota
	AppendEntries
	RequestVote
	HeartbeatTimeout
	ElectionTimeout
)

type Message struct {
	Type      Type
	TimeStamp time.Time
	Content   interface{}
}

type AppendEntriesContent struct {
	LeaderID     int64
	Term         int64
	PrevLogIndex int64
	PrevLogTerm  int64
	Entries      []logs.Entry
	LeaderCommit int64
}

type RequestVoteContent struct {
	LeaderID     int64
	Term         int64
	LastLogIndex int64
	LastLogTerm  int64
}
