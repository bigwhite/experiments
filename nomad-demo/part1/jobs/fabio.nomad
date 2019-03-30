job "fabio" {
  datacenters = ["dc1"]
  type = "system"

  group "fabio" {
    task "fabio" {
      driver = "docker"
      config {
        image = "fabiolb/fabio"
        network_mode = "host"
	logging {
   	  type = "json-file"
	}
        args = [
	  "-proxy.addr=:9999;proto=http,:9997;proto=tcp,:9996;proto=tcp+sni",
	  "-log.level=TRACE",
	  "-log.access.target=stdout"
	]
      }

      resources {
        cpu    = 200
        memory = 128
        network {
          mbits = 20
        }
      }
    }
  }
}
