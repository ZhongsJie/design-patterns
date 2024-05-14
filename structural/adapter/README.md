
# 适配器模式
> 适配器模式的作用是能够实现两个不兼容或弱兼容接口之间的适配桥接作用


## 核心角色

1. 目标 target：是一类含有指定功能的接口
2. 使用方 client：需要使用 target 的用户
3. 被适配的类 adaptee：和目标类 target 功能类似，但不完全吻合
4. 适配器类：adapter：能够将 adaptee 适配转换成 target 的功能类


## Interface

1. 基于 golang 中这种隐式实现的特性，使得 interface 本身也具备了适配器的功能.
