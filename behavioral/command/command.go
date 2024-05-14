package main

import "fmt"

type Command interface {
	Call()
}

type Receiver interface {
	Execute()
}

type Invoker struct {
	cmds []Command
}

func (in *Invoker) AddCommand(c Command) {
	if in == nil {
		return
	}
	in.cmds = append(in.cmds, c)
}

func (in *Invoker) ExecuteCommand() {
	if in == nil || len(in.cmds) == 0 {
		return
	}
	for _, cmd := range in.cmds {
		cmd.Call()
	}
}

type ReceiverA struct {
}

func (a *ReceiverA) Execute() {
	fmt.Println("接收者A处理请求")
}

// OneCommand 创建具体command, 指定接收者
type OneCommand struct {
	Receiver *ReceiverA
}

func (o OneCommand) Call() {
	o.Receiver.Execute()
}

func main() {
	// 接收者
	receiver := &ReceiverA{}
	// command
	command := OneCommand{Receiver: receiver}
	// 调用者
	invoker := new(Invoker)
	invoker.AddCommand(command)
	// 执行命令
	invoker.ExecuteCommand()
}
