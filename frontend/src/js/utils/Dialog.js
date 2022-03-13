let instance

const Dialog = function (app) {
    this.app = app
}

Dialog.prototype.openDirectoryDialog = async function() {
    return await this.app.OpenDirectoryDialog()
}

export const useDialog = function(app) {
    if (!instance) {
        instance = new Dialog(app)
    }

    return instance
}
