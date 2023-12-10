# Distributed Sys Project2

## How to run miner and client

First `make`, then `cd build`, run `./miner --addr_idx=xxx`, where "xxx" is the index of the node address in `addres_list` in `utils.go`. You should start miner on all the nodes in `addres_list`, then start client by `./client` on any nodes (or your local machine).
You will see the output to terminal on miner node.
