#include <node.h>
#include "../metricsql.h"

namespace metricsql {

using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

const char* ToCString(const String::Utf8Value& value) {
  return *value ? *value : "<string conversion failed>";
}

void Method(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  v8::String::Utf8Value str(isolate, args[0]);
  const char* cstr = ToCString(str);
  //char * charstr = const_cast<char *>(cstr);
  char* result = Parse((char*)cstr);
  // args.GetReturnValue().Set(String::NewFromUtf8(isolate, result));
  args.GetReturnValue().Set(String::NewFromUtf8(isolate, result).ToLocalChecked());
  

}

void init(Local<Object> exports) {
  NODE_SET_METHOD(exports, "parse", Method);
}

NODE_MODULE(NODE_GYP_MODULE_NAME, init)

}  // namespace metricsql
