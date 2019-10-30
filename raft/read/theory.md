初始化（此时所有节点都是follower）
node a election timeout 到时间了，
node a 会变成candidate， vote 给自己， term++ （term：a1 b0 c0）
node a 发送拉票请求给其他节点
node b,c 接收到拉票请求，如果node b，c在这个term里还未vote给别的节点，就可以同意投票给node a， 同时term++(term: a1 b1 c1), node b,c 重置election timeout
node a 如果收到了大部分的投票，就从candidate升级成了leader
node a 升级之后，发送 append entries 给其followers(应该就是给它投票的节点)，发送频率依据 heartbeat timeout 而定
node b,c 收到后会回复 append entries

上述过程会一直重复，直到一个follower没有收到heartbeat，并且其将变成一个candidate



某节点断连
如果leader 也就是 node a断连了， node b,c中某个首先 heartbeat timeout，又变成了 candidate， 发起新一轮的选举
此时 node b，c的term变成了2 （term：a1 b2 c2）



日志复制（同步） log replication
所有外部client提交的变更都需要通过leader node进行转发给其他所有node，转发信息一样包含在上述 append entries中
当leader node得到大部分节点的上述消息的回复时，就可以回复client，commit成功
回复client后，leader node再通知所有follower node，上述commit已经成功了，这些follow node就应用本次commit，从而系统状态达到一致

network partitions
初始状态有 node a,b,c,d,e 5 个节点，其中 node a为leader，此时 所有节点上的term都是1
发生network partitions，导致 a,b 和 c,d,e之间断开连接
c,d,e因为无法收到lead a 发来的 append entries，3个节点内重新选举，c成为了这三个节点中的新的leader，
此时c,d,e三个节点的term变成了2， a,b的term还是1，整个系统内出现了 a  c 两个leader，只是term不同

leader a，c 同时对client提供服务
如果client的请求打到了a上，此时，因为a的append entries只能得到b的回复，无法达到半数以上，所以client的请求，a会返回失败
而如果client的请求打到了c上，此时c的append entries 能得到半数以上的回复，所有client的请求，c会返回成功

network partitions修复之后
原来的leader a收到 新的leader c的append entries，会自动降级为follower，因为c 的 term 要大于 a 的term，b 的leader也会变成 node c，同时它们的term会同步变更为 leader c的term 2
a，b更换leader 之后，会把之前的 未提交的 ，也就是之前client返回失败的 log 回滚，并将leader c上 断连这段时间的log给补上。


