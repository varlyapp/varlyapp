package lib

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func ShowInfoModal(ctx context.Context, title string, msg string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type: runtime.InfoDialog,
		Title: title,
		Message: msg,
	})
}

func ShowErrorModal(ctx context.Context, title string, msg string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type: runtime.ErrorDialog,
		Title: title,
		Message: msg,
	})
}
