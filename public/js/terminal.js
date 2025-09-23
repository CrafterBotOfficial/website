messages.forEach(message => writeCommand(`echo "${message}" >> crafterbot.info`, message))

writeCommand("clear", "", clear)
writeCommand("cat ./crafterbot.info", messages.join("<br>"))

writeCommand('echo "My time is currently: $(date "+%I:%M:%S %p %Z")"', 'My time is currently: ' + serverTime)
