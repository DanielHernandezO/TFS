import grpc
import namenode_pb2_grpc 
import namenode_pb2 
from urllib.parse import urlparse,ParseResult
from threading import Thread
import time
class NameNodeClient:

    def __init__(self,namenode_uri, datanode_uri : str,) -> None:
        self.namenode_uri= namenode_uri
        self.channel : grpc.Channel = grpc.insecure_channel(self.namenode_uri)
        self.datanode_uri = datanode_uri
        self.heartbeat_stub : namenode_pb2_grpc.HeartBeatStub = namenode_pb2_grpc.HeartBeatStub(self.channel);
        uri_parse : ParseResult = urlparse(self.datanode_uri)
        self.socket = namenode_pb2.Socket(ip=uri_parse.hostname,port=str(uri_parse.port))
        
    def send_heartbeat(self):
        init_response : namenode_pb2.Response = self.heartbeat_stub.HeartBeat(self.socket)
        return init_response
    
    def start_heartbeat(self):

        tr = Thread(target=self.heartbeat_job)
        tr.start()

    

    def heartbeat_job(self):
        while True:
            res = self.send_heartbeat()
            if res.code != 200:
                raise Exception(f"Heartbeat was unsucessful with code {res.code}: {res.message}")
            else:
                print(f"sent heartbeat: {res.code}")
            time.sleep(60)






    
        



