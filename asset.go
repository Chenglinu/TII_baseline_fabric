package main

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Define the Smart Contract structure
type DigitalAssetContract struct {
    contractapi.Contract
}

// User structure
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    // Add other user properties as needed
}

// Asset structure
type Asset struct {
    Hash       string `json:"hash"`
    OwnerID    string `json:"ownerId"`
    Authorized []string `json:"authorized"` // List of user IDs authorized to access the asset
}

// RegisterUser registers a new user
func (s *DigitalAssetContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID, userName string) error {
    user := User{
        ID:   userID,
        Name: userName,
    }
    userAsBytes, _ := json.Marshal(user)
    return ctx.GetStub().PutState(userID, userAsBytes)
}

// RegisterAsset registers a new digital asset
func (s *DigitalAssetContract) RegisterAsset(ctx contractapi.TransactionContextInterface, assetHash, ownerID string) error {
    asset := Asset{
        Hash:       assetHash,
        OwnerID:    ownerID,
        Authorized: []string{},
    }
    assetAsBytes, _ := json.Marshal(asset)
    return ctx.GetStub().PutState(assetHash, assetAsBytes)
}

// AuthorizeAsset allows an asset's owner to authorize another user to access it
func (s *DigitalAssetContract) AuthorizeAsset(ctx contractapi.TransactionContextInterface, assetHash, userID string) error {
    assetAsBytes, _ := ctx.GetStub().GetState(assetHash)
    if assetAsBytes == nil {
        return fmt.Errorf("Asset not found")
    }

    asset := Asset{}
    _ = json.Unmarshal(assetAsBytes, &asset)

    // Check if the requesting user is the owner
    if asset.OwnerID != userID {
        return fmt.Errorf("Only the asset owner can authorize access")
    }

    // Add user to the authorized list
    asset.Authorized = append(asset.Authorized, userID)

    updatedAssetAsBytes, _ := json.Marshal(asset)
    return ctx.GetStub().PutState(assetHash, updatedAssetAsBytes)
}

func main() {
    digitalAssetContract := new(DigitalAssetContract)
    chaincode, err := contractapi.NewChaincode(digitalAssetContract)
    if err != nil {
        fmt.Printf("Error create digital asset chaincode: %s", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting digital asset chaincode: %s", err.Error())
    }
}
