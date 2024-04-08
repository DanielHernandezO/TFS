from concurrent import futures
import os
from dotenv import load_dotenv
import grpc
from namenode_client import NameNodeClient
import threading
import datanode_pb2_grpc
from datanode_servicer import DataNodeServicer
def initialize():
    if os.path.exists(".env"):
        load_dotenv()
    
    MASTER_NAMENODE_URI = os.getenv("MASTER_NAMENODE_URI")
    SLAVE_NAMENODE_URI = os.getenv("SLAVE_NAMENODE_URI")
    DATANODE_URI = os.getenv("DATANODE_URI")

    print(MASTER_NAMENODE_URI,SLAVE_NAMENODE_URI,DATANODE_URI)

    namenode_client = NameNodeClient(MASTER_NAMENODE_URI,DATANODE_URI,SLAVE_NAMENODE_URI)

    init_response = namenode_client.send_heartbeat()

    if(init_response.code == 200):
        print(f"Sucessfully connected with Namenode at {MASTER_NAMENODE_URI} : {init_response.message}")
    else:
        print(f"Request failed {init_response.message}")

    
    namenode_client.start_heartbeat()


    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10),options=[
    ('grpc.max_send_message_length', 50 * 1024 * 1024),  # 50 MB
    ('grpc.max_receive_message_length', 50 * 1024 * 1024)  # 50 MB
])

    datanode_pb2_grpc.add_DataNodeServiceServicer_to_server(DataNodeServicer(MASTER_NAMENODE_URI,DATANODE_URI,SLAVE_NAMENODE_URI),server)

    datanode_port = DATANODE_URI.split(":")[1]
    port = server.add_insecure_port(f"0.0.0.0:{datanode_port}")

    server.start()

    print(f"DataNode grpc started server at  0.0.0.0:{port}")

    server.wait_for_termination()

    


    

    

    
    

    



    
        
    

def main():
    initialize()



if __name__=='__main__':
    main()

