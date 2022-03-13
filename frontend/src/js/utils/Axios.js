import axios from 'axios'

const language = window.language && window.language.length && window.language !== 'en' ? window.language : '';

const client = axios.create({
  baseURL: `/${language}`,
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
    'X-CSRF-TOKEN': window.csrfTokenValue,
  }
})

function postToController(action, params = {}) {
  return client.post(`actions/${action}${window.location.search}`, params)
}

export { postToController }
