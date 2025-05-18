package main

import (
    "TAOSubnetAlert/internal/config"
    "TAOSubnetAlert/internal/twitter"
    "TAOSubnetAlert/internal/monitor"
)

func main() {
    cfg := config.LoadConfig()
    twClient := twitter.NewTwitterClient(cfg.TwitterServiceURL)
    monitor.MonitorSubnetPrice(cfg.SubnetID, cfg.APIKey, cfg.Threshold, cfg.Interval, twClient)
}