let instance

const Theme = function () {
    this.isDark = false
}

Theme.prototype.run = function () {
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

Theme.prototype.enableDarkmode = function () {
    this.isDark = true
    document.querySelector('html').classList.add('dark')
}

Theme.prototype.disableDarkmode = function () {
    this.isDark = false
    document.querySelector('html').classList.remove('dark')
}

export const useTheme = function() {
    if (!instance) {
        instance = new Theme()
    }

    return instance
}
