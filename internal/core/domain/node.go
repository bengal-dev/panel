package domain

import "context"

type Node struct {
	ID          string
	IP          string
	TunnelUnits []TunnelUnit
}

type NodeStats struct {
	NodeID      string
	LoadPercent uint8
}

type NodeRepository interface {
	Create(ctx context.Context, node Node) (newNode Node, err error)
	Delete(ctx context.Context, id string) (err error)

	GetByID(ctx context.Context, id string) (node Node, err error)
	GetActive(ctx context.Context) (nodes []Node, err error)
	GetAll(ctx context.Context) (nodes []Node, err error)
}

type NodeController interface {
	StartNode(ctx context.Context) (node Node, err error)
	StopNode(ctx context.Context) (err error)
}
