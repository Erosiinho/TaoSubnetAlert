package main

import (
    "TAOSubnetAlert/internal/config"
    "TAOSubnetAlert/internal/twitter"
    "TAOSubnetAlert/internal/monitor"
    "TAOSubnetAlert/internal/discord"
)

func main() {
    cfg := config.LoadConfig()
    if cfg.Service == "Twitter" {
        twClient := twitter.NewTwitterClient(cfg.TwitterServiceURL)
        monitor.MonitorSubnetsWithTwitter(cfg.SubnetIDs, cfg.APIKey, cfg.Threshold, cfg.Interval, twClient)
    } else {
        discordClient := discord.NewDiscordClient(cfg.DiscordServiceURL)
        monitor.MonitorSubnetsWithDiscord(cfg.SubnetIDs, cfg.APIKey, cfg.Threshold, cfg.Interval, discordClient)
    }
}