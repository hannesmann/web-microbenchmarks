# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rservice.proto\x12\tbenchmark\"\x17\n\x07Request\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t\"\x18\n\x08Response\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t2J\n\x10\x42\x65nchmarkService\x12\x36\n\tBenchmark\x12\x12.benchmark.Request\x1a\x13.benchmark.Response\"\x00\x42\x12Z\x10./go/grpc/commonb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\020./go/grpc/common'
  _globals['_REQUEST']._serialized_start=28
  _globals['_REQUEST']._serialized_end=51
  _globals['_RESPONSE']._serialized_start=53
  _globals['_RESPONSE']._serialized_end=77
  _globals['_BENCHMARKSERVICE']._serialized_start=79
  _globals['_BENCHMARKSERVICE']._serialized_end=153
# @@protoc_insertion_point(module_scope)