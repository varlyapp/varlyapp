import type { Router } from 'vue-router'
import type { Store } from 'pinia'
import type { Settings } from '../wailsjs/go/models'
import type { go } from '../wailsjs/go'
import type { runtime as Runtime } from '../wailsjs/runtime'

let intl: any
let router: Router | null | any
let appStore: Store | null | any
let collectionStore: Store | null | any

const app: go['main']['App'] = window.go.main.App
const runtime: Runtime = window.runtime

type Params = {
    intl: null | any,
    router: Router | null | any
    appStore: Store | null | any
    collectionStore: Store | null | any
}

function load(params: Params | any) {
    intl = params.intl || null
    router = params.router || router || null
    appStore = params.appStore || appStore || null
    collectionStore = params.collectionStore || collectionStore || null
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

function switchLanguage() {
    intl.locale.value = 'es'
    appStore.setLocale('es')
}

function launchTwitter(): void {
    runtime.BrowserOpenURL('https://twitter.com/varlyapp')
}

async function confirmStartNewProject(title: string, message: string) {
    const response = await showMessageDialog({
        Title: title,
        Message: message,
        Buttons: [`OK`, `Cancel`],
        DefaultButton: `OK`
    })

    if (response.toLowerCase() === 'ok') {
        collectionStore.reset()
        navigate('artwork.layers')
    }
}

async function showMessageDialog(options: any): Promise<string> {
    try {
        return await app.MessageDialog(options)
    } catch (error: any) {
        runtime.LogError(JSON.stringify(error))
        return ''
    }
}

async function getPreview(): Promise<string> {
    return await app.GetPreview(getConfig())
}

function getConfig(): any {
    const layers = { ...collectionStore.layers }

    for (const trait in Object.keys(layers)) {
        if (layers.hasOwnProperty(trait)) {
            layers[trait] = layers[trait].map((layer) => {
                return {
                    ...layer,
                    Weight: parseInt(layer.Weight)
                }
            })
        }
    }

    const config = {
        Dir: '',
        Order: [...collectionStore.traits].map((item: any) => item.name),
        Layers: layers,
        Width: parseInt(collectionStore.width.toString(), 10),
        height: parseInt(collectionStore.height.toString(), 10),
        Size: parseInt(collectionStore.size.toString(), 10)
    }

    return config
}

export { app, runtime, load, stat, navigate, getSettings, switchLanguage, launchTwitter, confirmStartNewProject, getPreview, showMessageDialog }
