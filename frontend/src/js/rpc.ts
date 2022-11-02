import * as app from '@wails/go/main/App'
import * as CollectionService from '@wails/go/services/CollectionService'
import * as FileSystemService from '@wails/go/services/FileSystemService'
import * as SettingsService from '@wails/go/services/SettingsService'
import { WindowSetTitle, EventsOnMultiple } from '@wails/runtime'

const rpc = { app, CollectionService, FileSystemService, SettingsService, on, setPageTitle }

function on(event: string, callback: (...data: any) => void) {
    EventsOnMultiple(event, callback, -1)
}

async function setPageTitle(title: string) {
    const prefix = await app.Title()

    WindowSetTitle(`${prefix} — ${title}`)
}

export default rpc
