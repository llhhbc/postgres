#!/bin/bash

# Copyright The KubeDB Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -xeou pipefail

DOCKER_REGISTRY=${DOCKER_REGISTRY:-kubedb}

IMG=postgres-tools
SUFFIX=v3
DB_VERSION=9.6
PATCH=9.6.7

TAG="$DB_VERSION-$SUFFIX"
BASE_TAG="$PATCH-$SUFFIX"


docker pull "$DOCKER_REGISTRY/$IMG:$BASE_TAG"

docker tag "$DOCKER_REGISTRY/$IMG:$BASE_TAG" "$DOCKER_REGISTRY/$IMG:$TAG"
docker push "$DOCKER_REGISTRY/$IMG:$TAG"
