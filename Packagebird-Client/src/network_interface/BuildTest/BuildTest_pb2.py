# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: BuildTest.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='BuildTest.proto',
  package='BuildTest',
  syntax='proto3',
  serialized_options=b'Z\013./buildtest',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0f\x42uildTest.proto\x12\tBuildTest\",\n\x0bPackageInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\x03\",\n\tBuildInfo\x12\x0e\n\x06status\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t2J\n\x11\x42uildTestServices\x12\x35\n\x05\x42uild\x12\x16.BuildTest.PackageInfo\x1a\x14.BuildTest.BuildInfoB\rZ\x0b./buildtestb\x06proto3'
)




_PACKAGEINFO = _descriptor.Descriptor(
  name='PackageInfo',
  full_name='BuildTest.PackageInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='BuildTest.PackageInfo.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='BuildTest.PackageInfo.version', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=30,
  serialized_end=74,
)


_BUILDINFO = _descriptor.Descriptor(
  name='BuildInfo',
  full_name='BuildTest.BuildInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='BuildTest.BuildInfo.status', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='BuildTest.BuildInfo.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=76,
  serialized_end=120,
)

DESCRIPTOR.message_types_by_name['PackageInfo'] = _PACKAGEINFO
DESCRIPTOR.message_types_by_name['BuildInfo'] = _BUILDINFO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

PackageInfo = _reflection.GeneratedProtocolMessageType('PackageInfo', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGEINFO,
  '__module__' : 'BuildTest_pb2'
  # @@protoc_insertion_point(class_scope:BuildTest.PackageInfo)
  })
_sym_db.RegisterMessage(PackageInfo)

BuildInfo = _reflection.GeneratedProtocolMessageType('BuildInfo', (_message.Message,), {
  'DESCRIPTOR' : _BUILDINFO,
  '__module__' : 'BuildTest_pb2'
  # @@protoc_insertion_point(class_scope:BuildTest.BuildInfo)
  })
_sym_db.RegisterMessage(BuildInfo)


DESCRIPTOR._options = None

_BUILDTESTSERVICES = _descriptor.ServiceDescriptor(
  name='BuildTestServices',
  full_name='BuildTest.BuildTestServices',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=122,
  serialized_end=196,
  methods=[
  _descriptor.MethodDescriptor(
    name='Build',
    full_name='BuildTest.BuildTestServices.Build',
    index=0,
    containing_service=None,
    input_type=_PACKAGEINFO,
    output_type=_BUILDINFO,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_BUILDTESTSERVICES)

DESCRIPTOR.services_by_name['BuildTestServices'] = _BUILDTESTSERVICES

# @@protoc_insertion_point(module_scope)
