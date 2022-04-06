import { go } from '@/wailsjs/go/bindings'
// import * as models from '@/wailsjs/go/models'

import { runtime as Runtime } from '@/wailsjs/runtime'

function api(): go['main']['App'] {
    return window.go.main.App
}

function runtime(): Runtime {
    return window.runtime
}

function launch(url: string) {
    runtime().BrowserOpenURL(url)
}

export { api, runtime, launch }
