package partition_identity

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	clustering "github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/log"
)

type placementActor struct {
	cluster          *clustering.Cluster
	partitionManager *PartitionManager
	actors           map[string]GrainMeta
}

func newPlacementActor(c *clustering.Cluster, pm *PartitionManager) *placementActor {
	return &placementActor{
		cluster:          c,
		partitionManager: pm,
		actors:           map[string]GrainMeta{},
	}
}

func (p *placementActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Terminated:
		p.onTerminated(msg, ctx)
	case *clustering.IdentityHandoverRequest:
		p.onIdentityHandoverRequest(msg, ctx)
	case *clustering.ActivationRequest:
		p.onActivationRequest(msg, ctx)
	default:
		plog.Error("Invalid message", log.TypeOf("type", msg), log.PID("sender", ctx.Sender()))
	}
}

func (p *placementActor) onIdentityHandoverRequest(msg *clustering.IdentityHandoverRequest, ctx actor.Context) {

}

func (p *placementActor) onActivationRequest(msg *clustering.ActivationRequest, ctx actor.Context) {

}

func (p *placementActor) onTerminated(msg *actor.Terminated, ctx actor.Context) {

}