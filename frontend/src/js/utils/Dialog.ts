import type { go } from '@wails/go/bindings'

let instance: Dialog

class Dialog {
    app: go["main"]["App"]
    constructor(app: go["main"]["App"]) {
        this.app = app
    }

    async openDirectoryDialog(): Promise<string> {
        return await this.app.OpenDirectoryDialog()
    }
}

function useDialog(app: go["main"]["App"]): Dialog {
    if (!instance) {
        instance = new Dialog(app)
    }

    return instance
}

export { useDialog }
