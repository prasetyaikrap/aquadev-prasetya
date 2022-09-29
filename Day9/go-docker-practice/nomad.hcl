job "go-echo-prasetya" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "go-echo-prasetya" {
      driver = "docker"

      config {
        image = "prasetyaip/go-echo-prasetya:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "go-echo-prasetya"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.go-echo-prasetya-demo.rule=Host(\"prasetyaecho.cupang.efishery.ai\")",
      ]
      check {
        port        = "http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}