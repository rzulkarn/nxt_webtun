# nxt webtun

12/19/2020 - On iOS/MacOS, starting from 2019, it introduced System/Network Extension to allow apps developer to do content filtering, DNS Proxy, VPN types of application. It allows to run the app user permission aka. “Trusted app”. It usable Swift or Objective C, similarly VPNService on Android written in Java. OpenVPN written in C, uses freebsd specific library to implement VPN like service.  

```
Experiment on MacOS only
$ sudo go run nxt_tap.go
1. This module creates a virtual interface utunX
2. Assign point to point IP (P-t-P) 10.1.0.1 -> 10.1.0.2
3. ping 10.1.0.2 will send an ICMP packet to 10.1.0.2 via utunX
4. Modify the Dst Address to 127.0.0.1, inject the packet back to utunX
5. TODO .. route this packet to another go listener.
6. TODO .. program the default route to utun4

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
Cisco AnyConnect routing table (sample ifconfig and netstat when enabled)

rzulkarn$ ifconfig utun1
utun1: flags=80d1<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1250
       inet 192.168.25.207 --> 192.168.25.207 netmask 0xffffffff

Destination        Gateway            Flags        Refs      Use   Netif Expire
default            192.168.25.207     UGSc          114        0   utun1      
192.168.1          link#13            UCS             2        0   utun1      
192.168.1.1/32     link#13            UCS             0        0   utun1      
192.168.1.121      link#13            UHW3I           0        0   utun1      7
192.168.1.196      link#13            UHW3I           0        0   utun1      7
192.168.25.207/32  link#13            UCS             1        0   utun1      
192.168.25.207     link#13            UHWIir         25       17   utun1      
224.0.0/4          link#13            UmCS            1        0   utun1      
224.0.0.251        link#13            UHmW3I          0        0   utun1      7
255.255.255.255/32 link#13            UCS             0        0   utun1      
```
Packets sent by an operating system via a tun/tap device are delivered to a user-space program which attaches itself to the device. A user-space program may also pass packets into a tun/tap device. In this case the tun/tap device delivers (or “injects”) these packets to the operating-system network stack thus emulating their reception from an external source. tun/tap interfaces are software-only interfaces, meaning that they exist only in the kernel and, unlike regular network interfaces, they have no physical hardware component (and so there’s no physical wire connected to them).
