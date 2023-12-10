import matplotlib.pyplot as plt
import numpy as np
import subprocess
import time 

difficulties=[i for i in range()]


address = ["10.1.0.91:18051", "10.1.0.92:18052", "10.1.0.93:18053", "10.1.0.94:18054", "10.1.0.95:18055", "10.1.0.96:18056", "10.1.0.98:18057", "10.1.0.99:18058", "10.1.0.116:18059", "10.1.0.102:18060", "10.1.0.103:18061", "10.1.0.104:18062", "10.1.0.105:18063", "10.1.0.107:18064", "10.1.0.109:18065", "10.1.0.110:18066", "10.1.0.111:18067", "10.1.0.112:18068", "10.1.0.113:18069", "10.1.0.115:18070"]
ports = [8051, 8052, 8053, 8054, 8055, 8056, 8057, 8058, 8059, 8060, 8061, 8062, 8063, 8064, 8065, 8066, 8067, 8068, 8069, 8070]


def test_100blocks():
    miner_num=5
    for i in range(1,miner_num+1):
        server=subprocess.Popen([f"ssh -p {ports[i]} osgroup13@122.200.68.26 'cd ~/go/src/blockchain && ./miner --addr_idx={i}'"],shell=True)

    time.sleep(5)
    subprocess.call([f"./client"], shell=True)

    for i in range(1,miner_num+1):
            arg=f'"./miner --addr_idx={i}"'
            subprocess.call([f"ssh -p {ports[i]} osgroup13@122.200.68.26 'cd ~/go/src/p22 && ./kill.sh {arg}'"],shell=True)


def test_difficulty(): 
    plt.figure()
    plt.xlabel("number of leading zeros")
    plt.ylabel("mining speed(/s)")
    
    miner_num=5
    for num_zeros in difficulties:
         # FIXME: Pass the num_zeors to miner to adjust difficulty
        for i in range(1,miner_num+1):
            server=subprocess.Popen([f"ssh -p {ports[i]} osgroup13@122.200.68.26 'cd ~/go/src/blockchain && ./miner --addr_idx={i}'"],shell=True)

        time.sleep(5)
        subprocess.call([f"./client"], shell=True)

        for i in range(1,miner_num+1):
                arg=f'"./miner --addr_idx={i}"'
                subprocess.call([f"ssh -p {ports[i]} osgroup13@122.200.68.26 'cd ~/go/src/p22 && ./kill.sh {arg}'"],shell=True)

    
def test_fork():
    miner_num=5
    for i in range(1,miner_num):
        server=subprocess.Popen([f"ssh -p {ports[i]} osgroup13@122.200.68.26 'cd ~/go/src/blockchain && ./miner --addr_idx={i}'"],shell=True)
    
    time.sleep(15)
    server=subprocess.Popen([f"ssh -p {ports[1]} osgroup13@122.200.68.26 'cd ~/go/src/blockchain && ./miner --addr_idx={miner_num}'"],shell=True)

    time.sleep(10)

    for i in range(1,miner_num+1):
        arg=f'"./miner --addr_idx={i}"'
        subprocess.call([f"ssh -p {ports[i]} osgroup13@122.200.68.26 'cd ~/go/src/p22 && ./kill.sh {arg}'"],shell=True)


if __name__=='__main__':
    
    pass