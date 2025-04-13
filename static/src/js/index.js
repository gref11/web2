document.querySelector("#logo").addEventListener("click", clickLogo);
const menu = document.querySelector("#menu");
const menuList = document.querySelector("#menu-list");
const contactForm = document.getElementById('regForm');
const responseMessage = document.getElementById('responseMessage');

let media = window.matchMedia("(max-width: 800px)");

function getRandomInt(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min) + min); 
    }

function clickLogo(e) {
    e.target.src=`/static/src/img/logo${getRandomInt(2, 4)}.png`;
    e.target.removeEventListener("click", clickLogo);
}

function clickMenu() {
    menuList.classList.toggle("hidden");
    if (menuList.classList.contains("hidden"))
        menu.querySelector("span").innerHTML = "Меню ∨";
    else
        menu.querySelector("span").innerHTML = "Меню ∧";
}

function hideMenu(e) {
    if (e.target.id == "menu" || e.target.parentElement && e.target.parentElement.id == "menu" || menuList.classList.contains("hidden"))
        return;
    clickMenu();
}

document.addEventListener("click", hideMenu);
menu.addEventListener("click", clickMenu);

loadFormData();

regForm.onsubmit = async function(e) {
    e.preventDefault();

    const form = e.target;
    const data = {
        fullName: form.fullName.value,
        email: form.email.value,
        phone: form.phone.value,
        date: form.date.value,
        sex: form.sex.value,
        lang: Array.from(form.lang.selectedOptions).map(o => o.value),
        bio: form.bio.value,
        agreement: form.agreement.checked
    };

    try {
        const response = await fetch('/user/add/', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: { 
                'Content-Type': 'application/json',
                'Accept': 'text/plain'
            }
        });

        const responseText = await response.text();

        if (!response.ok) throw new Error(responseText);

        // const result = await response.json();
        console.log(response.body)
        responseMessage.textContent = responseText;
        responseMessage.style.color = 'green';

        saveFormData();
        // contactForm.reset();

    } catch (error){ 
        responseMessage.textContent = 'Ошибка отправки: ' + error.message;
        responseMessage.style.color = 'red';
        console.log('Response status:', response.status);
        console.log('Response body:', error.message);
    }
}


function saveFormData() {
    const formData = {
        fullName: document.getElementById('fullName').value,
        email: document.getElementById('email').value,
        phone: document.getElementById('phone').value,
        date: document.getElementById('date').value,
        sex: document.getElementById('sex').value,
        bio: document.getElementById('bio').value,
    };
    localStorage.setItem('regData', JSON.stringify(formData));
}

function loadFormData() {
    const savedData = localStorage.getItem('regData');
    if (savedData) {
        const formData = JSON.parse(savedData);
        document.getElementById('fullName').value = formData.fullName || '';
        document.getElementById('email').value = formData.email || '';
        document.getElementById('phone').value = formData.phone || '';
        document.getElementById('date').value = formData.date || '';
        document.getElementById('sex').value = formData.sex || '';
        document.getElementById('bio').value = formData.bio || '';
        
    }
}

// function sendForm(e) {
//     e.preventDefault();
//     saveFormData()
//     const form = e.target;
//     const data = {
//         fullName: form.fullName.value,
//         email: form.email.value,
//         phone: form.phone.value,
//         date: form.date.value,
//         sex: form.sex.value,
//         lang: Array.from(form.lang.selectedOptions).map(o => o.value),
//         bio: form.bio.value,
//         agreement: form.agreement.checked
//     };
//     console.log(data);

//     fetch('http://localhost:8080/user/add/', {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify(data)
//     })
//     .then(response => response.text())
//     .then(data => console.log(data))
//     .catch(error => console.error(error));
// }