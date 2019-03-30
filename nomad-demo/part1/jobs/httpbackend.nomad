job "httpbackend" {
  datacenters = ["dc1"]
  type = "service"

  group "httpbackend" {
    count = 2 

    task "httpbackend" {
      driver = "docker"
      config {
        image = "bigwhite/httpbackendservice:v1.0.0"
        port_map {
          http = 8081
        }
	logging {
	  type = "json-file"
	}
      }

      resources {
        network {
          mbits = 10
          port "http" {}
        }
      }

      service {
        name = "httpbackend"
        port = "http"
      }
    }
  }
}
