# MetarProvider

[![ReleaseCard]][Release]![ReleaseDataCard]![LastCommitCard]  
![BuildStateCard]![ProjectLanguageCard]![ProjectLicense]

MetarProvider 是一个用于格式化获取 METAR 数据并提供外部查询API的 Go 语言项目

本项目提供了RESTful API接口与GRPC接口，用于获取 METAR 数据

## Feature列表

- [ ] 格式化获取METAR数据
- [ ] 解析METAR数据
- [ ] 格式化获取TAF数据
- [ ] 解析TAF数据

## 快速开始

1. 获取项目可执行文件
    - 前往[Release]页面下载最新版本
    - 前往[Action]页面下载最新开发版本
    - 手动编译本项目
2. [可选]下载项目根目录中的`config.yaml`配置文件放置于可执行文件同级目录中
3. 运行可执行文件，如果配置文件存在，则使用配置文件，否则创建默认配置文件

## 开源协议

MIT License

Copyright © 2025 Half_nothing

无附加条款。

[ReleaseCard]: https://img.shields.io/github/v/release/Flyleague-Collection/metar-provider?style=for-the-badge&logo=github

[ReleaseDataCard]: https://img.shields.io/github/release-date/Flyleague-Collection/metar-provider?display_date=published_at&style=for-the-badge&logo=github

[LastCommitCard]: https://img.shields.io/github/last-commit/Flyleague-Collection/metar-provider?display_timestamp=committer&style=for-the-badge&logo=github

[BuildStateCard]: https://img.shields.io/github/actions/workflow/status/Flyleague-Collection/metar-provider/go-build.yml?style=for-the-badge&logo=github&label=Full-Build

[ProjectLanguageCard]: https://img.shields.io/github/languages/top/Flyleague-Collection/metar-provider?style=for-the-badge&logo=github

[ProjectLicense]: https://img.shields.io/badge/License-MIT-blue?style=for-the-badge&logo=github

[Release]: https://www.github.com/Flyleague-Collection/metar-provider/releases/latest

[Action]: https://github.com/Flyleague-Collection/metar-provider/actions/workflows/go-build.yml

[Release]: https://www.github.com/Flyleague-Collection/metar-provider/releases/latest

