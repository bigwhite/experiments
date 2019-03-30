job "httpsbackend-sni-1" {
  datacenters = ["dc1"]
  type = "service"

  group "httpsbackend-sni-1" {
    count = 2 
    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "httpsbackend-sni-1" {
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
        name = "httpsbackend-sni-1"
	tags = ["urlprefix-mysite-sni-1.com/ proto=tcp+sni"]
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
