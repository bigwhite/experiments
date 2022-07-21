#include <stdio.h>
#include <unistd.h>
#include <sys/resource.h>
#include <bpf/libbpf.h>
#include <linux/bpf.h>
#include "execve_counter.skel.h"

typedef __u64 u64;
typedef char stringkey[64];

static int libbpf_print_fn(enum libbpf_print_level level, const char *format, va_list args)
{
	return vfprintf(stderr, format, args);
}

int main(int argc, char **argv)
{
	struct execve_counter_bpf *skel;
	int err;

	libbpf_set_strict_mode(LIBBPF_STRICT_ALL);
	/* Set up libbpf errors and debug info callback */
	libbpf_set_print(libbpf_print_fn);

	/* Open BPF application */
	skel = execve_counter_bpf__open();
	if (!skel) {
		fprintf(stderr, "Failed to open BPF skeleton\n");
		return 1;
	}

	/* Load & verify BPF programs */
	err = execve_counter_bpf__load(skel);
	if (err) {
		fprintf(stderr, "Failed to load and verify BPF skeleton\n");
		goto cleanup;
	}

	/* init the counter */
	stringkey key = "execve_counter";
	u64 v = 0;
	err = bpf_map__update_elem(skel->maps.execve_counter, &key, sizeof(key), &v, sizeof(v), BPF_ANY);
	if (err != 0) {
		fprintf(stderr, "Failed to init the counter, %d\n", err);
		goto cleanup;
	}

	/* Attach tracepoint handler */
	err = execve_counter_bpf__attach(skel);
	if (err) {
		fprintf(stderr, "Failed to attach BPF skeleton\n");
		goto cleanup;
	}

	for (;;) {
			// read counter value from map
			//
			//LIBBPF_API int bpf_map__lookup_elem(const struct bpf_map *map,
            //        const void *key, size_t key_sz,
            //        void *value, size_t value_sz, __u64 flags);
			//        /usr/local/bpf/include/bpf/libbpf.h
			err = bpf_map__lookup_elem(skel->maps.execve_counter, &key, sizeof(key), &v, sizeof(v), BPF_ANY);
			if (err != 0) {
               fprintf(stderr, "Lookup key from map error: %d\n", err);
               goto cleanup;
			} else {
			   printf("execve_counter is %llu\n", v);
			}
			
			sleep(5);
	}

cleanup:
	execve_counter_bpf__destroy(skel);
	return -err;
}
