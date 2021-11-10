resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv-binding"
  location = "us-central1"
  project  = "scottsuarez-graphite"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_binding" "foo" {
  location = google_cloud_run_service.default.location
  project = google_cloud_run_service.default.project
  service = google_cloud_run_service.default.name
  role = "viewer"
  members = ["user:admin@hashicorptest.com"]
}