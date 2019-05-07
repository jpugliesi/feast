# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: feast/types/FeatureRow.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from feast.types import Feature_pb2 as feast_dot_types_dot_Feature__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='feast/types/FeatureRow.proto',
  package='feast.types',
  syntax='proto3',
  serialized_options=_b('\n\013feast.typesB\017FeatureRowProtoZ6github.com/gojek/feast/protos/generated/go/feast/types'),
  serialized_pb=_b('\n\x1c\x66\x65\x61st/types/FeatureRow.proto\x12\x0b\x66\x65\x61st.types\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x19\x66\x65\x61st/types/Feature.proto\"j\n\rFeatureRowKey\x12\x11\n\tentityKey\x18\x01 \x01(\t\x12\x32\n\x0e\x65ventTimestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x12\n\nentityName\x18\x04 \x01(\t\"\x8f\x01\n\nFeatureRow\x12\x11\n\tentityKey\x18\x01 \x01(\t\x12&\n\x08\x66\x65\x61tures\x18\x02 \x03(\x0b\x32\x14.feast.types.Feature\x12\x32\n\x0e\x65ventTimestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x12\n\nentityName\x18\x04 \x01(\tBV\n\x0b\x66\x65\x61st.typesB\x0f\x46\x65\x61tureRowProtoZ6github.com/gojek/feast/protos/generated/go/feast/typesb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,feast_dot_types_dot_Feature__pb2.DESCRIPTOR,])




_FEATUREROWKEY = _descriptor.Descriptor(
  name='FeatureRowKey',
  full_name='feast.types.FeatureRowKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entityKey', full_name='feast.types.FeatureRowKey.entityKey', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='eventTimestamp', full_name='feast.types.FeatureRowKey.eventTimestamp', index=1,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entityName', full_name='feast.types.FeatureRowKey.entityName', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=105,
  serialized_end=211,
)


_FEATUREROW = _descriptor.Descriptor(
  name='FeatureRow',
  full_name='feast.types.FeatureRow',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entityKey', full_name='feast.types.FeatureRow.entityKey', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='features', full_name='feast.types.FeatureRow.features', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='eventTimestamp', full_name='feast.types.FeatureRow.eventTimestamp', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entityName', full_name='feast.types.FeatureRow.entityName', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=214,
  serialized_end=357,
)

_FEATUREROWKEY.fields_by_name['eventTimestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_FEATUREROW.fields_by_name['features'].message_type = feast_dot_types_dot_Feature__pb2._FEATURE
_FEATUREROW.fields_by_name['eventTimestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['FeatureRowKey'] = _FEATUREROWKEY
DESCRIPTOR.message_types_by_name['FeatureRow'] = _FEATUREROW
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FeatureRowKey = _reflection.GeneratedProtocolMessageType('FeatureRowKey', (_message.Message,), dict(
  DESCRIPTOR = _FEATUREROWKEY,
  __module__ = 'feast.types.FeatureRow_pb2'
  # @@protoc_insertion_point(class_scope:feast.types.FeatureRowKey)
  ))
_sym_db.RegisterMessage(FeatureRowKey)

FeatureRow = _reflection.GeneratedProtocolMessageType('FeatureRow', (_message.Message,), dict(
  DESCRIPTOR = _FEATUREROW,
  __module__ = 'feast.types.FeatureRow_pb2'
  # @@protoc_insertion_point(class_scope:feast.types.FeatureRow)
  ))
_sym_db.RegisterMessage(FeatureRow)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)