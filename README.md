# nxt webtun

```
MacOS only
$ sudo go run nxt_tap.go
1. This module creates a virtual interface utunX
2. Assign point to point IP (P-t-P) 10.1.0.1 -> 10.1.0.2
3. ping 10.1.0.2 will send an ICMP packet to 10.1.0.2 via utunX
4. Modify the Dst Address to 127.0.0.1, inject the packet back to utunX
5. TODO .. route this packet to another go listener.

$ netstat -rn | more
default            192.168.1.1        UGSc           en0       
10.1.0.2           10.1.0.1           UH           utun4       

$ ifconfig -a
utun4: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1500
	inet 10.1.0.1 --> 10.1.0.2 netmask 0xff000000 

$ sudo go run nxt_tap.go
2020/12/18 18:31:30 Packet Received 84: 45 00 00 54 ee 0f 00 00 40 01 78 95 0a 01 00 01 0a 01 00 02 08 00 6a 13 4c 50 00 00 5f dd 66 02 00 05 90 b4 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f 30 31 32 33 34 35 36 37
Sending to remote: ver=4 hdrlen=20 tos=0x0 totallen=21524 id=0xee0f flags=0x0 fragoff=0x0 ttl=64 proto=1 cksum=0x7895 src=10.1.0.1 dst=10.1.0.2 (<nil>)

```

```
Packets sent by an operating system via a tun/tap device are delivered to a user-space program which attaches itself to the device. A user-space program may also pass packets into a tun/tap device. In this case the tun/tap device delivers (or “injects”) these packets to the operating-system network stack thus emulating their reception from an external source. tun/tap interfaces are software-only interfaces, meaning that they exist only in the kernel and, unlike regular network interfaces, they have no physical hardware component (and so there’s no physical wire connected to them).
```
