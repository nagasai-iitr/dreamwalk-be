package main

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"math/rand"
//	"os"
//	"time"
//
//	solana "github.com/gagliardetto/solana-go"
//	"github.com/gagliardetto/solana-go/rpc"
//)
//
//func triggerSmartContract(min, max float64) bool {
//	// Load ENV
//	programID := solana.MustPublicKeyFromBase58(os.Getenv("SOLANA_PROGRAM_ID"))
//	receiver := solana.MustPublicKeyFromBase58(os.Getenv("SOLANA_RECEIVER"))
//	cluster := os.Getenv("SOLANA_CLUSTER")
//	payerKeyPath := os.Getenv("SOLANA_PAYER_KEY")
//
//	// Load Payer Keypair
//	payer, err := loadKeypair(payerKeyPath)
//	if err != nil {
//		log.Printf("Keypair load error: %v", err)
//		return false
//	}
//
//	client := rpc.New(cluster)
//
//	// Random Amount [min, max]
//	rand.Seed(time.Now().UnixNano())
//	lamports := uint64((min + rand.Float64()*(max-min)) * 1e9)
//
//	// Build transaction (here just send lamports; customize if calling instructions)
//	tx, err := solana.NewTransaction(
//		[]solana.Instruction{
//			solana.NewSystemInstructionTransfer(payer.PublicKey(), receiver, lamports).Build(),
//		},
//		solana.NewAccountMetaSlice([]solana.AccountMeta{}),
//	)
//	if err != nil {
//		log.Printf("Tx creation error: %v", err)
//		return false
//	}
//
//	// Get recent blockhash
//	recent, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
//	if err != nil {
//		log.Printf("Blockhash error: %v", err)
//		return false
//	}
//
//	tx.SetRecentBlockHash(recent.Value.Blockhash)
//	tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
//		if key.Equals(payer.PublicKey()) {
//			return &payer.PrivateKey
//		}
//		return nil
//	})
//
//	// Send Transaction
//	sig, err := client.SendTransactionWithOpts(
//		context.Background(),
//		tx,
//		false,
//		rpc.CommitmentFinalized,
//	)
//	if err != nil {
//		log.Printf("Transaction failed: %v", err)
//		return false
//	}
//
//	log.Printf("Transaction submitted: %s", sig.String())
//	return true
//}
//
//func loadKeypair(path string) (*solana.Keypair, error) {
//	bytes, err := ioutil.ReadFile(path)
//	if err != nil {
//		return nil, err
//	}
//	var keyBytes []byte
//	if err := json.Unmarshal(bytes, &keyBytes); err != nil {
//		return nil, err
//	}
//	return solana.PrivateKeyFromBytes(keyBytes)
//}
