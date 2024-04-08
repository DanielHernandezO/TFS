from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Socket(_message.Message):
    __slots__ = ("ip", "port")
    IP_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    ip: str
    port: str
    def __init__(self, ip: _Optional[str] = ..., port: _Optional[str] = ...) -> None: ...

class ChunkLocation(_message.Message):
    __slots__ = ("name", "chunkId", "replicaId", "socket")
    NAME_FIELD_NUMBER: _ClassVar[int]
    CHUNKID_FIELD_NUMBER: _ClassVar[int]
    REPLICAID_FIELD_NUMBER: _ClassVar[int]
    SOCKET_FIELD_NUMBER: _ClassVar[int]
    name: str
    chunkId: int
    replicaId: int
    socket: Socket
    def __init__(self, name: _Optional[str] = ..., chunkId: _Optional[int] = ..., replicaId: _Optional[int] = ..., socket: _Optional[_Union[Socket, _Mapping]] = ...) -> None: ...

class Sockets(_message.Message):
    __slots__ = ("sockets",)
    SOCKETS_FIELD_NUMBER: _ClassVar[int]
    sockets: _containers.RepeatedCompositeFieldContainer[Socket]
    def __init__(self, sockets: _Optional[_Iterable[_Union[Socket, _Mapping]]] = ...) -> None: ...

class Chunk(_message.Message):
    __slots__ = ("chunkId", "sockets")
    CHUNKID_FIELD_NUMBER: _ClassVar[int]
    SOCKETS_FIELD_NUMBER: _ClassVar[int]
    chunkId: int
    sockets: _containers.RepeatedCompositeFieldContainer[Socket]
    def __init__(self, chunkId: _Optional[int] = ..., sockets: _Optional[_Iterable[_Union[Socket, _Mapping]]] = ...) -> None: ...

class FileMetadata(_message.Message):
    __slots__ = ("fileId", "chunks")
    FILEID_FIELD_NUMBER: _ClassVar[int]
    CHUNKS_FIELD_NUMBER: _ClassVar[int]
    fileId: str
    chunks: _containers.RepeatedCompositeFieldContainer[Chunk]
    def __init__(self, fileId: _Optional[str] = ..., chunks: _Optional[_Iterable[_Union[Chunk, _Mapping]]] = ...) -> None: ...

class Metadata(_message.Message):
    __slots__ = ("files",)
    FILES_FIELD_NUMBER: _ClassVar[int]
    files: _containers.RepeatedCompositeFieldContainer[FileMetadata]
    def __init__(self, files: _Optional[_Iterable[_Union[FileMetadata, _Mapping]]] = ...) -> None: ...

class VoidRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class Response(_message.Message):
    __slots__ = ("code", "message")
    CODE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    code: int
    message: str
    def __init__(self, code: _Optional[int] = ..., message: _Optional[str] = ...) -> None: ...
