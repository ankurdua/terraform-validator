resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv-policy"
  location = "us-central1"
  project  = "{{.Provider.project}}"

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

data "google_iam_policy" "foo" {
  binding {
    role = "viewer"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloud_run_service_iam_policy" "foo" {
  location = google_cloud_run_service.default.location
  project = google_cloud_run_service.default.project
  service = google_cloud_run_service.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}