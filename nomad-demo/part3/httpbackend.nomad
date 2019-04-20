job "httpbackend" {
  datacenters = ["dc1"]
  type = "service"

  group "httpbackend" {
    count = 3

    task "httpbackend" {
      driver = "docker"
      config {
        image = "bigwhite/httpbackendservice:v1.0.0"
        dns_servers =  ["192.168.56.3", "192.168.56.4", "192.168.56.5"]
        network_mode = "weave"
        logging {
          type = "json-file"
        }
      }

      resources {
        network {
          mbits = 10
        }
      }

      service {
        name = "httpbackend"
      }
    }
  }
}
