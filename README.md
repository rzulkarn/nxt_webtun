# nxt webtun

```
$ip addr show
6: sample_dev: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 3a:26:66:8d:e2:00 brd ff:ff:ff:ff:ff:ff

$sudo ip addr add 10.1.0.10/24 dev sample_dev
$sudo ip link set dev sample_dev up
$ip addr show
9: sample_dev: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 1000
    link/ether 26:6c:c8:a5:69:48 brd ff:ff:ff:ff:ff:ff
    inet 10.1.0.10/24 scope global sample_dev
       valid_lft forever preferred_lft forever
    inet6 fe80::246c:c8ff:fea5:6948/64 scope link 
       valid_lft forever preferred_lft forever
```

```
Packets sent by an operating system via a tun/tap device are delivered to a user-space program which attaches itself to the device. A user-space program may also pass packets into a tun/tap device. In this case the tun/tap device delivers (or “injects”) these packets to the operating-system network stack thus emulating their reception from an external source. tun/tap interfaces are software-only interfaces, meaning that they exist only in the kernel and, unlike regular network interfaces, they have no physical hardware component (and so there’s no physical wire connected to them).
```
