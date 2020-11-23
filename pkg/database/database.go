package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/willie-lin/KubeClipper/pkg/database/ent"
)

type Client struct {
	Ent *ent.Client
	CTX context.Context
	Log *log.Logger
}
