data_dir = "/root/nomad-install/nomad.d"

bind_addr = "192.168.56.3"

server {
  enabled = true
  bootstrap_expect = 3
}

client {
  enabled = true
}
