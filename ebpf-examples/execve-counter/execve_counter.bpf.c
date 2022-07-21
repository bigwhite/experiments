#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

typedef __u64 u64;
typedef char stringkey[64];

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 128);
    //__type(key, stringkey);
	stringkey* key;
    __type(value, u64);
} execve_counter SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_execve")
int bpf_prog(void *ctx) {
  stringkey key = "execve_counter";
  u64 *v = NULL;
  v = bpf_map_lookup_elem(&execve_counter, &key);
  if (v != NULL) {
    *v += 1;
    //bpf_map_update_elem(&execve_counter, &key, v, BPF_ANY);
    //bpf_printk("map value: %d\n", *v);
  }
  return 0;
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";
