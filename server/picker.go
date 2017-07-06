package server


type RemotePiker interface {
	Pick() *Remote
}