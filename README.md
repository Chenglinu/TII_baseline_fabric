Required:
go

Ubuntu environment

hyperledger frabic docker

If want to extend this Test Network（add peer nodes or add org）, extend fabric key with cryptogen, update system-channel definition, compose orderer nodes, compose peer nodes, join application channel and deploy the chaincode.

cd this filefolder

./network.sh up createChannel

peer lifecycle chaincode package basic.tar.gz --path ./MedicalAsset/  --lang golang --label basic_1.0

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode install basic.tar.gz

export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

peer lifecycle chaincode install basic.tar.gz

peer lifecycle chaincode queryinstalled

export CC_PACKAGE_ID=basic_1.0:The ID list before

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json

peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"

检查成功部署：
peer lifecycle chaincode querycommitted --channelID mychannel --name basic --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"




用户注册测试
User register test data：

用户A注册密钥：User_A_SK
MIIEowIBAAKCAQEAuyf3PpoMTEi+Lqr3TqjnvuohvB/olY95qrGwAO4WRkUPdgnr
hRxbUrlwJvAHVkYAE8D934hFmZeD99q5hkpIq++3dEE4P1knw+c2V+LFN+ch9oEg
7+R1d5RXsQjhowuRp/QqaKivRbo3cS4/ANGqX+6imj7BlwBks2JWJcxyHwqya5of
GafAeMkvFCwwT6l9+M7jcCoTWK1JduJAt1mCHBlobqhXyWXiwcKK/6q4W4BjyPD+
EI4TTt5sYzOL0L2kE3LzeLNfNML8r7b5KCen7hQXr/wrfdZXo2Mk6KaQxrt5XWXm
E+3ZREWduQKjA2CEiU69fDCb/K4NW3PtLSezdwIDAQABAoIBAAtPBumJdWTGoHdB
bWAbZKVskE1FwFAJM1jVE8V6nW3xjlWbk9efNwVwnw47QrY71JVc+/odydbWCOtZ
FAzBQFLjUTp9FmD9iiGUPvxgf3o5RRwYAV19eHuZQxM3birj8BEt98ILL0wPTHpS
SQxLnvfc+4ZGdHwjUfJk5r+x8tNdP5vV2PLIMpLWdxtJ44f4sSo38XMeiBBQoKG4
ijLqDZ6EFFGg6c6jkfeA6F4Lo1seYUTw4bJBFn/F/aBSain9ff7oGY2jBtoz87jP
Nev9zFXJ/R4lDdwICEOf5aL7+cjVa1RSfOVKc5uaRiyddtNQKiTJrUAdt8hPVHre
lyZzi0ECgYEAzsIpUkZBdU2aS0IaCeI+p3+UGXZphAvWf2ZuWByK9tUJmPK//v5K
R6VrlAcdb514JjQgFhGRdneBiDz6IHZuGwJBsqH9cnG8K5K1hefP/3EF/87GpR3S
duEblKAhsJrywZEVSAmce3A9UjVMSf0cgPriZKa0N3BowLqjapx/jO0CgYEA57qr
9zFzZNS/koLhGFfwtjiC9LPvaqfRCgqwKg/uDpo3cw72mwUzW0zEnQz61X5OzzWE
c/JlCFt388YHF82phjJqHMhzb8kNKDUfhzbK92JFPiRNDlIwE7j1xNk5D3CejtQ2
rhZBRq5qcVcAFLUij+gqBGtGGKa0Ln1yxXixWXMCgYAfMexKHY+Cw1KkSDTliN0r
KHSP4u3InoCeeDXt1WCiHUJ1cSrGrldGuA6jJu+qB3g5S2QL8FqiJSXGCG00uKmk
KZMAALDcs4xQhrIcof0f7U2aavhNsVIv3Ybrxb1PiBFYYyty0wBpH2YhISmBgE7s
pu3BgeVu9+bWLVK6oyIbgQKBgQCDDJ/KGS5APMzmh5vTD5CzDLyKtOPWNnfSrP65
mu6vVWm8aR7vxn4nyP98LeYBLQBW0NZKWC/pDQmGVFyYipevq/00r+wQsOv+7CQb
bNJWGz47iX4GdlZ4IObk31AUukUBU2RlCXL7DRJnbKNAErwsFEkG3L/0mvpVPF7W
7I2nnwKBgH/Vd4h/VVwGFsCpPSDsLebRK2wf3DYYEE9iwzAA9N7J/OkBzZD62AaR
UQqUe1jLM9QBcNumag5n2y7VdhJ/NcuitH6/7h+xuD+ZbUdR94uOvUWourQI9tC1
VneuCl5+vXhiiG0XSChPiSbbnOHgl8IJ9RpyxSLF06mLrAKKW9R3

公钥：
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuyf3PpoMTEi+Lqr3Tqjn
vuohvB/olY95qrGwAO4WRkUPdgnrhRxbUrlwJvAHVkYAE8D934hFmZeD99q5hkpI
q++3dEE4P1knw+c2V+LFN+ch9oEg7+R1d5RXsQjhowuRp/QqaKivRbo3cS4/ANGq
X+6imj7BlwBks2JWJcxyHwqya5ofGafAeMkvFCwwT6l9+M7jcCoTWK1JduJAt1mC
HBlobqhXyWXiwcKK/6q4W4BjyPD+EI4TTt5sYzOL0L2kE3LzeLNfNML8r7b5KCen
7hQXr/wrfdZXo2Mk6KaQxrt5XWXmE+3ZREWduQKjA2CEiU69fDCb/K4NW3PtLSez
dwIDAQAB

用户注册
User register
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"RegisterUser","Args":["Alice","MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuyf3PpoMTEi+Lqr3Tqjn
vuohvB/olY95qrGwAO4WRkUPdgnrhRxbUrlwJvAHVkYAE8D934hFmZeD99q5hkpI
q++3dEE4P1knw+c2V+LFN+ch9oEg7+R1d5RXsQjhowuRp/QqaKivRbo3cS4/ANGq
X+6imj7BlwBks2JWJcxyHwqya5ofGafAeMkvFCwwT6l9+M7jcCoTWK1JduJAt1mC
HBlobqhXyWXiwcKK/6q4W4BjyPD+EI4TTt5sYzOL0L2kE3LzeLNfNML8r7b5KCen
7hQXr/wrfdZXo2Mk6KaQxrt5XWXmE+3ZREWduQKjA2CEiU69fDCb/K4NW3PtLSez
dwIDAQAB"]}'


资产注册
Asset register
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peer --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"RegisterAsset","Args":["asset1Hash","Alice"]}'


资产授权
Asset Authorize
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"AuthorizeAsset","Args":["Alice","signatureHex(use User_A_SK)","asset1Hash:Bob"]}'

