# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/configuration.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='common/configuration.proto',
  package='common',
  syntax='proto3',
  serialized_pb=_b('\n\x1a\x63ommon/configuration.proto\x12\x06\x63ommon\" \n\x10HashingAlgorithm\x12\x0c\n\x04name\x18\x01 \x01(\t\"*\n\x19\x42lockDataHashingStructure\x12\r\n\x05width\x18\x01 \x01(\r\"%\n\x10OrdererAddresses\x12\x11\n\taddresses\x18\x01 \x03(\t\"\x1a\n\nConsortium\x12\x0c\n\x04name\x18\x01 \x01(\tBS\n$org.hyperledger.fabric.protos.commonZ+github.com/ledgerone/fabric-ledgerone/protos/commonb\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_HASHINGALGORITHM = _descriptor.Descriptor(
  name='HashingAlgorithm',
  full_name='common.HashingAlgorithm',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='common.HashingAlgorithm.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=38,
  serialized_end=70,
)


_BLOCKDATAHASHINGSTRUCTURE = _descriptor.Descriptor(
  name='BlockDataHashingStructure',
  full_name='common.BlockDataHashingStructure',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='width', full_name='common.BlockDataHashingStructure.width', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=72,
  serialized_end=114,
)


_ORDERERADDRESSES = _descriptor.Descriptor(
  name='OrdererAddresses',
  full_name='common.OrdererAddresses',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='addresses', full_name='common.OrdererAddresses.addresses', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=116,
  serialized_end=153,
)


_CONSORTIUM = _descriptor.Descriptor(
  name='Consortium',
  full_name='common.Consortium',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='common.Consortium.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=155,
  serialized_end=181,
)

DESCRIPTOR.message_types_by_name['HashingAlgorithm'] = _HASHINGALGORITHM
DESCRIPTOR.message_types_by_name['BlockDataHashingStructure'] = _BLOCKDATAHASHINGSTRUCTURE
DESCRIPTOR.message_types_by_name['OrdererAddresses'] = _ORDERERADDRESSES
DESCRIPTOR.message_types_by_name['Consortium'] = _CONSORTIUM

HashingAlgorithm = _reflection.GeneratedProtocolMessageType('HashingAlgorithm', (_message.Message,), dict(
  DESCRIPTOR = _HASHINGALGORITHM,
  __module__ = 'common.configuration_pb2'
  # @@protoc_insertion_point(class_scope:common.HashingAlgorithm)
  ))
_sym_db.RegisterMessage(HashingAlgorithm)

BlockDataHashingStructure = _reflection.GeneratedProtocolMessageType('BlockDataHashingStructure', (_message.Message,), dict(
  DESCRIPTOR = _BLOCKDATAHASHINGSTRUCTURE,
  __module__ = 'common.configuration_pb2'
  # @@protoc_insertion_point(class_scope:common.BlockDataHashingStructure)
  ))
_sym_db.RegisterMessage(BlockDataHashingStructure)

OrdererAddresses = _reflection.GeneratedProtocolMessageType('OrdererAddresses', (_message.Message,), dict(
  DESCRIPTOR = _ORDERERADDRESSES,
  __module__ = 'common.configuration_pb2'
  # @@protoc_insertion_point(class_scope:common.OrdererAddresses)
  ))
_sym_db.RegisterMessage(OrdererAddresses)

Consortium = _reflection.GeneratedProtocolMessageType('Consortium', (_message.Message,), dict(
  DESCRIPTOR = _CONSORTIUM,
  __module__ = 'common.configuration_pb2'
  # @@protoc_insertion_point(class_scope:common.Consortium)
  ))
_sym_db.RegisterMessage(Consortium)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n$org.hyperledger.fabric.protos.commonZ+github.com/ledgerone/fabric-ledgerone/protos/common'))
try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
