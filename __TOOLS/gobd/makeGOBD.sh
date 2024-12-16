#!/bin/bash
rm go.mod
rm go.sum
rm /opt/SCRIPTS/go.mod
rm /opt/SCRIPTS/go.sum

LEGACY_GOBUILD_SCRIPT="./gobuild_STILL_NEEDED_for_GOBD_deploy.sh"

${LEGACY_GOBUILD_SCRIPT} --intel --d /opt/SCRIPTS -name gobd.mac
${LEGACY_GOBUILD_SCRIPT} --arm64 --d /opt/SCRIPTS -name gobd_ARM.mac

/opt/SCRIPTS/quickCommit.sh
