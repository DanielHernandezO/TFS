from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ChunkMetadata(_message.Message):
    __slots__ = ("name", "chunkId", "replicaId")
    NAME_FIELD_NUMBER: _ClassVar[int]
    CHUNKID_FIELD_NUMBER: _ClassVar[int]
    REPLICAID_FIELD_NUMBER: _ClassVar[int]
    name: str
    chunkId: int
    replicaId: int
    def __init__(self, name: _Optional[str] = ..., chunkId: _Optional[int] = ..., replicaId: _Optional[int] = ...) -> None: ...

class DataNodeSocket(_message.Message):
    __slots__ = ("ip", "port")
    IP_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    ip: str
    port: str
    def __init__(self, ip: _Optional[str] = ..., port: _Optional[str] = ...) -> None: ...

class ChunkWriteRequest(_message.Message):
    __slots__ = ("chunkMetadata", "pipelineDataNodes", "chunkBytes")
    CHUNKMETADATA_FIELD_NUMBER: _ClassVar[int]
    PIPELINEDATANODES_FIELD_NUMBER: _ClassVar[int]
    CHUNKBYTES_FIELD_NUMBER: _ClassVar[int]
    chunkMetadata: ChunkMetadata
    pipelineDataNodes: _containers.RepeatedCompositeFieldContainer[DataNodeSocket]
    chunkBytes: bytes
    def __init__(self, chunkMetadata: _Optional[_Union[ChunkMetadata, _Mapping]] = ..., pipelineDataNodes: _Optional[_Iterable[_Union[DataNodeSocket, _Mapping]]] = ..., chunkBytes: _Optional[bytes] = ...) -> None: ...

class ChunkReadRequest(_message.Message):
    __slots__ = ("chunkMetadata",)
    CHUNKMETADATA_FIELD_NUMBER: _ClassVar[int]
    chunkMetadata: ChunkMetadata
    def __init__(self, chunkMetadata: _Optional[_Union[ChunkMetadata, _Mapping]] = ...) -> None: ...

class ChunkWriteResponse(_message.Message):
    __slots__ = ("code", "message", "chunkMetadata")
    CODE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CHUNKMETADATA_FIELD_NUMBER: _ClassVar[int]
    code: int
    message: str
    chunkMetadata: ChunkMetadata
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ..., chunkMetadata: _Optional[_Union[ChunkMetadata, _Mapping]] = ...) -> None: ...

class ChunkReadResponse(_message.Message):
    __slots__ = ("code", "message", "chunkMetadata", "chunkBytes")
    CODE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CHUNKMETADATA_FIELD_NUMBER: _ClassVar[int]
    CHUNKBYTES_FIELD_NUMBER: _ClassVar[int]
    code: int
    message: str
    chunkMetadata: ChunkMetadata
    chunkBytes: bytes
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ..., chunkMetadata: _Optional[_Union[ChunkMetadata, _Mapping]] = ..., chunkBytes: _Optional[bytes] = ...) -> None: ...
