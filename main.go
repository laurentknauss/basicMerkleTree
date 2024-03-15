
package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"log"
)



// Defining a struct named 'Node' representing a node in the Merkle Trie with 3 fields. 
type Node struct { 
	Hash string // stores the hash value of the node 's data . 
	Left *Node  // Pointer to the left child node .
	Right *Node // Pointer to the right node . 
}

// NewNode creates a new Node with calculated hash . 
// Given 2 child nodes , it creates a new 'Node' instance representing their parent node . 
// 
func NewNode(left, right *Node) *Node { 
	node := &Node{
		Left: left, 
		Right: right,
	}
	node.Hash = node.calculateHash() 
	return node 
}

// calculateHash computes the hash for the node combining child hashes . 
func(n *Node) calculateHash() string {
	hashes := "" // this initializes an empty string variable named `hashes`, this variable will be used to concatenate the hashes of child nodes .

	// This below block checks if the `Left` child node of the current node (`n`) is not nil, indicating it exists.
	if n.Left !=nil { 
	// If the left child node exists, its hash is concatenated to the `hashes` string.
		hashes += n.Left.Hash
	}

	// This below block checks if the `Right` child node of the current node (`n`) is not nil, indicating it exits.
	if n.Right != nil { 
	// If the right child node exists, its hash is concatenated  to the `hashes` string . 
		hashes += n.Right.Hash
	}

	// Computing the SHA-256 hash of the `hashes` string. `sha256.Sum256` takesa byte slice as input and returns an array of 32 bytes representing the hash.
	hash := sha256.Sum256([]byte(hashes))
	// After computing the hash, the function checks idf the length of the hash is zero, which would indicate an error during hash computation
	if len(hash) == 0 { 
		// If such error is detected , the function logs an error message using `log.Fatal()` and terminates the program.
		log.Fatal("Error: unable to compute hash for node.")
	}
	// This converts the byte array `hash` to a hexadecimal string representation using 'hex.encodeToString' for better human-readibility and further console printing.
	// and returns it as the result of the `calculateHash` function . 
	return hex.EncodeToString(hash[:])
}


// LeafNode is a function used to create a leaf node in the Merkle Tree . 
// The function takesa single poarameter data of type string and returns a pointer to the `Node` struct .  
func LeafNode(data string) *Node { 
	hash := sha256.Sum256([]byte(data))
	return &Node{
		Hash: hex.EncodeToString(hash[:]),
}
}

func main() {
	// Example usage: 
	// Constructing a small Merkle Trie for demonstration . 

	// Creating 4 leaf nodes with dat  "data1" , "data2", "data3", "data4" 
	leaf1 := LeafNode("data1")
	leaf2 := LeafNode("data2") 
	leaf3 := LeafNode("data3") 
	leaf4 := LeafNode("data4") 

	// Creating 2 middle nodes by combining the leaf nodes . 
	middleNode1 := NewNode(leaf1,leaf2)
	middleNode2 := NewNode(leaf3,leaf4)  

	// Creating the root node by combining the middle nodes 
	rootNode := NewNode(middleNode1, middleNode2) 

	// Printing the root hash of the constructed Merkle Trie.
	fmt.Println("Root Hash:", rootNode.Hash) 
	// Output: rootHash: <some value> 
}