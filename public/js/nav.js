function setActiveNavItem() {
    let url =  window.location.href.split('/').filter(x => x != "")
    let name = url[url.length - 1]

    document.querySelector(`#${name}`).children[0].classList.add("active")
}

setActiveNavItem()
