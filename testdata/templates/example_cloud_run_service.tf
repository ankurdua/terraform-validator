# provider "google" {
#   project="scottsuarez-graphite"
# }

resource "google_cloud_run_service" "default" {
  name     = "cloudrun-to-get-cai"
  location = "us-central1"
  project="scottsuarez-graphite"

  metadata {
  namespace = "scottsuarez-graphite"
  annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        args  = ["arrgs"]
        ports {
          container_port = 8080
        }
      }
	  container_concurrency = 10
	  timeout_seconds = 600
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}