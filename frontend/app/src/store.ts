import {Store as TauriStore} from 'tauri-plugin-store-api'

export function createStore(path: string): TauriStore | LocalStore {
  if (!import.meta.env.SSR) {
    return new TauriStore(path)
  } else {
    return new LocalStore(path)
  }
}

type LocalStorageStore<T = unknown> = {
  [key: string]: T
}

export class LocalStore {
  path: string

  constructor(path: string) {
    this.path = path
  }

  set(key: string, value: unknown): Promise<void> {
    return setFallback({path: this.path, key, value})
  }

  get<T>(key: string): Promise<T | null> {
    return getFallback({path: this.path, key})
  }

  entries<T>(): Promise<{[key: string]: T}> {
    return getStoreFallback<T>(this.path)
  }
}

function getStoreFallback<T = unknown>(path: string): Promise<LocalStorageStore<T>> {
  return new Promise((resolve, reject) => {
    try {
      let state = window.localStorage.getItem(path) || '{}'
      resolve(JSON.parse(state))
    } catch (err) {
      reject(`getStoreFallback error for ${path}: ${err}`)
    }
  })
}

type SetParams = {
  path: string
  key: string
  value: unknown
}

function setFallback({path, key, value}: SetParams): Promise<void> {
  return getStoreFallback(path)
    .then(setValueFallback)
    .catch((err) => {
      console.error(`setFallback error for ${path}: ${err}`)
    })

  function setValueFallback(store: LocalStorageStore) {
    store[key] = value
    window.localStorage.setItem(path, JSON.stringify(store))
  }
}

type GetParams = {
  path: string
  key: string
}

function getFallback<T>({path, key}: GetParams): Promise<T | null> {
  return getStoreFallback(path)
    .then(getValueFallback)
    .catch((err) => {
      console.error(`getFallback error for ${path}: ${err}`)
    })

  function getValueFallback(store: LocalStorageStore) {
    let value = (store[key] as T) || null

    return value
  }
}