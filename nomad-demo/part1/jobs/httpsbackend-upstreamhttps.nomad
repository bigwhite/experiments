job "httpsbackend" {
  datacenters = ["dc1"]
  type = "service"

  group "httpsbackend" {
    count = 2 
    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "httpsbackend" {
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
        name = "httpsbackend"
	tags = ["urlprefix-mysite-https.com:9999/ proto=https tlsskipverify=true"]
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
