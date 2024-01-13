// 转换时间
export function convertTime(a: number): string {
    const units = [
        {max: 0, unit: 'ms'},
        {max: 1000, unit: 's'},
        {max: 60000, unit: 'm'},
        {max: 3600000, unit: 'h'},
        {max: 86400000, unit: '天'},
        {max: 2592000000, unit: '月'},
    ];
    let s: string = ""
    for (let i = units.length - 1; i >= 0; i--) {
        if (a < 1000) {
            s = `${a} ${units[0].unit}`
            break
        }
        if (a >= units[i].max) {
            let value: string = (a / units[i].max).toFixed(0);
            s = `${value} ${units[i].unit}`
            break
        }
    }
    return s
}