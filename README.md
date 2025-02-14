# Dropify - eBPF Packet Dropper

Please note that this project is still work in progress, and it shouldn't be used in a production environment.

This project uses eBPF and Go to drop all packets on a specified network interface using Traffic Control (TC).

## Prerequisites
- Linux kernel with eBPF support
- Go 1.18+

## Installation & Usage

### 1. Clone the repo
```sh
git clone https://github.com/yourusername/ebpf-drop-packets.git
cd ebpf-drop-packets
```

### 2. Running
```sh
go run main.go -iface eth0
```

Replace eth0 with your interface.
