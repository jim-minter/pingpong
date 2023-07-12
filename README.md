# Let's play pingpong!

### Create cluster

```bash
NAME=jimminter
LOCATION=eastus

az group create -g $NAME -l $LOCATION
az aks create -g $NAME -n $NAME --network-plugin azure
az aks get-credentials -g $NAME -n $NAME

kubectl apply -f pingpong.yaml
```

### Some tools to investigate

1. `kubectl get node,pod,svc -A -o wide` to show nodes/pods/services and IPs

1. [ssh-k8s](https://raw.githubusercontent.com/jim-minter/utils/master/ssh-k8s) to connect to a node

1. `apt install net-tools` to install net-tools

1. `ifconfig` (`ip a show`) to show interfaces

1. `route` (`ip route show`) to show route table

1. `arp` (`cat /proc/net/arp`) to show arp table

1. `iptables -t nat -nvL` to show iptables NAT rules

1. `netstat -lnp` to list processes listening on sockets

1. `tcpdump -n -p -i $IF -s0 -X` to dump network trafic on an interface

1. `nsenter -n -t $PID bash` to enter the network namespace of a PID
