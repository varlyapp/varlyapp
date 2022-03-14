import axios from 'axios'

const client = axios.create({
  baseURL: `/`,
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
  }
})

function postToController(action: Object, params = {}) {
  return client.post(`actions/${action}${window.location.search}`, params)
}

export { postToController }
