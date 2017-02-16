package socket

import "github.com/davyxu/cellnet"

// Peer间的共享数据
type peerBase struct {
	cellnet.EventQueue
	name          string
	address       string
	maxPacketSize int

	recvHandler cellnet.EventHandler

	sendHandler cellnet.EventHandler

	*cellnet.DispatcherHandler

	codec cellnet.Codec
}

func (self *peerBase) nameOrAddress() string {
	if self.name != "" {
		return self.name
	}

	return self.address
}

func (self *peerBase) Address() string {
	return self.address
}

func (self *peerBase) PacketCodec() cellnet.Codec {
	return self.codec
}

func (self *peerBase) SetPacketCodec(c cellnet.Codec) {
	self.codec = c
}

func (self *peerBase) SetHandler(recv, send cellnet.EventHandler) {
	self.recvHandler = recv
	self.sendHandler = send
}

func (self *peerBase) GetHandler() (recv, send cellnet.EventHandler) {
	return self.recvHandler, self.sendHandler
}

func (self *peerBase) SetName(name string) {
	self.name = name
}

func (self *peerBase) Name() string {
	return self.name
}

func (self *peerBase) SetMaxPacketSize(size int) {
	self.maxPacketSize = size
}

func (self *peerBase) MaxPacketSize() int {
	return self.maxPacketSize
}

var DefaultCodec string = "pb"

func newPeerBase(queue cellnet.EventQueue) *peerBase {

	self := &peerBase{
		EventQueue:        queue,
		codec:             cellnet.FetchCodec(DefaultCodec),
		DispatcherHandler: cellnet.NewDispatcherHandler(),
	}

	self.recvHandler = BuildRecvHandler(EnableMessageLog, self.DispatcherHandler, queue)

	self.sendHandler = BuildSendHandler(EnableMessageLog)

	return self
}
