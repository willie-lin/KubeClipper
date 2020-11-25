package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/willie-lin/KubeClipper/pkg/csconfig"
	"github.com/willie-lin/KubeClipper/pkg/database/ent"
)

type Client struct {
	Ent *ent.Client
	CTX context.Context
	Log *log.Logger
}

func NewClient(config *csconfig.DatabaseCfg) (*Client, error) {
	var _ *ent.Client
	var err error
	if config == nil {
		return &Client{}, fmt.Errorf("DB config is empty")
	}
	switch config.Type {

	case "sqlite":
		_, err = ent.Open("sqlite3", fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1", config.DbPath))
		if err != nil {
			return &Client{}, fmt.Errorf("faild opening connection to sqlite: %v", err)
		}

	}
}
