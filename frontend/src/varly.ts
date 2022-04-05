import { reactive, App, inject } from 'vue'
import { Router, useRouter } from 'vue-router'
import type { Store } from 'pinia'
import type { go } from '@/wailsjs/go'
import type { runtime as Runtime } from '@/wailsjs/runtime'
import type { Settings } from '@/wailsjs/go/models'
import { useStore, useCollectionStore } from '@/store'

const VarlySymbol = '__Varly__'

export { VarlySymbol, createVarly, useVarly }

function useVarly(): Varly {
    return inject<Varly>(VarlySymbol)!
}

function createVarly() {
    return function install(app: App) {
        const varly = reactive(new Varly(app))

        app.config.globalProperties.$varly = varly
        app.provide(VarlySymbol, varly)
    }
}

class Varly {
    vue: App
    app: go['main']['App']
    runtime: Runtime
    appStore: Store | any
    collectionStore: Store | any
    isDark: boolean

    constructor(vue: App) {
        this.vue     = vue
        this.app     = window.go.main.App
        this.runtime = window.runtime

        this.appStore = useStore()
        this.collectionStore = useCollectionStore()

        this.isDark = false

        if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
            this.setIsDark(false)
        }

        window.matchMedia('(prefers-color-scheme: dark)')
            .addEventListener('change', event => {
                if (event.matches) {
                    this.setIsDark(false)
                } else {
                    this.setIsDark(false)
                }
            })
    }

    setIsDark(value: boolean) {
        if (value) {
            this.isDark = true
            document.querySelector('html')?.classList.add('dark')
        } else {
            this.isDark = false
            document.querySelector('html')?.classList.remove('dark')
        }
    }

    getIsDark(): boolean {
        return this.isDark
    }

    async stat(file: string) {
        return await this.app.GetImageStats(file)
    }

    async getSettings(): Promise<Settings> {
        return await this.app.GetSettings()
    }

    launchTwitter(): void {
        this.runtime.BrowserOpenURL('https://twitter.com/varlyapp')
    }

    async confirmStartNewProject(title: string, message: string) {
        const response = await this.showMessageDialog({
            Title: title,
            Message: message,
            Buttons: [`OK`, `Cancel`],
            DefaultButton: `OK`
        })

        return response.toLowerCase() === 'ok'
    }

    async showMessageDialog(options: any): Promise<string> {
        try {
            return await this.app.MessageDialog(options)
        } catch (error: any) {
            this.runtime.LogError(JSON.stringify(error))
            return ''
        }
    }

    async getPreview(): Promise<string> {
        return await this.app.GetPreview(this.getConfig())
    }

    getConfig(): any {
        const layers = { ...this.collectionStore.layers }

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
            Order: [...this.collectionStore.traits].map((item: any) => item.name),
            Layers: layers,
            Width: parseInt(this.collectionStore.width.toString(), 10),
            height: parseInt(this.collectionStore.height.toString(), 10),
            Size: parseInt(this.collectionStore.size.toString(), 10)
        }

        return config
    }
    async openDirectoryDialog(): Promise<string> {
        return await this.app.OpenDirectoryDialog()
    }
}
