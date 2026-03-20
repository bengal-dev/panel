package domain

import (
	"errors"
	"fmt"
)

type TunnelProtocol uint8

const (
	TunnelProtocolWireGuard TunnelProtocol = iota
	TunnelProtocolMTProto
	TunnelProtocolVMess
	TunnelProtocolVLESS
	TunnelProtocolTrojan
	TunnelProtocolShadowsocks
)

var (
	ErrUnknownProtocolType = errors.New("unknown protocol type")
	ErrInvalidNodeInfo     = errors.New("invalid tunnel unit data")
	ErrInvalidAccessType   = errors.New("invalid access type for protocol")
)

type TunnelAccessData interface {
	isTunnelAccessData()
}

type WireGuardAccess struct {
	PublicKey  string
	PrivateKey string
}

func (WireGuardAccess) isTunnelAccessData() {}

type VLESSAccess struct {
	UUID string
}

func (VLESSAccess) isTunnelAccessData() {}

type VMessAccess struct {
	UUID string
}

func (VMessAccess) isTunnelAccessData() {}

type TrojanAccess struct {
	Password string
}

func (TrojanAccess) isTunnelAccessData() {}

type ShadowsocksAccess struct {
	Password string
	Method   string
}

func (ShadowsocksAccess) isTunnelAccessData() {}

type MTProtoAccess struct {
	Secret string
}

func (MTProtoAccess) isTunnelAccessData() {}

type TunnelUnit struct {
	IP       string
	Port     uint16
	Protocol TunnelProtocol
	Access   TunnelAccessData
}

func (t TunnelUnit) IsValid() (err error) {
	if t.IP == "" || t.Port == 0 || t.Access == nil {
		return ErrInvalidNodeInfo
	}

	switch t.Protocol {
	case TunnelProtocolWireGuard:
		if _, ok := t.Access.(WireGuardAccess); !ok {
			return ErrInvalidAccessType
		}

	case TunnelProtocolVLESS:
		if _, ok := t.Access.(VLESSAccess); !ok {
			return ErrInvalidAccessType
		}

	case TunnelProtocolVMess:
		if _, ok := t.Access.(VMessAccess); !ok {
			return ErrInvalidAccessType
		}

	case TunnelProtocolTrojan:
		if _, ok := t.Access.(TrojanAccess); !ok {
			return ErrInvalidAccessType
		}

	case TunnelProtocolShadowsocks:
		if _, ok := t.Access.(ShadowsocksAccess); !ok {
			return ErrInvalidAccessType
		}

	case TunnelProtocolMTProto:
		if _, ok := t.Access.(MTProtoAccess); !ok {
			return ErrInvalidAccessType
		}

	default:
		return ErrUnknownProtocolType
	}

	return nil
}

func NewTunnelUnit(ip string, port uint16, protocol TunnelProtocol, access TunnelAccessData) (uint TunnelUnit, err error) {
	t := TunnelUnit{
		IP:       ip,
		Port:     port,
		Protocol: protocol,
		Access:   access,
	}

	if err := t.IsValid(); err != nil {
		return TunnelUnit{}, err
	}

	return t, nil
}

func (t TunnelUnit) Address() (address string) {
	return fmt.Sprintf("%s:%d", t.IP, t.Port)
}
