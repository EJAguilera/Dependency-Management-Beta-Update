# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: FileTransfer.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='FileTransfer.proto',
  package='FileTransfer',
  syntax='proto3',
  serialized_options=b'Z\016./FileTransfer',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12\x46ileTransfer.proto\x12\x0c\x46ileTransfer\"2\n\x04\x46ile\x12\x0e\n\x04name\x18\x01 \x01(\tH\x00\x12\x0f\n\x05\x63hunk\x18\x02 \x01(\x0cH\x00\x42\t\n\x07\x63ontent\"\'\n\x07Request\x12\x0e\n\x06header\x18\x01 \x01(\t\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t\"(\n\x08Response\x12\x0e\n\x06header\x18\x01 \x01(\t\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t2\xbf\x01\n\x0b\x46ileService\x12\x39\n\x08\x44ownload\x12\x15.FileTransfer.Request\x1a\x12.FileTransfer.File\"\x00\x30\x01\x12\x38\n\x06Upload\x12\x12.FileTransfer.File\x1a\x16.FileTransfer.Response\"\x00(\x01\x12;\n\x08NameFile\x12\x15.FileTransfer.Request\x1a\x16.FileTransfer.Response\"\x00\x42\x10Z\x0e./FileTransferb\x06proto3'
)




_FILE = _descriptor.Descriptor(
  name='File',
  full_name='FileTransfer.File',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='FileTransfer.File.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='chunk', full_name='FileTransfer.File.chunk', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
    _descriptor.OneofDescriptor(
      name='content', full_name='FileTransfer.File.content',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=36,
  serialized_end=86,
)


_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='FileTransfer.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='FileTransfer.Request.header', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='body', full_name='FileTransfer.Request.body', index=1,
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
  serialized_start=88,
  serialized_end=127,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='FileTransfer.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='FileTransfer.Response.header', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='body', full_name='FileTransfer.Response.body', index=1,
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
  serialized_start=129,
  serialized_end=169,
)

_FILE.oneofs_by_name['content'].fields.append(
  _FILE.fields_by_name['name'])
_FILE.fields_by_name['name'].containing_oneof = _FILE.oneofs_by_name['content']
_FILE.oneofs_by_name['content'].fields.append(
  _FILE.fields_by_name['chunk'])
_FILE.fields_by_name['chunk'].containing_oneof = _FILE.oneofs_by_name['content']
DESCRIPTOR.message_types_by_name['File'] = _FILE
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

File = _reflection.GeneratedProtocolMessageType('File', (_message.Message,), {
  'DESCRIPTOR' : _FILE,
  '__module__' : 'FileTransfer_pb2'
  # @@protoc_insertion_point(class_scope:FileTransfer.File)
  })
_sym_db.RegisterMessage(File)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), {
  'DESCRIPTOR' : _REQUEST,
  '__module__' : 'FileTransfer_pb2'
  # @@protoc_insertion_point(class_scope:FileTransfer.Request)
  })
_sym_db.RegisterMessage(Request)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'FileTransfer_pb2'
  # @@protoc_insertion_point(class_scope:FileTransfer.Response)
  })
_sym_db.RegisterMessage(Response)


DESCRIPTOR._options = None

_FILESERVICE = _descriptor.ServiceDescriptor(
  name='FileService',
  full_name='FileTransfer.FileService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=172,
  serialized_end=363,
  methods=[
  _descriptor.MethodDescriptor(
    name='Download',
    full_name='FileTransfer.FileService.Download',
    index=0,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_FILE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Upload',
    full_name='FileTransfer.FileService.Upload',
    index=1,
    containing_service=None,
    input_type=_FILE,
    output_type=_RESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='NameFile',
    full_name='FileTransfer.FileService.NameFile',
    index=2,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_FILESERVICE)

DESCRIPTOR.services_by_name['FileService'] = _FILESERVICE

# @@protoc_insertion_point(module_scope)
