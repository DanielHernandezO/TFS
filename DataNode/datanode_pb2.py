# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datanode.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0e\x64\x61tanode.proto\"A\n\rChunkMetadata\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07\x63hunkId\x18\x02 \x01(\x05\x12\x11\n\treplicaId\x18\x03 \x01(\x05\"*\n\x0e\x44\x61taNodeSocket\x12\n\n\x02ip\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\t\"z\n\x11\x43hunkWriteRequest\x12%\n\rchunkMetadata\x18\x01 \x01(\x0b\x32\x0e.ChunkMetadata\x12*\n\x11pipelineDataNodes\x18\x02 \x03(\x0b\x32\x0f.DataNodeSocket\x12\x12\n\nchunkBytes\x18\x03 \x01(\x0c\"9\n\x10\x43hunkReadRequest\x12%\n\rchunkMetadata\x18\x01 \x01(\x0b\x32\x0e.ChunkMetadata\"Z\n\x12\x43hunkWriteResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12%\n\rchunkMetadata\x18\x03 \x01(\x0b\x32\x0e.ChunkMetadata\"m\n\x11\x43hunkReadResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12%\n\rchunkMetadata\x18\x03 \x01(\x0b\x32\x0e.ChunkMetadata\x12\x12\n\nchunkBytes\x18\x04 \x01(\x0c\x32}\n\x0f\x44\x61taNodeService\x12\x35\n\nWriteChunk\x12\x12.ChunkWriteRequest\x1a\x13.ChunkWriteResponse\x12\x33\n\tReadChunk\x12\x12.ChunkReadResponse\x1a\x12.ChunkReadResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'datanode_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_CHUNKMETADATA']._serialized_start=18
  _globals['_CHUNKMETADATA']._serialized_end=83
  _globals['_DATANODESOCKET']._serialized_start=85
  _globals['_DATANODESOCKET']._serialized_end=127
  _globals['_CHUNKWRITEREQUEST']._serialized_start=129
  _globals['_CHUNKWRITEREQUEST']._serialized_end=251
  _globals['_CHUNKREADREQUEST']._serialized_start=253
  _globals['_CHUNKREADREQUEST']._serialized_end=310
  _globals['_CHUNKWRITERESPONSE']._serialized_start=312
  _globals['_CHUNKWRITERESPONSE']._serialized_end=402
  _globals['_CHUNKREADRESPONSE']._serialized_start=404
  _globals['_CHUNKREADRESPONSE']._serialized_end=513
  _globals['_DATANODESERVICE']._serialized_start=515
  _globals['_DATANODESERVICE']._serialized_end=640
# @@protoc_insertion_point(module_scope)
