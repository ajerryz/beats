# filebeat.yml

# setup
核心结论：Filebeat 的 setup 相关配置是初始化阶段的核心开关，用于在 Filebeat 启动采集数据前，自动创建 / 配置 ES 索引模板、Kibana 可视化资源、ILM 策略等基础依赖，避免手动操作，核心作用是 “一键搭建数据链路所需的底层资源”，按功能可分为 “模板配置、Kibana 配置、ILM 配置、其他辅助配置” 四大类。

## 一、核心前提：setup 阶段的触发时机
- 执行 filebeat setup 命令时，会强制执行所有 setup 配置对应的初始化操作。
- 启动 Filebeat 时（filebeat run），若 setup.* 配置为 true（默认部分开启），会自动执行初始化（仅首次执行或资源不存在时生效）。
- 生产环境建议：先执行 filebeat setup 完成初始化，再启动 Filebeat 采集数据，避免启动时因初始化耗时影响采集。


## 二、setup 核心配置分类详解

- setup.template.enabled: true或false,是否在setup阶段执行template相关操作
- setup.template.name: 自定义 Filebeat 自动创建的模板名称。默认为 `filebeat-%{[agent.version]}`
- setup.template.pattern: 自定义模板匹配的索引模式。默认为 `filebeat-*`
- setup.template.overwrite: 当模板已存在时，是否强制覆盖。默认为 false
- setup.template.json.enabled: 是否从一个 JSON 文件加载模板，而不是使用内置的模板。
- setup.template.json.path:  指向自定义模板 JSON 文件的路径
- setup.template.settings.* : 为自动创建的模板添加 ES settings


- setup.kibana.enabled: 开关：是否连接 Kibana 并初始化资源
- setup.kibana.hosts: Kibana 服务地址
- setup.kibana.username/password: Kibana 认证账号密码
- setup.kibana.ssl.* : Kibana 启用 HTTPS 时的 SSL 配置
- setup.dashboards.enabled: 开关：是否导入 Filebeat 内置仪表盘
- setup.dashboards.index: 仪表盘关联的 ES 索引模式
- setup.dashboards.id: 指定导入的仪表盘 ID（按需筛选）
- setup.index_pattern.enabled: 开关：是否自动创建 Kibana Index Pattern
- setup.index_pattern.name: 自动创建的 Kibana Pattern 名称


- setup.ilm.enabled: 开关：是否自动创建 / 关联 ILM 策略
- setup.ilm.policy_name: 自定义 ILM 策略名称
- setup.ilm.policy_file: 本地 ILM 策略 JSON 文件路径
- setup.ilm.rollover_alias: ILM 滚动别名
- setup.ilm.pattern: ILM 策略匹配的索引模式
