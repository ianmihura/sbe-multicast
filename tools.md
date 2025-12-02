# iface and router
```sh
sudo ip link set dev lo multicast on  # enable mcast on iface
ip route get 224.0.0.1  # what is the default mcast iface
sudo ip route add 224.0.0.0/4 dev lo  # set as default mcast
```

# sniff upd packets on wlan0 interface for port 9999
```sh
sudo tcpdump -ni wlan0 udp port 6200 -X
```


# send simple packet to 
```sh
echo "hello" | socat - UDP4-DATAGRAM:239.222.222.2:6200,ip-multicast-if=0.0.0.0
```

# receive packet
```sh
socat -u UDP4-RECVFROM:6200,ip-add-membership=239.222.222.2:192.168.1.132 -
socat -u UDP4-DATAGRAM:239.222.222.2:6200,ip-multicast-loop=0,ip-add-membership=239.222.222.2:127.0.0.1 -
sudo tcpdump -i wlan0 -U port 9999 -w sample.pcap  # write to sample.pcap
```

# working with pcapng files
```sh
sudo tcpdump -r ./sbe/rfq.pcapng -X -s0
tcprewrite --portmap=6100:9999,37608:49112 --pnat=127.0.0.1:192.168.1.132 --enet-dmac=38:68:93:ea:87:0b --enet-smac=38:68:93:ea:87:0b --fixcsum --infile=./../sbe/rfq.pcapng --outfile=./../sbe/rfq_9999_wlan0.pcapng
sudo tcpreplay --topspeed --loop 10 -vi wlan0 ./sbe/rfq.pcapng
```

# Other nice things to have:
Save these changes to a file to persist them
```sh
# Backlog for half-open SYNs
sudo sysctl -w net.ipv4.tcp_max_syn_backlog=4096  # it was 1024

# Increase local port range for heavy outbound clients
sudo sysctl -w net.ipv4.ip_local_port_range="1024 65535"  # it was 32768 60999

# Bump r/w memory
sudo sysctl -w net.core.rmem_max=134217728  # it was 212992
sudo sysctl -w net.core.wmem_max=134217728  # it was 212992

# Allow reuse + fast recycling of TIME_WAIT
sudo sysctl -w net.ipv4.tcp_tw_reuse=1 # it was 2
```

sudo sysctl net.ipv4.conf.all.rp_filter=0  # it was 1
sudo sysctl net.ipv4.conf.wlan0.rp_filter=0  # it was 2
sudo sysctl net.ipv4.conf.default.rp_filter=0  # it was 1

# Add ips
```sh
for i in `seq 1 20`; do sudo ip addr add 192.168.254.$i/24 dev enp4s0; done
```
# Trace utils
```sh
watch 'ss -uml' 
watch 'ss -ums'
watch 'cat /proc/interrupts | column -t'
```
