messages.forEach(message => writeCommand(`echo "${message}" >> crafterbot.info`, message))

writeCommand("clear", "", clear)
writeCommand("cat ./crafterbot.info", messages.join("<br>"))
