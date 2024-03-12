const btn = document.getElementById('btn-login');

function Login() {
    event.preventDefault();
    const data = {
        email_fisioterapeuta: document.getElementById('input-email').value,
        senha: document.getElementById('input-senha').value
    };
    fetch("http://127.0.0.1:8080/api/login", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro ao realizar login');
        }
        return response.json();
    })
    .then(data => {
        window.location = '/frontend/html/hello.html';

    })
    .catch(error => {
        const successDiv = document.getElementById('login-error');
        successDiv.textContent = 'Erro ao realizar login';
        successDiv.style.width = '200px';
        successDiv.style.padding = '20px';
        successDiv.style.marginBottom = '20px';
        successDiv.style.borderRadius = '5px';
        successDiv.style.border = '2px solid red';
    });
}

btn.addEventListener('click', Login);
