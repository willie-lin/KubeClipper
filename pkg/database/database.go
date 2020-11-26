package database

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/willie-lin/KubeClipper/pkg/csconfig"
	"github.com/willie-lin/KubeClipper/pkg/database/ent"
	"github.com/willie-lin/KubeClipper/pkg/types"
)

type Client struct {
	Ent *ent.Client
	CTX context.Context
	Log *log.Logger
}

func NewClient(config *csconfig.DatabaseCfg) (*Client, error) {
	var client *ent.Client
	var err error
	if config == nil {
		return &Client{}, fmt.Errorf("DB config is empty")
	}
	switch config.Type {

	case "sqlite":
		client, err = ent.Open("sqlite3", fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1", config.DbPath))
		if err != nil {
			return &Client{}, fmt.Errorf("faild opening connection to sqlite: %v", err)
		}
	case "mysql":
		client, err = ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", config.User, config.Password,
			config.Host, config.Port, config.DbName))
		if err != nil {
			return &Client{}, fmt.Errorf("failed opening connection to mysql: %v", err)
		}
	case "postgres", "postgresql":
		client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			config.Host, config.Port, config.User, config.DbName, config.Password))
		if err != nil {
			return &Client{}, fmt.Errorf("failed opening connection postgres: %v", err)
		}
	default:
		return &Client{}, fmt.Errorf("unknown database type")

	}
	/*The logger that will be used by db operations*/
	clog := log.New()
	if err := types.ConfigureLogger(clog); err != nil {
		return nil, errors.Wrap(err, "while configuring db logger")
	}

	if config.LogLevel != nil {
		clog.SetLevel(*config.LogLevel)
		if *config.LogLevel >= log.TraceLevel {
			log.Debugf("Enabling request debug")
			client = client.Debug()
		}
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		//if err = client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed createing scheam resource: %v", err)
	}
	return &Client{client, context.Background(), clog}, nil
}

//func (c *Client) StartFlushScheduler(config *csconfig.FlushDBCfg) (*gocron.Scheduler, error)  {
//	//maxItems := 0
//	//maxAge := ""
//	//if config.MaxItems != nil && *config.MaxItems <= 0 {
//	//	return nil, fmt.Errorf("max_items can't be zero or negative number")
//	//
//	//}
//	//if config.MaxItems != nil {
//	//		maxItems = *config.MaxItems
//	//}
//	//if config.MaxAge != nil && *config.MaxAge != "" {
//	//		maxAge = *config.MaxAge
//	//}
//	//
//	//schedler := gocron.NewScheduler(time.UTC)
//	//schedler.Every(1).Minute().Do(c.FlushAlerts, maxAge, maxItems)
//	//schedler.StartAsync()
//	//return schedler, nil
//
//}
