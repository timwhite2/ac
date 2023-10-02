package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func generateBitcoinAddresses(num int) ([]string, error) {
	addresses := make([]string, num)

	for i := 0; i < num; i++ {
		address, err := generateBitcoinAddress()
		if err != nil {
			return nil, err
		}
		addresses[i] = address
	}

	return addresses, nil
}

func generateBitcoinAddress() (string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", err
	}

	pubKey := privKey.PubKey()
	address, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}

func writeAddressesToFile(filename string, addresses []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, address := range addresses {
		_, err := file.WriteString(fmt.Sprintf("%s\n", address))
		if err != nil {
			return err
		}
	}

	return nil
}

func GenBtcAddr(num int) {
	filename := fmt.Sprintf("files/addresses_%d.txt", num)
	addresses, err := generateBitcoinAddresses(num)
	if err != nil {
		log.Fatal(err)
	}

	err = writeAddressesToFile(filename, addresses)
	if err != nil {
		log.Fatal(err)
	}
}
