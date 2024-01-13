import {v4 as uuid4} from 'uuid';

export default function Uuid() {
    let u = uuid4();
    u = u.replace(/-/g, "").toUpperCase()
    return u;
}