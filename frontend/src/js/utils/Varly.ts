import type { Settings } from '@wails/go/models'
import type { runtime } from '@wails/runtime'
import type { Store } from 'pinia'
import type { Router } from 'vue-router'

const app = window.go.main.App
const runtime: runtime = window.runtime

let router: Router | null | any
let appStore: Store | null | any
let collectionStore: Store | null | any
let initialized: boolean

type Params = {
    router: Router | null | any
    appStore: Store | null | any
    collectionStore: Store | null | any
}

function init(params: Params|any) {
    if (!initialized) {
        router = params.router || null
        appStore = params.appStore || null
        collectionStore = params.collectionStore || null
        initialized = true
    }
}

function panic(msg: string, data?: any) {
    if (data) {
        msg = `${msg}: ${JSON.stringify(data)}`
    }

    runtime.LogError(msg)
}

async function stat(file: string) {
    return await app.GetImageStats(file)
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

async function confirmStartNewProject() {
    const response = await showMessageDialog({
        Title: "Are you sure you want to start a new project?",
        Message: "Starting a new project without saving your current one, will tell Varly to discard your progress",
        Buttons: ["Ok", "Cancel"],
        DefaultButton: "Ok"
    })

    if (response.toLowerCase() === 'ok') {
        collectionStore ? collectionStore.reset() : ''
        navigate('artwork.layers')
    }
}

async function showMessageDialog(options: any): Promise<string> {
    try {
        return await app.MessageDialog(options)
    } catch (error: any) {
        panic(error)
        return ''
    }
}

export { app, runtime, init, panic, stat, navigate, getSettings, launchTwitter, confirmStartNewProject, showMessageDialog }
