const btn = document.getElementById('btn-cadastro');

function CadastroUsuario() {
    event.preventDefault();
    const data = {
        nome_fisioterapeuta: document.getElementById('input-nome').value,
        email_fisioterapeuta: document.getElementById('input-email').value,
        senha: document.getElementById('input-senha').value
    };
    fetch("http://127.0.0.1:8080/api/cadastro", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Erro ao cadastrar usuário');
        }
        return response.json();
    })
    .then(data => {
        const successDiv = document.getElementById('cadastro-success');
        successDiv.textContent = 'Cadastro realizado com sucesso';
        successDiv.style.width = '200px';
        successDiv.style.padding = '20px';
        successDiv.style.marginBottom = '20px';
        successDiv.style.borderRadius = '5px';
        successDiv.style.border = '2px solid black';

        setTimeout(() => {
            window.location = '/frontend/html/login.html';
        }, 3000); 
    })
    .catch(error => {
        const successDiv = document.getElementById('cadastro-error');
        successDiv.textContent = 'Erro ao realizar cadastro';
        successDiv.style.width = '200px';
        successDiv.style.padding = '20px';
        successDiv.style.marginBottom = '20px';
        successDiv.style.borderRadius = '5px';
        successDiv.style.border = '2px solid red';
    });
}

btn.addEventListener('click', CadastroUsuario);
