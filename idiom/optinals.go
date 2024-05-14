package main

type LockOptions struct {
	isBlock             bool
	blockWaitingSeconds int64
	expireSeconds       int64
	watchDogMode        bool
}

type LockOption func(options *LockOptions)

func WithBlock() LockOption {
	return func(options *LockOptions) {
		options.isBlock = true
	}
}

func WithBlockWaitingSeconds(waitingSeconds int64) LockOption {
	return func(options *LockOptions) {
		options.blockWaitingSeconds = waitingSeconds
	}
}

func WithExpireSeconds(expireSeconds int64) LockOption {
	return func(options *LockOptions) {
		options.expireSeconds = expireSeconds
	}
}

func NewLockOptions(options ...LockOption) *LockOptions {
	o := &LockOptions{}
	for _, option := range options {
		option(o)
	}
	return o
}
