# QueueService

## 一. 功能
	开服排队系统，对到达服务器的大量用户进行队列缓冲，名为QueueService，根据服务器压力情况，逐步让队列中的用户拿到登录服务器的令牌（token），代表该用户请求可以被处理了，从而缓解登录高峰，排队中用户要能够“实时”知道自己在队伍中的位置变更。
	
	已经支持功能
	1.排队的用户能实时查到自己当前所处的位置，一旦排到 异步返回登录服务器令牌（token）且马上进入游戏状态（不允许重复排队，重复排队以第一次为准）
	2.正在排队的用户退出、从缓存中踢出该用户（逻辑踢出并未真正踢出，其他用户位置不变，直到处理到已被踢出的用户消息 直接跳过，其他用户位置才真正变化）
	3.开发人员能够实时看到现在游戏中的用户数和正在排队中的用户数
	
	待开发功能
	1.正在游戏的用户退出(未支持)

## 二. 实现
![](define/cmdNo.go)

## 三. 使用

## 四. 测试

## 五. 部署

## 六. 待改进
