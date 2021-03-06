## kubegrid roadmap 

### alpha（2022-06 ~ 2022-09）

#### 标准:

* kubespray 功能基本就绪，基线完成 70%
* 拥有充分的 e2e 测试用例
* Unit-Test 覆盖率 50%
* ClusterAPI 调研

#### 关键 milestone

- [ ]  最小功能 mvp（快速验证，通过 kubespray operator 创建集群）
- [ ]  支持主流操作系统(CentOS、RedHat)
- [ ]  集群 LCM 管理(升级业务需要连续不中断)
    - [ ]  集群一键创建(支持私有云虚拟化场景、物理机 给到 IP 和登录凭证)
    - [ ]  集群升级
    - [ ]  集群删除
    - [ ]  kubernetes 连续大版本升级,应用向下兼容
    - [ ]  一键轮滚证书
    - [ ]  详尽的安装/运维日志记录
- [ ]  节点 LCM 管理
    - [ ]  新增节点(在线+离线)
    - [ ]  删除节点
    - [ ]  节点污点管理
    - [ ]  主机的标签管理
    - [ ]  详尽的安装/运维日志记录
    - [ ]  CRI 支持
        - [ ]  containerd
        - [ ]  docker
        - [ ]  runc
        - [ ]  docker
        - [ ]  kata
        - [ ]  gvisor
- [ ]  集群检查
    - [ ]  安装前检查
    - [ ]  安装后检查
- [ ]  集群组件高可用
    - [ ]  单台 master 节点宕机不影响集群正常使用
    - [ ]  管理节点故障后，集群运行状态不受影响，业务正常使用
    - [ ]  容器管理平台节点故障，可自动弹探测、切换，不影响平台自身的对外提供服务
- [ ]  异构架构环境支持(海光、鲲鹏、麒麟操作系统 arm64 && amd64) (2022-07)
- [ ]  性能
    - [ ]  集群安装、运维效率(先确定量化方法论，再得出测试结论)
- [ ]  审计日志开关功能
- [ ]  kubean 离线资源的维护(通过 CI 自动化，拒绝人肉)
- [ ]  对于基线中涉及到 ClusterAPI 相关的内容做可行性探索(2022-09)
    - [ ]  节点一键扩缩容(支持私有云+公有云)。
    - [ ]  支持集群的主机/虚拟机根据规则自动化 扩容和缩容
    - [ ]  节点池管理
    - [ ]  物理节点的平滑扩容
- [ ]  整理十条最常见的问题排查 FAQ


### beta（2022 Q4）

#### 标准:

* kubean 功能开发完成基线的 90%
* 异构架构支持
* 基本功能全部就绪
* UT 单元测试达到 70% 
* 完善常见的 FAQ
* clusterAPI 接入、就绪

#### 关键 milestone

- [ ]  支持主流操作系统(Suse、Ubuntu)
- [ ]  支持 GPU
- [ ]  集群 LCM 管理(升级业务需要连续不中断)
    - [ ]  kubernetes 版本降级
    - [ ]  UI 测升级网络插件
- [ ]  节点 LCM 管理
    - [ ]  支持集群的主机/虚拟机根据规则自动化 扩容和缩容
    - [ ]  物理节点的平滑扩容
    - [ ]  节点一键扩缩容(支持私有云+公有云)。
    - [ ]  节点池管理
- [ ]  集群检查
    - [ ]  支持集群层面、主机、服务巡检，出报告
- [ ]  大规模
    - [ ]  支持 200+ 节点
    - [ ]  支持 2000+ pod
    - [ ]  支持 1000+ namespaces
- [ ] SR-IOV 硬件加速 

### GA

#### 标准:

* 完成基线划定的内容
* 拥有充分的 e2e 测试用例
* UT 单元测试达到 80% 
* 实现部分魅力型需求

#### 关键 milestone
- [ ]  支持 linux + windows 异构环境
- [ ]  节点 LCM 管理
    - [ ]  CRI 支持
        - [ ]  beyondVM
        - [ ]  pouch
        - [ ]  kubevirt
- [ ]  容灾备份
    - [ ]  平台组件备份还原
    - [ ]  应用备份还原
- [ ]  安全
    - [ ]  cis benchmark


## 文档参考

- [第五代产品的详细关键里程碑](https://dwiki.daocloud.io/pages/viewpage.action?pageId=103948195)

- [kubean 基线梳理](https://dwiki.daocloud.io/pages/viewpage.action?pageId=121174913)