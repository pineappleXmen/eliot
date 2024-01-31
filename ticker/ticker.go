package ticker

import (
	"eliot/message"
	"time"
)

// Ticker 是 Raft 协议中的时间管理模块
type Ticker struct {
	heartbeatInterval time.Duration
	electionTimeout   time.Duration
	stopChan          chan struct{}
	messageChan       chan message.Message
}

// NewTicker 创建一个新的 Ticker 实例
func NewTicker(heartbeatInterval, electionTimeout time.Duration) *Ticker {
	return &Ticker{
		heartbeatInterval: heartbeatInterval,
		electionTimeout:   electionTimeout,
		stopChan:          make(chan struct{}),
		messageChan:       make(chan message.Message),
	}
}

// Start 启动 Ticker
func (t *Ticker) Start() {
	go t.runTicker()
}

// runTicker 执行 Ticker 的主循环
func (t *Ticker) runTicker() {
	heartbeatTicker := time.NewTicker(t.heartbeatInterval)
	electionTicker := time.NewTicker(t.electionTimeout)
	defer func() {
		heartbeatTicker.Stop()
		electionTicker.Stop()
		close(t.messageChan)
	}()

	for {
		select {
		case <-heartbeatTicker.C:
			// 发送心跳定时事件
			t.messageChan <- message.Message{Type: message.HeartBeat, TimeStamp: time.Now(), Content: nil}
		case <-electionTicker.C:
			// 发送选举超时定时事件
			t.messageChan <- message.Message{Type: message.ElectionTimeout, TimeStamp: time.Now(), Content: nil}
		case <-t.stopChan:
			// 收到停止信号，退出循环
			return
		}
	}
}
