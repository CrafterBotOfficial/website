var container = document.getElementById('container')

const header = `
Welcome to Ubuntu 24.9 (GNU/Linux 6.10.0-1007-raspi aarch64)<br>
    <br>
    * Documentation:  https://help.ubuntu.com<br>
    * Management:     https://landscape.canonical.com<br>
    * Support:        https://ubuntu.com/pro<br>
    <br>
    Last login: ${getLastLoginTime()}<br>
    `

const prefix = '<span id="context">crafterbot@ubuntu</span>:<span id="context-path">~</span>$ ' //"crafterbot@ubuntu:~$ " //
const writerTempo = 10

var typewriter = new Typewriter(container, {
    delay: writerTempo,
})
var wrapper = document.querySelector(".Typewriter__wrapper")

wrapper.innerHTML = header + prefix

function clear() {
    wrapper.innerHTML = prefix
}

function writeCommand(command, output, action) {
    const hasAction = action != undefined 
    typewriter
        .pauseFor(!hasAction ? 250 : 0)
        .typeString(command)
        .pauseFor(250)
        .pasteString("<br>" + output)
        .pasteString("<br>" + prefix)
        .pauseFor(!hasAction ? 1000 : 0)
        .callFunction(() => {
            if (action != undefined) {
                action()
            }
        })
        .start()
}
