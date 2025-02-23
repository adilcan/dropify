# Dropify - eBPF Packet Dropper

TODO: Forgot to add the bpf loader for the ((current)) main go lol.

Please note that this project is still work in progress, and it shouldn't be used in a production environment.

This project uses eBPF and Go to drop TODO packets on a specified network interface using Traffic Control (TC).

Note: This README is a draft, do not rely on it.

## Prerequisites
- Linux kernel with eBPF support
- Go 1.18+
- Clang
- XDP Tools

## Installation & Usage

### 1. Clone
```sh
git clone https://github.com/yourusername/ebpf-drop-packets.git
cd ebpf-drop-packets
```

### 2. Running
```sh
go run main.go -iface eth0
```
Replace eth0 with your interface.

### 3. Compiling the flowz
```
clang -O2 -g -target bpf -I/usr/include/asm-generic -I/usr/include/asm -I/usr/include/bpf -c flowz.c -o flowz.o
```

### 4. Attaching flowz to an interface
```
sudo ip link set dev eth0 xdp obj flowz.o
```
