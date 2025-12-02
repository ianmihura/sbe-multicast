# SBE Multicast tooling

Simple Binary Encode (SBE) parsing app that receives market data packets via multicast and prints them out to stdout.

Using Deribit (a crypto derivatives exchange) as example. Find the following page with relevant documentation https://support.deribit.com/hc/en-us/articles/29392445838877-Multicast-Developer-Guide.

# How to use

If your machine already supports multicast (and that's a big if), simply run `go run . -v` to see the program load a sample capture, send it via multicast and parse it. You can also compile for your system with `go build .`.

Available flags:

```
# go run . --help
  -h    Hex dump received network pkts
          (if not looping, app may end before printing full hex dump)
  -iface ip link
        Network interface.
          Check available ifaces with ip link (default "wlan0")
  -l    Loop: inifinite loop for pkt replay. Otherwise:
          ping: 10 pkts.
          sample: replay sample once
  -m    Monitoring network pkts (sent & received)
  -mode string
        Replay mode.
          Options: [ping, sample]. (default "sample")
  -p    Pretty-Print parsed SBE structs
```


# Architecture

We use goroutines to split the work up. Find boot implementation in `main.go` file.
1. Spinup the listener
    - Its only task is to receive the messages and relay them to the `dataCh` for workers to read
2. Spinup workers
    - We may have as many as procesors in the machine: `runtime.NumCPU()`
    - They do the bulk of the work: parsing from bytes stream to our stdmessage class, ready to be used
    - Relay the result via `syncCh`
3. Spinup sync goroutine
    - It will receive the finished work of the workers, via `syncCh`
    - Here is where we will print to stdout, monitor performance, sort packets and verify pkt drops
4. Lastly: spinup the packet replayer
    - It loads the sample pcap captures and sends them to an mcast group

### Monitoring performance

Monitoring e2e work (reception -> parsing -> sync):

```sh
# taskset -c 0 go run .
2025/11/29 14:09:02.126383 Listening on &{5 1500 wlan0 38:68:93:ea:87:0b up|broadcast|multicast|running} from 239.222.222.2:6200
2025/11/29 14:09:03.156624 Sending on 239.222.222.2:6200
2025/11/29 14:09:04.512231 Total Rcv: 10000 | PPS: 15428
2025/11/29 14:09:05.032443 Total Rcv: 20000 | PPS: 19225
2025/11/29 14:09:05.520770 Total Rcv: 30000 | PPS: 20480
2025/11/29 14:09:06.037357 Total Rcv: 40000 | PPS: 19359
2025/11/29 14:09:06.458361 Total Rcv: 50000 | PPS: 23755
```

# Additional tooling

- Find a list of usefull commands documented in `tools.md`
- (upcoming) Some articles telling how I built and tested this app

### Deribit Multicast Dev Guide
For updated pcapng captures and SBE classes, refer to [Deribit Dev Guide](https://support.deribit.com/hc/en-us/articles/29392445838877-Multicast-Developer-Guide)

### TODO explore further:
- sort pkts:
    - out-of-order
    - network gaps (drops)
- state interpretation:
    - mapping instrument_id : instrument_name
- state repr + recovery (protobuf?)
- snapshot + incremental replay
- multiple channels - ip:port listen to different asset bases
- perf:
    - pprof
        https://dev.to/jones_charles_ad50858dbc0/a-hands-on-guide-to-supercharging-your-go-apps-with-pprof-57m2
        https://go.dev/blog/pprof
        https://pkg.go.dev/net/http/pprof
    - size of buffs
    - kernel tuning
    - cpu affinity - monitor cpus and goroutines closer


// https://jewelhuq.medium.com/mastering-high-performance-tcp-udp-socket-programming-in-go-996dc85f5de1
// https://stackoverflow.com/questions/60337662/how-to-maximise-udp-packets-per-second-with-go
// https://blog.cloudflare.com/how-to-receive-a-million-packets/
// https://tungdam.medium.com/linux-network-ring-buffers-cea7ead0b8e8
// https://ntk148v.github.io/posts/linux-network-performance-ultimate-guide/
// https://balodeamit.blogspot.com/2013/10/receive-side-scaling-and-receive-packet.html
// https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/10/html/network_troubleshooting_and_performance_tuning/tuning-network-adapter-settings
// https://blog.packagecloud.io/monitoring-tuning-linux-networking-stack-receiving-data/

