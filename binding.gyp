{
  "targets": [
    {
      "target_name": "addon",
      "sources": [ "src/metricsql.cc" ],
      "libraries": [ "<!(pwd)/metricsql.so" ],
      "ldflags" : [ "-Wl,-s" ]
    }
  ]
}
