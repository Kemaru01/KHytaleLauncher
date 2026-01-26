package progress

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	appCtx       context.Context = nil
	progressText string          = ""
)

func InitilaizeProgess(ctx context.Context) {
	appCtx = ctx
}

func SetProgressStatus(text string, present int32) {
	progressText = text
	runtime.EventsEmit(appCtx, "progress:status", text, present)
}

func SetProgressPresent(present int32) {
	runtime.EventsEmit(appCtx, "progress:status", progressText, present)
}

func ClearProgress() {
	progressText = ""
	runtime.EventsEmit(appCtx, "progress:status", "", -1)
}
