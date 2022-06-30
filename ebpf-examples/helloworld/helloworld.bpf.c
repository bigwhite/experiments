#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

SEC("tracepoint/syscalls/sys_enter_execve")

int bpf_prog(void *ctx) {
  char msg[] = "Hello, World!";
  bpf_printk("invoke bpf_prog: %s\n", msg);
  return 0;
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";
