# Filebeat配置
只是单独分组展开了各个配置项，具体还是需要查看filebeat.refrence.yml文件



# 配置

## Modules Configuration
- System Module
- Apache Module
- Audit Module
- Elasticsearch Module
- HAProxy Module
- Icinga Module
- IIS Module
- Kafka Module
- Kibana Module
- Logstash Module
- Mongodbb Module
- .....


## Filebeat inputs(重要)
- log
- filestream(重要)
- stdin
- redis
- udp
- tcp
- syslog
- container


### filestream

## Filebeat autodiscover
自动发现允许您检测系统中的变化并在发生变化时产生新的模块或输入。


## Filebeat Global options
```yaml
# filebeat.registry.path: ${path.data}/registry
# filebeat.registry.file_permissions: 0600
# filebeat.registry.flush: 0s
# filebeat.registry.migrate_file: ${path.data}/registry
# filebeat.overwrite_pipelines: false
# filebeat.shutdown_timeout: 0
#filebeat.config:
  #inputs:
    #enabled: false
    #path: inputs.d/*.yml
    #reload.enabled: true
    #reload.period: 10s
  #modules:
    #enabled: true
    #path: modules.d/*.yml
    #reload.enabled: true
    #reload.period: 10s
```

## General
//TODO

## Processors

## Elastic Cloud
一般用不着，只有使用Elastic云才有用

## Outputs(重要)

### Elasticsearch Output(重要)


### Logstash Output(重要)


### Kafka Output

### Redis Output

### File Output

### Console Output

## Paths

## Keystore


## Dashboards

## Template
默认情况下ES的模版处于启用状体啊，并且会加载模版。可以调整配置，以加载自己的模版。
```yaml
#setup.template.enabled: true
#setup.template.type: index
#setup.template.name: "filebeat-%{[agent.version]}"
#setup.template.pattern: "filebeat-%{[agent.version]}-*"
#setup.template.fields: "${path.config}/fields.yml"
#setup.template.append_fields:
#- name: field_name
#  type: field_type
#setup.template.json.name: ""
#setup.template.settings:
    #index:
      #number_of_shards: 1
      #codec: best_compression
    #_source:
      #enabled: false
```


## Index Lifecycle Management (ILM)(*)

## Kibana(*)

## Logging

## X-Pack Monitoring

## HTTP Endpoint

## Process Security

## Instrumentation

## Migration
