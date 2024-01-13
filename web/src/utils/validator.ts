// @return true 是一个URL链接 | false 不是一个URL链接
export function isValidURL(url: string): boolean {
    try {
        if (url.length === 0) {
            return true
        }
        let u = new URL(url)
        // 判断协议
        return !(u.protocol !== "http:" && u.protocol !== "https:");
    } catch (err) {
        return false;
    }
}