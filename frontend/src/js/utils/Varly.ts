import type { Settings } from '@wails/go/models'
import type { runtime } from '@wails/runtime'
import type { Store } from 'pinia'
import type { Router } from 'vue-router'

const app = window.go.main.App
const runtime: runtime = window.runtime

let router: Router | null
let appStore: Store | null
let collectionStore: Store | null
let initialized: boolean

type Params = {
    router: Router | null
    appStore: Store | null
    collectionStore: Store | null
}

function init(params: Params|any) {
    router = params.router || null
    appStore = params.appStore || null
    collectionStore = params.collectionStore || null
    initialized = true
}

function panic(msg: string, data?: any) {
    if (data) {
        msg = `${msg}: ${JSON.stringify(data)}`
    }

    runtime.LogError(msg)
}

async function navigate(name: string): Promise<void> {
    try {
        await router?.push({ name })
    } catch (error) {
        console.error(`unable to navigate: ${error}`)
    }
}

async function getSettings(): Promise<Settings> {
    return await app.GetSettings()
}

function launchTwitter(): void {
    runtime.BrowserOpenURL('https://twitter.com/varlyapp')
}

export { app, runtime, init, panic, navigate, getSettings, launchTwitter }
