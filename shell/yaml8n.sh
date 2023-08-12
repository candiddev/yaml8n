#!/usr/bin/env bash

export APP_NAME=yaml8n
export GITHUB_REPOSITORY_ID=678385646
export INSTALL_ALL="install-go install-golangci-lint install-hugo install-shellcheck install-vault"

run-yaml8n-hugo () {
	"${DIR}/${BUILD_NAME_HOMECHART}" -c "${DIR}/homechart_config.yaml" run
}
