job "httpbackend" {
  datacenters = ["dc1"]
  type = "service"

  group "httpbackend" {
    count = 2 
    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

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
	tags = ["urlprefix-mysite.com:9999/"]
        port = "http"
      }
    }
  }
}
