package utils

import "encoding/json"

func GetMintPayload() []byte {
    payload := map[string]string{
        "name":    "Test Sword",
        "address": "0xABC123",
    }
    jsonBytes, _ := json.Marshal(payload)
    return jsonBytes
}

func GetTransferPayload() []byte {
    payload := map[string]string{
        "name":    "Test Shield",
        "address": "0xDEF456",
    }
    jsonBytes, _ := json.Marshal(payload)
    return jsonBytes
}
