play_canvas:
	go run *.go -debug

train_ai: ./ai/train.go
	bash ./ai/train.go

.PHONY: train_ai