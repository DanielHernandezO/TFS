import os
from dotenv import load_dotenv
import grpc
from namenode_client import NameNodeClient
import threading
def initialize():
    if os.path.exists(".env"):
        load_dotenv()
    
    MASTER_NAMENODE_URI = os.getenv("MASTER_NAMENODE_URI")
    SLAVE_NAMENODE_URI = os.getenv("SLAVE_NAMENODE_URI")
    DATANODE_URI = os.getenv("DATANODE_URI")

    print(MASTER_NAMENODE_URI,SLAVE_NAMENODE_URI,DATANODE_URI)

    namenode_client = NameNodeClient(MASTER_NAMENODE_URI,DATANODE_URI)

    init_response = namenode_client.send_heartbeat()

    if(init_response.code == 200):
        print(f"Sucessfully connected with Namenode at {MASTER_NAMENODE_URI} : {init_response.message}")
    else:
        print(f"Request failed {init_response.message}")

    
    namenode_client.start_heartbeat()

    


    

    

    
    

    



    
        
    

def main():
    initialize()



if __name__=='__main__':
    main()

