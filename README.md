# go-bitlight
A bitcoin light client

It is a pruned version of the bitcoin blockchain. The client only stores the block headers and supports this one query:
1. Given a raw transaction and a merkle path, verifies if the transaction exists in the Bitcoin blockchain.
