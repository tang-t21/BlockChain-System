# Distributed Sys Project2

**Final report is in `blockchain_v2/report/`.**
`blockchain_v2` contains the code of part2 version.

## How to run miner and client in Part1

First `make`, then `cd build`, run `./miner --addr_idx=xxx`, where "xxx" is the index of the node address in `address_list` in `utils.go`. You should start miner on all the nodes in `address_list`, then start client by `./client` on any nodes (or your local machine).

You will see the output to terminal on miner node.

For now, we fix the `address_list=[]String{address[0],address[1]}`, if you want to add miners on other node, you should add the address of the node into the `address_list` in `utils.go`.

