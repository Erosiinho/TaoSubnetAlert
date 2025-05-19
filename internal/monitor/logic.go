package monitor

import (
    "log"
    "time"
    "TAOSubnetAlert/internal/taoapp"
    "TAOSubnetAlert/internal/twitter"
    "TAOSubnetAlert/internal/discord"
)

func MonitorSubnetsWithDiscord(netuids []int, apiKey string, threshold float64, intervalMinutes int, client *discord.Client) {
    lastPrices, err := taoapp.FetchSubnetsPrices(netuids, apiKey)
    if err != nil {
        log.Fatalf("Erreur initiale de récupération des prix: %v", err)
    }
    
    for netuid, price := range lastPrices {
        log.Printf("Prix initial subnet %d: %.4f $TAO", netuid, price)
    }

    ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
    defer ticker.Stop()

    for {
        <-ticker.C
        currentPrices, err := taoapp.FetchSubnetsPrices(netuids, apiKey)
        if err != nil {
            log.Printf("Erreur de récupération des prix: %v", err)
            continue
        }
        for netuid, currentPrice := range currentPrices {
            lastPrice := lastPrices[netuid]
            variation := ((currentPrice - lastPrice) / lastPrice) * 100

            if variation >= threshold {
                discord.SendAlert(netuid, currentPrice, variation, true, client)
                lastPrices[netuid] = currentPrice
            } else if variation <= -threshold {
                discord.SendAlert(netuid, currentPrice, -variation, false, client)
                lastPrices[netuid] = currentPrice
            } else {
                log.Printf("Subnet %d - Variation: %.2f%% (seuil %.2f%%)", 
                    netuid, variation, threshold)
            }
        }
    }
}

func MonitorSubnetsWithTwitter(netuids []int, apiKey string, threshold float64, intervalMinutes int, client *twitter.Client) {
    lastPrices, err := taoapp.FetchSubnetsPrices(netuids, apiKey)
    if err != nil {
        log.Fatalf("Erreur initiale de récupération des prix: %v", err)
    }
    
    for netuid, price := range lastPrices {
        log.Printf("Prix initial subnet %d: %.4f $TAO", netuid, price)
    }

    ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
    defer ticker.Stop()

    for {
        <-ticker.C
        currentPrices, err := taoapp.FetchSubnetsPrices(netuids, apiKey)
        if err != nil {
            log.Printf("Erreur de récupération des prix: %v", err)
            continue
        }
        for netuid, currentPrice := range currentPrices {
            lastPrice := lastPrices[netuid]
            variation := ((currentPrice - lastPrice) / lastPrice) * 100

            if variation >= threshold {
                twitter.TweetAlert(netuid, currentPrice, variation, true, client)
                lastPrices[netuid] = currentPrice
            } else if variation <= -threshold {
                twitter.TweetAlert(netuid, currentPrice, -variation, false, client)
                lastPrices[netuid] = currentPrice
            } else {
                log.Printf("Subnet %d - Variation: %.2f%% (seuil %.2f%%)", 
                    netuid, variation, threshold)
            }
        }
    }
}