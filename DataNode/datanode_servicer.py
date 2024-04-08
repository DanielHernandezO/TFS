from urllib.parse import ParseResult, urlparse
import grpc
from concurrent import futures
import datanode_pb2,datanode_pb2_grpc
import namenode_pb2_grpc,namenode_pb2
class DataNodeServicer(datanode_pb2_grpc.DataNodeServiceServicer):

    def __init__(self,namenode_uri:str,datanode_uri:str,namenode_slave_uri:str) -> None:
        super().__init__()

        
        self.datanode_uri : ParseResult = urlparse("http://"+datanode_uri)
        self.datanode_slave_uri = namenode_slave_uri
        print(self.datanode_uri,self.datanode_uri.hostname,self.datanode_uri.port)
        self.namenode_uri : ParseResult = urlparse("http://"+namenode_uri)
        namenode_channel = grpc.insecure_channel(namenode_uri)
        slave_channel = grpc.insecure_channel(namenode_slave_uri)

        self.namenode_stub = namenode_pb2_grpc.LocateChunkStub(namenode_channel)
        self.namenode_slave_stub = namenode_pb2_grpc.LocateChunkStub(slave_channel)

    def WriteChunk(self, request:datanode_pb2.ChunkWriteRequest, context):


        chunk_metadata = request.chunkMetadata

        filename = f"{chunk_metadata.name}_{chunk_metadata.chunkId}_{chunk_metadata.replicaId}.dat"
        with open(f"chunks/{filename}","wb+") as file:

            print(f"writing {filename}")
            file.write(request.chunkBytes)




        response = datanode_pb2.ChunkWriteResponse(
            code=200,
            message="Chunk successfully written",
            chunkMetadata=request.chunkMetadata
            )
        
        locate_chunk_request = namenode_pb2.ChunkLocation(name=response.chunkMetadata.name,chunkId=response.chunkMetadata.chunkId,replicaId=response.chunkMetadata.replicaId,socket=namenode_pb2.Socket(ip=self.datanode_uri.hostname,port=str(self.datanode_uri.port)))
        
        print(locate_chunk_request)

        try:
            self.namenode_stub.LocateChunk(locate_chunk_request)
        except grpc.RpcError as e:
            try:
                error_status = e.code()
                if(error_status == grpc.StatusCode.DEADLINE_EXCEEDED or error_status == grpc.StatusCode.UNAVAILABLE):
                    print(f"MASTER WRITE REQUEST FAILED, SWITCHING TO SLAVE AT: {self.datanode_slave_uri}")
                    self.namenode_slave_stub.LocateChunk(locate_chunk_request)
                else:
                    raise e
            except grpc.RpcError as slave_e:
                if(error_status == grpc.StatusCode.DEADLINE_EXCEEDED or error_status == grpc.StatusCode.UNAVAILABLE):
                    print("BOTH MASTER AND SLAVE FAILED")
                raise slave_e

        if request.pipelineDataNodes != []:

            next_datanode = request.pipelineDataNodes[0]

            next_datanode_chunk_metadata:datanode_pb2.ChunkMetadata= datanode_pb2.ChunkMetadata()
            
            next_datanode_chunk_metadata.CopyFrom(request.chunkMetadata)

            next_datanode_chunk_metadata.replicaId = request.chunkMetadata.replicaId + 1

            datanode_channel = grpc.insecure_channel(f"{next_datanode.ip}:{next_datanode.port}")
            datanode_stub : datanode_pb2_grpc.DataNodeServiceStub = datanode_pb2_grpc.DataNodeServiceStub(datanode_channel)
            
            next_datanode_request = datanode_pb2.ChunkWriteRequest(chunkMetadata=next_datanode_chunk_metadata,pipelineDataNodes=request.pipelineDataNodes[1:],chunkBytes=request.chunkBytes)
            
            next_datanode_res = datanode_stub.WriteChunk(next_datanode_request)

            if next_datanode_res.code == 200:
                print(next_datanode_res)



        return response

    def ReadChunk(self, request:datanode_pb2.ChunkWriteRequest, context):
        
        
        chunk_metadata = request.chunkMetadata

        filename = f"{chunk_metadata.name}_{chunk_metadata.chunkId}_{chunk_metadata.replicaId}.dat"
        chunkBytes = None
        with open(f"chunks/{filename}","rb") as chunk:
            chunkBytes = chunk.read()

        response = datanode_pb2.ChunkReadResponse(
            code=200,
            message="Chunk successfully read",
            chunkMetadata=request.chunkMetadata,
            chunkBytes=chunkBytes
        )
        return response
