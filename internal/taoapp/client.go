package taoapp

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Subnet struct {
    Netuid              int     `json:"netuid"`
    NetworkRegisteredAt int     `json:"network_registered_at"`
    OwnerColdkey        string  `json:"owner_coldkey"`
    OwnerHotkey         string  `json:"owner_hotkey"`
    Tempo               int     `json:"tempo"`
    Price              float64 `json:"price"`
    TaoIn              float64 `json:"tao_in"`
    AlphaIn            float64 `json:"alpha_in"`
    AlphaOut           float64 `json:"alpha_out"`
    SubnetName         string  `json:"subnet_name"`
    GithubRepo         string  `json:"github_repo"`
    SubnetContact      string  `json:"subnet_contact"`
    SubnetUrl          string  `json:"subnet_url"`
    SubnetWebsite      string  `json:"subnet_website"`
    Discord            string  `json:"discord"`
    Additional         string  `json:"additional"`
    Symbol             string  `json:"symbol"`
    RootProp           float64 `json:"root_prop"`
}

func FetchSubnetsPrices(netuids []int, apiKey string) (map[int]float64, error) {
    url := "https://api.tao.app/api/beta/analytics/subnets/info"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("X-API-KEY", apiKey)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var subnets []Subnet
    if err := json.Unmarshal(body, &subnets); err != nil {
        return nil, err
    }

    desiredNetuids := make(map[int]bool)
    for _, id := range netuids {
        desiredNetuids[id] = true
    }

    result := make(map[int]float64)
    for _, subnet := range subnets {
        if desiredNetuids[subnet.Netuid] {
            result[subnet.Netuid] = subnet.Price
        }
    }

    if len(result) != len(netuids) {
        var missing []int
        for _, id := range netuids {
            if _, found := result[id]; !found {
                missing = append(missing, id)
            }
        }
        return nil, fmt.Errorf("certains subnets n'ont pas été trouvés: %v", missing)
    }

    return result, nil
}