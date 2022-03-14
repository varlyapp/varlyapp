let instance: Theme

class Theme {
    isDark: boolean = false

    run() {
        if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
            this.enableDarkmode()
        }

        window.matchMedia('(prefers-color-scheme: dark)')
            .addEventListener('change', event => {
                if (event.matches) {
                    this.enableDarkmode()
                } else {
                    this.disableDarkmode()
                }
            })
    }

    enableDarkmode() {
        this.isDark = true
        document.querySelector('html')?.classList.add('dark')
    }

    disableDarkmode() {
        this.isDark = false
        document.querySelector('html')?.classList.remove('dark')
    }
}

function useTheme(): Theme {
    if (!instance) {
        instance = new Theme()
    }

    return instance
}

export { useTheme }
