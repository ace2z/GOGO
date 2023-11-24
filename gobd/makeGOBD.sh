#!/bin/bash
rm go.mod
rm go.sum
rm /opt/SCRIPTS/go.mod
rm /opt/SCRIPTS/go.sum


gobuild.sh --intel --d /opt/SCRIPTS -name gobd.mac
gobuild.sh --arm64 --d /opt/SCRIPTS -name gobd_ARM.mac

/opt/SCRIPTS/quickCommit.sh
