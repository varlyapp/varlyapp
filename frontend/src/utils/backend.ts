import { go } from '@/wailsjs/go/bindings'
import { runtime as Runtime } from '@/wailsjs/runtime'
import { LoggerType } from '@/types'

// import * as models from '@/wailsjs/go/models'

function api(): go['main']['App'] {
    return window.go.main.App
}

function runtime(): Runtime {
    return window.runtime
}

function log(msg: string, type?: LoggerType): void {
    if (!type) {
        runtime().LogInfo(msg)
        return
    }

    const logger = runtime()[`log${type}`]

    logger(msg)
}

function launch(url: string) {
    runtime().BrowserOpenURL(url)
}

export { api, runtime, log, launch }
