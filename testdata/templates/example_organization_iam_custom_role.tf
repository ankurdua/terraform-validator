/**
 * Copyright 2021 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> {{.Provider.version}}"
    }
  }
}

provider "google" {
{{if .Provider.credentials}}credentials = "{{.Provider.credentials}}"{{end}}
}

resource "google_organization_iam_custom_role" "my-custom-role" {
org_id = "{{.OrgID}}"
role_id = "myCustomRole"
title = "My Custom Role"
description = "A description"
permissions = ["iam.roles.list", "iam.roles.create", "iam.roles.delete"]
}