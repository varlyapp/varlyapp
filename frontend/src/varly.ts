import { App, inject, reactive } from 'vue'

const VarlySymbol = '__VARLY__'

export { VarlySymbol, createVarly, useVarly }

function useVarly(): Varly {
    return inject<Varly>(VarlySymbol)!
}

function createVarly() {
    return function install(app: App) {
        const varly = reactive(new Varly())

        app.config.globalProperties.$varly = varly
        app.provide(VarlySymbol, varly)
    }
}

class Varly {
    isDark: boolean

    constructor() {
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
}
