function setActiveNavItem() {
    let url =  window.location.href.split('/').filter(x => x != "")
    let name = url[url.length - 1]

    let navObject = document.querySelector(`#${name}`)
    if (navObject != null) {
        navObject.children[0].classList.add("active")
        return
    }
    document.querySelector('#articles').children[0].classList.add("active")
}

setActiveNavItem()
