import * as app from '@/wailsjs/go/main/App'
import * as CollectionService from '@/wailsjs/go/services/CollectionService'
import * as FileSystemService from '@/wailsjs/go/services/FileSystemService'
import * as SettingsService from '@/wailsjs/go/services/SettingsService'
import { WindowSetTitle, EventsOnMultiple } from '@/wailsjs/runtime'

const rpc = { app, CollectionService, FileSystemService, SettingsService, on, setPageTitle }

function on(event: string, callback: (...data: any) => void) {
    EventsOnMultiple(event, callback, -1)
}

async function setPageTitle(title: string) {
    const prefix = await app.Title()

    WindowSetTitle(`${prefix} — ${title}`)
}

export default rpc
