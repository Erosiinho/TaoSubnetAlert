package monitor

import (
    "log"
    "time"
    "TAOSubnetAlert/internal/taoapp"
    "TAOSubnetAlert/internal/twitter"
)

func MonitorSubnetPrice(netuid int, apiKey string, threshold float64, intervalMinutes int, client *twitter.Client) {
    lastPrice, err := taoapp.FetchSubnetPrice(netuid, apiKey)
    if err != nil {
        log.Fatalf("Erreur prix initial subnet %d : %v", netuid, err)
    }
    log.Printf("Prix initial subnet %d : %.4f $TAO", netuid, lastPrice)

    ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
    defer ticker.Stop()

    for {
        <-ticker.C

        currentPrice, err := taoapp.FetchSubnetPrice(netuid, apiKey)
        if err != nil {
            log.Printf("Erreur récupération prix : %v", err)
            continue
        }

        variation := ((currentPrice - lastPrice) / lastPrice) * 100

        if variation >= threshold {
            twitter.TweetAlert(netuid, currentPrice, variation, true, client)
            lastPrice = currentPrice
        } else if variation <= -threshold {
            twitter.TweetAlert(netuid, currentPrice, -variation, false, client)
            lastPrice = currentPrice
        } else {
            log.Printf("Pas de variation suffisante : %.2f%% (seuil %.2f%%)", variation, threshold)
        }
    }
}