export default function JsonFormat(json: string): string {
    let p = [],
        push = function (m) {
            return '\\' + p.push(m) + '\\';
        },
        pop = function (m, i) {
            return p[i - 1]
        },
        tabs = function (count) {
            return new Array(count + 1).join('\t')
        }

    let out = "",
        indent = 0

    json = json
        .replace(/\\./g, push)
        .replace(/(".*?"|'.*?')/g, push)
        .replace(/\s+/, '')

    for (let i = 0; i < json.length; i++) {
        let c = json.charAt(i)

        switch (c) {
            case '{':
            case '[':
                out += c + "\r\n" + tabs(++indent)
                break
            case '}':
            case ']':
                out += "\r\n" + tabs(--indent) + c
                break
            case ',':
                out += ",\r\n" + tabs(indent)
                break
            case ':':
                out += ": "
                break
            default:
                out += c
                break
        }
    }

    out = out
        .replace(/\[[\d,\s]+?\]/g, function (m) {
            return m.replace(/\s/g, '');
        })
        .replace(/\\(\d+)\\/g, pop)
        .replace(/\\(\d+)\\/g, pop)

    return out
}