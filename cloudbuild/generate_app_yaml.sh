#!/bin/bash

cat > ./app.yaml <<EOF
runtime: go116
service: pdf-api

automatic_scaling:
    max_concurrent_requests: 80
    min_pending_latency: 100ms
    max_pending_latency: 200ms
    target_throughput_utilization: 0.95
instance_class: F4

handlers:
- url: /.*
  secure: always
  script: auto

env_variables:
  GCLOUD_STORAGE_BUCKET: '${PROJECT_ID}.appspot.com'

EOF
