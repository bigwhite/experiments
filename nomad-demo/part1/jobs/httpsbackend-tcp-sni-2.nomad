job "httpsbackend-sni-2" {
  datacenters = ["dc1"]
  type = "service"

  group "httpsbackend-sni-2" {
    count = 2 
    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "httpsbackend-sni-2" {
      driver = "docker"
      config {
        image = "bigwhite/httpsbackendservice:v1.0.1"
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
        name = "httpsbackend-sni-2"
	tags = ["urlprefix-mysite-sni-2.com/ proto=tcp+sni"]
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
