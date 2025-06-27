# 环境变量字段说明

## 区块链与节点相关

- **RPC_URL**：以太坊/Polygon 节点的 RPC 接口地址，用于链上数据同步和交互。
- **BLOCKCHAIN**：指定使用的区块链类型（如 polygon、ethereum 等），影响链上交互逻辑。
- **INIT_BLOCK_NUMBER**：服务启动时同步的起始区块号。
- **INIT_BLOCK_HASH**：服务启动时同步的起始区块哈希。
- **POLYGON_ZKEVM_ADDRESS**：Polygon zkEVM 合约地址，用于监听和交互。
- **PROPOSAL_BATCHES_LIMITNUM**：每次批量处理的最大 proposal 数量。
- **INIT_PROPOSAL_ID**：服务启动时同步的起始 proposal ID。

## DANode 相关
- **B2NODE_PRIVATE_KEY**：DANode 节点的私钥，用于签名链上交易。
- **B2NODE_ADDRESS**：DANode 节点的链上地址。
- **B2NODE_CHAIN_ID**：DANode 所在链的 Chain ID。
- **B2NODE_GRPC_HOST**：DANode gRPC 服务主机名。
- **B2NODE_GRPC_PORT**：DANode gRPC 服务端口。
- **B2NODE_RPC_URL**：DANode 的 RPC 接口地址。
- **B2NODE_COIN_DENOM**：DANode 链上代币单位（如 abecoin）。

## 业务 API 相关
- **ENDPOINT**：Abelian 业务 API 的主接口地址。
- **RPCENDPOINT**：Abelian 业务 API 的 RPC 接口地址。
- **USERNAME**：Abelian 业务 API 认证用户名。
- **PASSWORD**：Abelian 业务 API 认证密码。
- **APPID**：Abelian 业务 API 分配的应用 ID。
- **REQUEST_SIGNATURE**：Abelian 业务 API 请求签名，用于接口安全校验。
- **USERID**：Abelian 业务 API 用户 ID。
- **FROM**：发起转账的 Abelian 用户地址。
- **RECIPIENT**：收款方 Abelian 用户地址。
- **PRIVATE_KEY**：Abelian 业务 API 用户私钥。
- **AUTHTOKEN**：Abelian 业务 API 认证 Token。

## 镜像相关
- **ETHERMINT_IMAGE**：Ethermint 节点镜像名及标签，用于 docker-compose 启动 ethermint 服务。

---

如需进一步了解每个字段的具体用法，请参考项目源码或联系开发者。 