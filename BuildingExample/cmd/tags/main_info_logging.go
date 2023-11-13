//go:build log_info
// +build log_info

package main

import log "github.com/sirupsen/logrus"

func init() {
	log.SetLevel(log.InfoLevel)
}
