package lib

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func InfoModal(ctx context.Context, title string, msg string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type: runtime.InfoDialog,
		Title: title,
		Message: msg,
	})
}

func ErrorModal(ctx context.Context, title string, msg string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type: runtime.ErrorDialog,
		Title: title,
		Message: msg,
	})
}
