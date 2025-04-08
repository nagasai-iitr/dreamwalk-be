package main

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/gagliardetto/solana-go/rpc"
//
//	//"fmt"
//	"io/ioutil"
//	"log"
//	"math/rand"
//	"os"
//	"time"
//
//	solana "github.com/gagliardetto/solana-go"
//	//"github.com/gagliardetto/solana-go/rpc"
//)
//
//func triggerAnchorTransfer(min, max float64, receiverPubKey string) bool {
//	rand.Seed(time.Now().UnixNano())
//	amount := uint64((min + rand.Float64()*(max-min)) * 1e9)
//
//	cluster := os.Getenv("SOLANA_CLUSTER")
//	programID := solana.MustPublicKeyFromBase58(os.Getenv("SOLANA_PROGRAM_ID"))
//	receiver := solana.MustPublicKeyFromBase58(receiverPubKey)
//	authority, err := loadKeypair(os.Getenv("SOLANA_PAYER_KEY"))
//	if err != nil {
//		log.Printf("Keypair error: %v", err)
//		return false
//	}
//
//	client := rpc.New(cluster)
//
//	// Derive vault PDA
//	vault, bump, err := solana.FindProgramAddress([][]byte{[]byte("vault")}, programID)
//	if err != nil {
//		log.Printf("PDA derive error: %v", err)
//		return false
//	}
//	log.Printf("Vault PDA: %s (bump: %d)", vault.String(), bump)
//
//	// Build Anchor-compatible instruction: transfer(amount)
//	discriminator := []byte{2} // Instruction 2 = transfer
//	amtBytes := make([]byte, 8)
//	for i := 0; i < 8; i++ {
//		amtBytes[i] = byte(amount >> (8 * i))
//	}
//	data := append(discriminator, amtBytes...)
//
//	accounts := solana.AccountMetaSlice{
//		{PublicKey: vault, IsWritable: true, IsSigner: false},
//		{PublicKey: receiver, IsWritable: true, IsSigner: false},
//		{PublicKey: authority.PublicKey(), IsWritable: false, IsSigner: true},
//		{PublicKey: solana.SystemProgramID, IsWritable: false, IsSigner: false},
//	}
//
//	instruction := solana.NewInstruction(programID, accounts, data)
//
//	// Get recent blockhash
//	recentBlockhashResp, err := client.GetLatestBlockhash(context.Background())
//	if err != nil {
//		log.Printf("blockhash error: %v", err)
//		return false
//	}
//	blockhash := recentBlockhashResp.Value.Blockhash
//
//	// Create transaction
//	tx, err := solana.NewTransaction(
//		[]solana.Instruction{instruction},
//		blockhash,
//		solana.TransactionPayer(authority.PublicKey()),
//	)
//	if err != nil {
//		log.Printf("Tx build error: %v", err)
//		return false
//	}
//
//	// Sign transaction
//	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
//		if key.Equals(authority.PublicKey()) {
//			return &authority.PrivateKey
//		}
//		return nil
//	})
//	if err != nil {
//		log.Printf("Signing error: %v", err)
//		return false
//	}
//
//	// Send transaction
//	sig, err := client.SendTransactionWithOpts(context.Background(), tx, rpc.TransactionOpts{
//		SkipPreflight:       false,
//		PreflightCommitment: rpc.CommitmentFinalized,
//	})
//	if err != nil {
//		log.Printf("Send error: %v", err)
//		return false
//	}
//
//	log.Printf("✅ Transfer sent! Signature: %s", sig.String())
//	return true
//}
//
//func loadKeypair(path string) (*solana.Keypair, error) {
//	bytes, err := ioutil.ReadFile(path)
//	if err != nil {
//		return nil, err
//	}
//	var key []byte
//	if err := json.Unmarshal(bytes, &key); err != nil {
//		return nil, err
//	}
//	return solana.PrivateKeyFromBytes(key)
//}
