const app = window.go.main.App
const rpc = { app, ...window.go.services, setPageTitle }

async function setPageTitle(title: string) {
    const prefix = await app.Title()

    window.runtime.WindowSetTitle(`${prefix} — ${title}`)
}

export default rpc
