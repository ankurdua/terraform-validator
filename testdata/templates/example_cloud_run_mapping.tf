resource "google_cloud_run_service" "default" {
    name     = "tf-test-cloudrun-srv-beep"
    location = "us-central1"
    project  = "{{.Provider.project}}"

    metadata {
      namespace = "{{.Provider.project}}"
    }

    template {
      spec {
        containers {
          image = "us-docker.pkg.dev/cloudrun/container/hello"
        }
      }
    }
  }

resource "google_cloud_run_domain_mapping" "default" {
  location = "us-central1"
  name     = "tf-test-domain-meep.gcp.tfacc.hashicorptest.com"
  project  = "{{.Provider.project}}"

  metadata {
    namespace = "{{.Provider.project}}"
  }

  spec {
    route_name = google_cloud_run_service.default.name
  }
}