package metric

import (
	"context"
	"log"
	"time"

	"github.com/Windfarer/powerup-follower-exporter/pkg/api"
	"github.com/Windfarer/powerup-follower-exporter/pkg/config"
)

func MonitorStatFunc(ctx context.Context, client *api.Client, uids []string) {
	for _, uid := range uids {
		log.Printf("fetch %s", uid)

		res, err := client.GetUidStat(ctx, uid)
		if err != nil {
			log.Printf("fetch stat for %s failed, detail: %s", uid, err)
			continue
		}
		followerMetric.WithLabelValues(uid).Set(float64(res.Follower))
	}
}

func StartMonitFulfillment(ctx context.Context, client *api.Client, config *config.Config) {
	timer := time.NewTimer(0)

	go func() {
		for {
			select {
			case <-timer.C:
				MonitorStatFunc(ctx, client, config.Uids)
				timer.Reset(config.Interval)
			case <-ctx.Done():
				return
			}
		}
	}()
}
