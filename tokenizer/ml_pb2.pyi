from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class HelloRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class Hello(_message.Message):
    __slots__ = ("res",)
    RES_FIELD_NUMBER: _ClassVar[int]
    res: str
    def __init__(self, res: _Optional[str] = ...) -> None: ...

class SymptomsRequest(_message.Message):
    __slots__ = ("text",)
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class SymptomsResponse(_message.Message):
    __slots__ = ("symptoms",)
    SYMPTOMS_FIELD_NUMBER: _ClassVar[int]
    symptoms: _containers.RepeatedCompositeFieldContainer[Symptom]
    def __init__(self, symptoms: _Optional[_Iterable[_Union[Symptom, _Mapping]]] = ...) -> None: ...

class Symptom(_message.Message):
    __slots__ = ("type", "name", "loc", "confidence")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    LOC_FIELD_NUMBER: _ClassVar[int]
    CONFIDENCE_FIELD_NUMBER: _ClassVar[int]
    type: str
    name: str
    loc: str
    confidence: float
    def __init__(self, type: _Optional[str] = ..., name: _Optional[str] = ..., loc: _Optional[str] = ..., confidence: _Optional[float] = ...) -> None: ...
