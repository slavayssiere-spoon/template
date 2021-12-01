# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: template.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from protoc_gen_openapiv2.options import annotations_pb2 as protoc__gen__openapiv2_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='template.proto',
  package='template',
  syntax='proto3',
  serialized_options=b'Z&github.com/slavayssiere-spoon/template\222A\234\004\022\214\001\n\032Spoon - Template Golang WS\"g\n\020Spoon Cloud Team\0223https://gitlab.com/SpoonQIR/Cloud/services/template\032\036sebastien.lavayssiere@spoon.ai2\0050.0.1*\002\002\0012\020application/json:\020application/jsonRP\n\003403\022I\nGReturned when the user does not have permission to access the resource.R;\n\003404\0224\n*Returned when the resource does not exist.\022\006\n\004\232\002\001\007RW\n\003418\022P\n\rI\'m a teapot.\022?\n=\032;.grpc.gateway.examples.internal.proto.examplepb.NumericEnumZ#\n!\n\nApiKeyAuth\022\023\010\002\032\rAuthorization \002b\020\n\016\n\nApiKeyAuth\022\000rD\n\rlink for docs\0223https://gitlab.com/SpoonQIR/Cloud/services/template',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0etemplate.proto\x12\x08template\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"e\n\x0cTemplateData\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12,\n\x08lastData\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0c\n\x04test\x18\x04 \x01(\x08\x32\xb8\x01\n\x08Template\x12R\n\x03Get\x12\x16.template.TemplateData\x1a\x16.template.TemplateData\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/template_golang\x12X\n\x06\x43reate\x12\x16.template.TemplateData\x1a\x16.template.TemplateData\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/template_golang:\x01*B\xc8\x04Z&github.com/slavayssiere-spoon/template\x92\x41\x9c\x04\x12\x8c\x01\n\x1aSpoon - Template Golang WS\"g\n\x10Spoon Cloud Team\x12\x33https://gitlab.com/SpoonQIR/Cloud/services/template\x1a\x1esebastien.lavayssiere@spoon.ai2\x05\x30.0.1*\x02\x02\x01\x32\x10\x61pplication/json:\x10\x61pplication/jsonRP\n\x03\x34\x30\x33\x12I\nGReturned when the user does not have permission to access the resource.R;\n\x03\x34\x30\x34\x12\x34\n*Returned when the resource does not exist.\x12\x06\n\x04\x9a\x02\x01\x07RW\n\x03\x34\x31\x38\x12P\n\rI\'m a teapot.\x12?\n=\x1a;.grpc.gateway.examples.internal.proto.examplepb.NumericEnumZ#\n!\n\nApiKeyAuth\x12\x13\x08\x02\x1a\rAuthorization \x02\x62\x10\n\x0e\n\nApiKeyAuth\x12\x00rD\n\rlink for docs\x12\x33https://gitlab.com/SpoonQIR/Cloud/services/templateb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,protoc__gen__openapiv2_dot_options_dot_annotations__pb2.DESCRIPTOR,])




_TEMPLATEDATA = _descriptor.Descriptor(
  name='TemplateData',
  full_name='template.TemplateData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='template.TemplateData.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='email', full_name='template.TemplateData.email', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lastData', full_name='template.TemplateData.lastData', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='test', full_name='template.TemplateData.test', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=139,
  serialized_end=240,
)

_TEMPLATEDATA.fields_by_name['lastData'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['TemplateData'] = _TEMPLATEDATA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TemplateData = _reflection.GeneratedProtocolMessageType('TemplateData', (_message.Message,), {
  'DESCRIPTOR' : _TEMPLATEDATA,
  '__module__' : 'template_pb2'
  # @@protoc_insertion_point(class_scope:template.TemplateData)
  })
_sym_db.RegisterMessage(TemplateData)


DESCRIPTOR._options = None

_TEMPLATE = _descriptor.ServiceDescriptor(
  name='Template',
  full_name='template.Template',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=243,
  serialized_end=427,
  methods=[
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='template.Template.Get',
    index=0,
    containing_service=None,
    input_type=_TEMPLATEDATA,
    output_type=_TEMPLATEDATA,
    serialized_options=b'\202\323\344\223\002\025\022\023/v1/template_golang',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='template.Template.Create',
    index=1,
    containing_service=None,
    input_type=_TEMPLATEDATA,
    output_type=_TEMPLATEDATA,
    serialized_options=b'\202\323\344\223\002\030\"\023/v1/template_golang:\001*',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TEMPLATE)

DESCRIPTOR.services_by_name['Template'] = _TEMPLATE

# @@protoc_insertion_point(module_scope)
