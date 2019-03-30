job "httpsbackend-tcp" {
  datacenters = ["dc1"]
  type = "service"

  group "httpsbackend-tcp" {
    count = 2 
    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "httpsbackend-tcp" {
      driver = "docker"
      config {
        image = "bigwhite/httpsbackendservice:v1.0.0"
        port_map {
          https = 7777
        }
	logging {
	  type = "json-file"
	}
      }

      resources {
        network {
          mbits = 10
          port "https" {}
        }
      }

      service {
        name = "httpsbackend-tcp"
	tags = ["urlprefix-:9997 proto=tcp"]
        port = "https"
	check {
          name     = "alive"
          type     = "tcp"
          path     = "/"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }
}
