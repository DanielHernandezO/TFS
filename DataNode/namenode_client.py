import grpc
import namenode_pb2_grpc 
import namenode_pb2 
from urllib.parse import urlparse,ParseResult
from threading import Thread
import time
class NameNodeClient:

    def __init__(self,namenode_uri, datanode_uri,datanode_slave_uri : str,) -> None:
        self.namenode_uri= namenode_uri
        self.datanode_slave_uri = datanode_slave_uri

        self.slave_channel = grpc.insecure_channel(self.datanode_slave_uri)
        self.channel : grpc.Channel = grpc.insecure_channel(self.namenode_uri)

        self.datanode_uri = datanode_uri
        self.heartbeat_stub : namenode_pb2_grpc.HeartBeatStub = namenode_pb2_grpc.HeartBeatStub(self.channel);
        uri_parse : ParseResult = urlparse("http://"+self.datanode_uri)
        self.socket = namenode_pb2.Socket(ip=uri_parse.hostname,port=str(uri_parse.port))
        
    def send_heartbeat(self):
        try:
            init_response : namenode_pb2.Response = self.heartbeat_stub.HeartBeat(self.socket)
            return init_response
        except grpc.RpcError as e:
            try:
                error_status = e.code()
                if(error_status == grpc.StatusCode.DEADLINE_EXCEEDED or error_status == grpc.StatusCode.UNAVAILABLE):
                    print(f"MASTER HEARBEAT FAILED, SWITCHING TO SLAVE AT: {self.datanode_slave_uri}")
                    slave_stub : namenode_pb2_grpc.HeartBeatStub = namenode_pb2_grpc.HeartBeatStub(self.slave_channel);
                    return slave_stub.HeartBeat(self.socket)
                else:
                    raise e
            except grpc.RpcError as slave_e:
                if(error_status == grpc.StatusCode.DEADLINE_EXCEEDED or error_status == grpc.StatusCode.UNAVAILABLE):
                    return namenode_pb2.Response(code="500",message="Slave and Master failed")
                else:
                    raise slave_e


            

    
    def start_heartbeat(self):

        tr = Thread(target=self.heartbeat_job)
        tr.start()

    

    def heartbeat_job(self):
        while True:
            res = self.send_heartbeat()
            if res.code != 200:
                raise Exception(f"Heartbeat was unsucessful with code {res.code}: {res.message}")
                exit(1)
            else:
                print(f"sent heartbeat: {res.code}")
            time.sleep(60)






    
        



