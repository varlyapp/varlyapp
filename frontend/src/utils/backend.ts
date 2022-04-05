import { go } from '@/wailsjs/go/bindings'
import { runtime as Runtime } from '@/wailsjs/runtime'

function api(): go['main']['App'] {
    return window.go.main.App
}

function runtime(): Runtime {
    return window.runtime
}

export { api, runtime }
