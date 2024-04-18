package main

import (
    "crypto"
   
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/json"
    "encoding/pem"
    "fmt"
   
    "encoding/hex" // 如果您在代码中使用了hex相关的功能
    "strings" // 如果您使用了strings包中的功能

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// User structure
type User struct {
    ID        string `json:"id"`
    PublicKey string `json:"publicKey"` // User's RSA public key
}

// Asset structure
type Asset struct {
    Hash       string   `json:"hash"`
    OwnerID    string   `json:"ownerId"`
    Authorized []string `json:"authorized"`
}

// DigitalAssetContract defines the smart contract structure
type DigitalAssetContract struct {
    contractapi.Contract
}

// RegisterUser registers a new user with their public key
func (s *DigitalAssetContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID, publicKey string) error {
    user := User{
        ID:        userID,
        PublicKey: publicKey,
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

// AuthorizeAsset allows an asset's owner to authorize another user to access it using a signature
// AuthorizeAsset allows an asset's owner to authorize another user to access it using a signature
func (s *DigitalAssetContract) AuthorizeAsset(ctx contractapi.TransactionContextInterface, ownerID, signatureHex, message string) error {
    // Get the owner user's public key
    userAsBytes, _ := ctx.GetStub().GetState(ownerID)
    if userAsBytes == nil {
        return fmt.Errorf("User not found")
    }
    user := User{}
    _ = json.Unmarshal(userAsBytes, &user)

    // Decode the public key
    pubKeyBytes, _ := pem.Decode([]byte(user.PublicKey))
    if pubKeyBytes == nil {
        return fmt.Errorf("Could not decode public key")
    }
    parsedKey, err := x509.ParsePKIXPublicKey(pubKeyBytes.Bytes)
    if err != nil {
        return fmt.Errorf("Could not parse public key: %s", err.Error())
    }
    publicKey, ok := parsedKey.(*rsa.PublicKey)
    if !ok {
        return fmt.Errorf("Public key is not of type RSA")
    }

    // Verify the signature
    hashed := sha256.Sum256([]byte(message))
    signatureBytes, _ := hex.DecodeString(signatureHex)
    err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signatureBytes)
    if err != nil {
        return fmt.Errorf("Signature verification failed: %s", err.Error())
    }

    // Parse the message for asset hash and the ID of the user to authorize
    parts := strings.Split(message, ":")
    if len(parts) != 2 {
        return fmt.Errorf("Invalid message format")
    }
    assetHash, authorizedUserID := parts[0], parts[1]

    // Proceed with authorization
    assetAsBytes, _ := ctx.GetStub().GetState(assetHash)
    if assetAsBytes == nil {
        return fmt.Errorf("Asset not found")
    }
    asset := Asset{}
    _ = json.Unmarshal(assetAsBytes, &asset)

    // Add the authorized user to the asset's authorized list
    asset.Authorized = append(asset.Authorized, authorizedUserID)
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
