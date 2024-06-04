## 架构

![img](https://cdn.nlark.com/yuque/0/2024/png/366030/1716364494355-f3d3e104-63aa-412a-a26a-648fbc5eeafb.png)

## 流程

1. 基于初始化高度同步 **DA node** 的区块数据
2. 过滤 **DA** 区块中的**VerifyBatchesTrustedAggregator**事件，并保存到数据库 
3. 根据**VerifyBatchesTrustedAggregator** 事件，以**proposal**的形式提交**proofRootHash**和**stateRootHash**到DA 节点,此时**proposal**等待投票状态
4. 循环校验上述**proposal**的投票状态,只要不是投票状态就更新**DA**里**proposal**的状态到本地数据库
5. 根据**proposal**状态进行**incribe**处理

1. 1. 如果已经是成功状态，意味着得到了**abeTXhash** ，直接更新数据库
   2. 如果是**pending**状态，并且此轮**proposal**提交的胜者是自己，则构造一笔TXMemo交易提交**proposal**的**stateroot**哈希和**proofroot**哈希到**abe**链，拿到**txHash**,之后将提交**BitcoinTx**，将**proposalID**和 **tx**哈希进行关联，只有超过半数的人进行了投票才会把**proposal**的状态设置为**successedStatus**，同时保存到数据库，注意这里数据库的的**proposal**状态还是**pending**状态
   3. 当其他**committer** 确认已经有了在**abe**链上的**txhash**，并且proposa状态为**pendingstatus**时候，其余人也会提交**abeTx**，实际还是处理上面的交易，只是算是一个投票，当大于半数则**DA**节点上的**proposal**状态会为**succedStatus**

## 数据库设计

### sync_blocks

| 列名         | 数据类型    | 描述                                       | 索引情况                  |
| ------------ | ----------- | ------------------------------------------ | ------------------------- |
| id           | bigint      | 自增ID                                     | 主键                      |
| created_at   | datetime    | 创建时间，默认为当前时间戳                 | 无                        |
| updated_at   | datetime    | 更新时间，默认为每次更新时的当前时间戳     | 无                        |
| blockchain   | varchar(32) | DA链名                                     | 无                        |
| miner        | varchar(42) | 出块人的地址，以太坊地址一般为42字符       | 无                        |
| block_time   | bigint      | 出块时间，一般为Unix时间戳格式             | 无                        |
| block_number | bigint      | 区块高度，链上的连续编号                   | 无                        |
| block_hash   | varchar(66) | 区块的哈希值，长度通常为66字符（含前导0x） | 无                        |
| tx_count     | bigint      | 该区块包含的交易数量                       | 索引（tx_count_index）    |
| event_count  | bigint      | 区块内verifyTrustedAggegator事件的总数     | 无                        |
| parent_hash  | varchar(66) | 父区块的哈希值                             | 无                        |
| status       | varchar(32) | 状态                                       | 索引（status_index）      |
| check_count  | bigint      | 代码未涉及                                 | 索引（check_count_index） |



### sync_events

| 列名              | 数据类型    | 描述                                      | 索引情况             |
| ----------------- | ----------- | ----------------------------------------- | -------------------- |
| id                | bigint      | 自增ID                                    | 主键                 |
| created_at        | datetime    | 创建时间，默认为当前时间戳                | 无                   |
| updated_at        | datetime    | 更新时间，默认为每次更新时的当前时间戳    | 无                   |
| sync_block_id     | bigint      | 关联的同步区块ID，对应`sync_blocks`表的ID | 无                   |
| blockchain        | varchar(32) | DA 链                                     | 无                   |
| block_time        | bigint      | 区块时间，一般为Unix时间戳格式            | 无                   |
| block_number      | bigint      | 事件所在区块高度                          | 无                   |
| block_hash        | varchar(66) | 事件所在区块哈希                          | 无                   |
| block_log_indexed | bigint      | 日志在区块中的索引位置                    | 无                   |
| tx_index          | bigint      | 交易在区块中的索引位置                    | 无                   |
| tx_hash           | varchar(66) | 交易的哈希值                              | 无                   |
| event_name        | varchar(32) | 事件名称                                  | 无                   |
| event_hash        | varchar(66) | 事件的哈希值                              | 无                   |
| contract_address  | varchar(42) | 触发事件的合约地址                        | 无                   |
| data              | json        | 事件携带的具体数据内容                    | 无                   |
| status            | varchar(32) | 状态                                      | 索引（status_index） |
| retry_count       | bigint      | 处理失败时的重试次数，默认为0             | 无                   |



### proposal

| 列名               | 数据类型     | 描述                                               | 索引情况                  |
| ------------------ | ------------ | -------------------------------------------------- | ------------------------- |
| id                 | bigint       | 自增ID                                             | 主键                      |
| created_at         | datetime     | 创建时间，默认为当前时间戳                         | 无                        |
| updated_at         | datetime     | 更新时间，默认为每次更新时的当前时间戳             | 无                        |
| proposal_id        | bigint       | 提案的唯一标识符                                   | 索引（proposal_id_index） |
| proposer           | varchar(128) | 提议者的地址                                       | 无                        |
| state_root_hash    | varchar(128) | 状态根哈希                                         | 无                        |
| proof_root_hash    | varchar(128) | 证明根哈希                                         | 无                        |
| start_batch_num    | bigint       | 批次开始编号                                       | 无                        |
| end_batch_num      | bigint       | 批次结束编号                                       | 无                        |
| btc_commit_tx_hash | varchar(128) | 代码未涉及                                         | 无                        |
| btc_reveal_tx_hash | varchar(128) | 发送的memo交易，存储上面的staterootHash和proof哈希 | 无                        |
| block_height       | bigint       | proposal创建高度                                   | 无                        |
| winner             | varchar(128) | 提案的获胜者                                       | 无                        |
| status             | bigint       | 提案的状态编码，用于区分提案的不同阶段             | 无                        |