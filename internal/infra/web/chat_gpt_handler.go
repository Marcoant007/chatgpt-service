package web

import "github.com/marcoant007/chatgptservice/internal/usecase/chatcompletionstream"

type WebChatGPTHandler struct {
	CompletionUseCase chatcompletionstream.ChatCompletionUseCase
}
