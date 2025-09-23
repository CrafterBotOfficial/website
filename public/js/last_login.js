function getLastLoginTime() {
    const KEY = "last-login"
    let Now = new Date().getTime()

    var lastLogin = localStorage.getItem(KEY)
    localStorage.setItem(KEY, getTimeFormat(Now))

    if (lastLogin == null) return "n/a"
    return lastLogin + " from 127.0.0.1"
}

function getTimeFormat(time) {
    const date = new Date(+time)
    const options = { weekday: 'short', month: 'short', day: '2-digit', year: 'numeric', hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false }
    return new Intl.DateTimeFormat('en-US', options).format(date).replaceAll(',', '')
}
