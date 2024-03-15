


# Merkle Tree Implementation in Go

This repository contains an implementation of a Merkle Tree data structure in Go.

## Overview

A Merkle Tree is a binary tree in which each leaf node is a hash of a data block, and each non-leaf node is a hash of its child nodes. It is commonly used in peer-to-peer networks and distributed systems for efficient verification of large datasets.

This implementation includes:

- `Node` struct: Represents a node in the Merkle Tree with fields for storing hash value and pointers to left and right child nodes.
- `NewNode` function: Creates a new node by combining two child nodes and computing the hash.
- `calculateHash` method: Computes the hash for a node by combining the hashes of its child nodes.
- `LeafNode` function: Creates a leaf node with a hash value calculated from input data.
- Example usage in `main` function: Demonstrates the construction of a small Merkle Tree and prints the root hash.
