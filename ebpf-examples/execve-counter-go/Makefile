CLANG ?= clang-10
CFLAGS ?= -O2 -g -Wall -Werror

LIBEBPF_TOP = /home/tonybai/go/src/github.com/cilium/ebpf
EXAMPLES_HEADERS = $(LIBEBPF_TOP)/examples/headers

all: generate

generate: export BPF_CLANG=$(CLANG)
generate: export BPF_CFLAGS=$(CFLAGS)
generate: export BPF_HEADERS=$(EXAMPLES_HEADERS)
generate:
	go generate ./...
