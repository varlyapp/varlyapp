import { go } from '@/wailsjs/go/bindings'
import { runtime as Runtime } from '@/wailsjs/runtime'
import { LoggerType } from '@/types'
import { NewCollectionConfig } from '@/wailsjs/go/models'

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

async function getPreview(config: NewCollectionConfig) {
    return await api().GetPreview(config)
}

async function getFileInfo(file: string) {
    return await api().GetImageStats(file)
}

async function openFile(): Promise<string> {
    return await api().OpenFileDialog()
}

async function openFileContents(): Promise<string> {
    const path = await api().OpenFileDialog()

    if (path) {
        const contents = await api().GetFileContents(path)

        console.log(contents)

        return contents
    }
    console.log(path)

    return ''
}

export { api, runtime, log, launch, getPreview, getFileInfo, openFile, openFileContents }
