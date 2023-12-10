# Distributed Sys Project2

## How to run miner and client

First `make`, then `cd build`, run `./miner --addr_idx=xxx`, where "xxx" is the index of the node address in `addres_list` in `utils.go`. You should start miner on all the nodes in `addres_list`, then start client by `./client` on any nodes (or your local machine).

You will see the output to terminal on miner node.

For now, we fix the `addres_list=[]String{address[0],address[1]}`, if you want to add miners on other node, you should add the address of the node into the `addres_list` in `utils.go`.
