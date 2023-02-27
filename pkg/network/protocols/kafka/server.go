// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package kafka

import (
	"regexp"
	"testing"
	"time"

	"github.com/DataDog/datadog-agent/pkg/network/protocols/http/testutil"
	protocolsUtils "github.com/DataDog/datadog-agent/pkg/network/protocols/testutil"
)

func RunServer(t *testing.T, serverAddr, serverPort string) {
	env := []string{
		"KAFKA_ADDR=" + serverAddr,
		"KAFKA_PORT=" + serverPort,
	}

	t.Helper()
	dir, _ := testutil.CurDir()
	protocolsUtils.RunDockerServer(t, "kafka", dir+"/testdata/docker-compose.yml", env, regexp.MustCompile(".*Started socket server acceptors and processors.*"), 10*time.Minute)
}