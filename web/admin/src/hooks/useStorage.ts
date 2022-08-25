// useStorage
class UseStorage {
    setStorage(key: string, value: any) {
        localStorage.setItem(key, JSON.stringify(value));
    }
    getStorage(key: string) {
        const value = localStorage.getItem(key);
        return value? JSON.parse(value): null;
    }
    removeStorage(key: string) {
        localStorage.removeItem(key);
    }
    clearStorage() {
        localStorage.clear();
    }
}

export  const useStorage = new UseStorage();