#!/bin/sh
# Installs the quota-events ILM policy, ingest pipeline and data stream template.
# Usage: setup.sh [ES_URL]   (default http://localhost:9200)
set -eu
ES="${1:-http://localhost:9200}"
DIR="$(cd "$(dirname "$0")" && pwd)"

put() {
  printf '%s ... ' "$1"
  curl -fsS -X PUT "$ES/$1" -H 'Content-Type: application/json' --data-binary "@$DIR/$2"
  echo
}

put _ilm/policy/quota-events        ilm-policy.json
put _ingest/pipeline/quota-events   pipeline.json
put _index_template/quota-events    index-template.json
echo "elasticsearch setup done"
