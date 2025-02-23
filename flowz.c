#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct flow_key {
    __u32 src_ip;
    __u32 dst_ip;
    __u16 src_port;
    __u16 dst_port;
    __u8  proto;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 10240);
    __type(key, struct flow_key);
    __type(value, __u64);
} flows SEC(".maps");

SEC("xdp")
int count_flow(struct __sk_buff *skb) {
    struct flow_key key = {0};
    __u64 *count, one = 1;

    // Extract flow data (for simplicity, using a basic 5-tuple: src_ip, dst_ip, src_port, dst_port, protocol)
    key.src_ip = skb->remote_ip4;
    key.dst_ip = skb->local_ip4;
    key.src_port = skb->local_port;
    key.dst_port = skb->remote_port;
    key.proto = skb->protocol;

    count = bpf_map_lookup_elem(&flows, &key);
    if (count)
        __sync_fetch_and_add(count, 1);
    else
        bpf_map_update_elem(&flows, &key, &one, BPF_ANY);

    return XDP_PASS;
}

char _license[] SEC("license") = "Apache";
