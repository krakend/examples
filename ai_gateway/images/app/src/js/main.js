import '../scss/styles.scss'
import Keycloak from 'keycloak-js'
import { marked } from 'marked'
import * as bootstrap from 'bootstrap'

window.addEventListener('DOMContentLoaded', async function () {
    var apiUrl = document.location.protocol + '//' + document.location.hostname + ':8080';

    const keycloak = new Keycloak({
        url: 'http://localhost:8085',
        realm: 'krakend',
        clientId: 'playground',
        flow: 'implicit'
    });

    try {
        await keycloak.init({ onLoad: 'check-sso', flow: 'implicit', checkLoginIframe: false });
    } catch (error) {
        // Skip initialization error
    }

    const notLoggedInAlert = this.document.getElementById('__alert-not-logged-in')
    const loggedInMenu = this.document.getElementById('__menu-logged-in')
    const notLoggedInMenu = this.document.getElementById('__menu-not-logged-in')
    const loginBtn = this.document.getElementById('__menu-login')
    const loginLnk = this.document.getElementById('__section-login')
    const logoutBtn = this.document.getElementById('__menu-logout')
    const homeBtn = this.document.getElementById('__menu-home')
    const profileBtn = this.document.getElementById('__menu-profile')
    const submitBtn = this.document.getElementById('__btn-submit-ai')
    const callResponseSection = this.document.getElementById('__call-response')
    const aiSection = this.document.getElementById('__section-ai')
    const profileSection = this.document.getElementById('__section-profile')
    loginBtn.addEventListener('click', async () => {
        await keycloak.login()
    })
    loginLnk.addEventListener('click', async () => {
        await keycloak.login()
    })
    logoutBtn.addEventListener('click', async () => {
        await keycloak.logout()
    })
    homeBtn.addEventListener('click', () => {
        if (keycloak.authenticated) {
            show(aiSection)
        }
        hide(profileSection)
    })
    profileBtn.addEventListener('click', () => {
        if (keycloak.authenticated) {
            show(profileSection)
        }
        hide(aiSection)
        Prism.highlightAll();
    })

    submitBtn.addEventListener('click', async () => {
        const btnInner = submitBtn.getElementsByTagName('span')[0]
        submitBtn.disabled = true
        btnInner.innerHTML = 'Loading...'
        const errorBlock = document.querySelector('#__call-response > pre')
        const errorContent = document.querySelector('#__call-response > pre > code')
        const responseBlock = document.querySelector('#__call-response > div')
        const responseContent = document.querySelector('#__call-response > div > div')
        show(callResponseSection)
        try {
            hide(errorBlock)
            hide(responseBlock)
            errorContent.innerHTML = ''
            responseContent.innerHTML = ''
            const instructions = document.getElementById('__ai-instructions').value
            const question = document.getElementById('__ai-question').value
            if (!question || question.length < 1) {
                errorContent.innerHTML = 'Please enter a valid question.'
                show(errorBlock)
                return
            }
            const req = await fetch(apiUrl + '/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${keycloak.token}`
                },
                body: JSON.stringify({ contents: question, instructions: instructions || undefined })
            })
            if (req.ok) {
                const res = await req.json()
                const output = res.output[0] || []
                const contents = output.contents.join('\n\n')
                
                errorContent.innerHTML = JSON.stringify(res, null, 4)
                responseContent.innerHTML = marked.parse(contents)
                show(errorBlock)
                show(responseBlock)
            } else {
                errorContent.innerHTML = `Error (${req.status}): ${await req.text()}`
                show(errorBlock)
            }

        } catch (err) {
            errorContent.innerHTML = 'Error: ' + err
            show(errorBlock)
        } finally {
            submitBtn.disabled = false
            btnInner.innerHTML = 'Submit'
        }

        Prism.highlightAll();
    })

    if (keycloak.authenticated) {
        show(loggedInMenu)
        hide(notLoggedInMenu)
        hide(notLoggedInAlert)

        const info = await keycloak.loadUserInfo()
        this.document.getElementById('__menu-profile-username').innerHTML = info.name
        const table = this.document.querySelector('#__section-profile > table > tbody')

        const tableMapping = { sub: "ID", email: "Email", name: "Name", preferred_username: "Username", position: "Position", locale: "Locale" }
        for (const [k, v] of Object.entries(tableMapping)) {
            const row = this.document.createElement('tr')
            const cellA = this.document.createElement('td')
            const cellB = this.document.createElement('td')
            cellA.innerHTML = v
            cellB.innerHTML = info[k]

            row.appendChild(cellA)
            row.appendChild(cellB)
            table.appendChild(row)
        }
        this.document.querySelector('#__section-profile > pre > code').innerHTML = JSON.stringify(info, null, 4)
    } else {
        show(notLoggedInMenu)
        show(notLoggedInAlert)
        hide(loggedInMenu)
        hide(aiSection)
    }
})

function hide(el) {
    if (!el.classList.contains('hidden')) {
        el.classList.add('hidden')
    }
}

function show(el) {
    if (el.classList.contains('hidden')) {
        el.classList.remove('hidden')
    }
}
