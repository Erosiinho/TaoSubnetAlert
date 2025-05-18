package taoapp

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)


func FetchSubnetPrice(netuid int, apiKey string) (float64, error) {
    url := fmt.Sprintf("https://api.tao.app/api/beta/analytics/subnets/info/%d", netuid)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, err
    }

    req.Header.Set("X-API-KEY", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        return 0, err
    }

    price, ok := data["price"].(float64)
    if !ok {
        return 0, fmt.Errorf("prix invalide")
    }

    return price, nil
}