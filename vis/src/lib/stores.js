import { browser } from '$app/env';
import { writable } from 'svelte/store';


const storage = (key, initValue) => {
  const store = writable(initValue);
  if (!browser) {
    return store
  };

  const storedValueStr = localStorage.getItem(key);
  if (!!storedValueStr) {
    store.set(JSON.parse(storedValueStr))
  };

  store.subscribe((val) => {
    if ([null, undefined].includes(val)) {
      localStorage.removeItem(key)
    } else {
      localStorage.setItem(key, JSON.stringify(val))
    }
  })

  window.addEventListener('storage', () => {
    const storedValueStr = localStorage.getItem(key);
    if (storedValueStr == null) {
      return;
    }

    const localValue = JSON.parse(storedValueStr)
    if (localValue !== get(store)) {
      store.set(localValue);
    }
  });

  return store;
}

export const topics = storage("topics", [])
export const processed = storage("processed", [])
