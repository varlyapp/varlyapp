import type { go } from '@wails/go/bindings'
import type { runtime } from '@wails/runtime'
import type { App } from 'vue'
import type { Router } from 'vue-router'

export interface Context {
    router: Router
}

export class Varly {
    app: go["main"]["App"]
    runtime: runtime
    router: Context['router']

    constructor(context: Context) {
        this.app = window.go.main.App
        this.runtime = window.runtime
        this.router = context.router
    }
}

export default {
    install: (app: App, options: Context) => {
        app.provide('varly', new Varly(options))
    }
}
