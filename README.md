# Kikitoru

一个同人音声专用的音乐流媒体服务器

**！本项目仍在早期开发阶段，可能会出现预期之外的问题或 API 变更！**

### 特点
- 重写的 [kikoeru](https://github.com/kikoeru-project/kikoeru-express)，尽量保持 API 不变，可以兼容原版前端及第三方 APP 的大部分功能
- 使用 Golang 编写，性能和稳定性更高
- 数据库使用 Postgresql
- 扫描速度快，900 部作品建立数据库并下载封面仅需 1min15s 
- 内存占用低 ~22M

### 功能介绍
- 从 DLSite 爬取音声元数据
- 对音声标记进度、打星、写评语
- 通过标签或关键字快速检索想要找到的音声
- 根据音声元数据对检索结果进行排序
- 支持在 Web 端修改配置文件和扫描音声库
- 支持为音声库添加多个根文件夹
