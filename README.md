proto.actor cluster provider by zookeeper
============================================

```
package main

import (
	"flag"
	"time"

	"cluster-metrics/shared"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/AsynkronIT/protoactor-go/remote"
	zk "github.com/qjpcpu/proto.actor-zk"
)

func main() {
	port := flag.Int("port", 0, "")
	ip := flag.String("ip", "127.0.0.1", "")
	flag.Parse()

	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(*ip, *port)

	props := actor.PropsFromProducer(newHelloActor)
	helloKind := cluster.NewKind("Hello", props)

	provider, _ := zk.New([]string{"127.0.0.1:2181"}, zk.WithBaseKey("proto.actors"))
	clusterConfig := cluster.Configure("my_cluster", provider, remoteConfig, helloKind)
	c := cluster.New(system, clusterConfig)
	c.Start()

	_, _ = console.ReadLine()
	c.Shutdown(true)
}
```