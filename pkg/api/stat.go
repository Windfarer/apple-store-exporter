package api

import (
	"context"
	"github.com/Windfarer/powerup-follower-exporter/pkg/models"
	"net/http"
	"net/url"
)

func (c *Client) GetUidStat(ctx context.Context, uid string) (*models.Stat, error) {
	res := models.StatResponse{}
	err := c.do(ctx, http.MethodGet, "x/relation/stat", url.Values{"vmid": []string{uid}}, nil, &res)
	return &res.Data, err
}
