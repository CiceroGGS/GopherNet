$(`#register-form`).on(`submit`, createUser);

function createUser(event) {
    event.preventDefault();

    if ($('#password').val() != $('#confirmPassword').val()) {
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $(`#name`).val(),
            nick: $(`#nick`).val(),
            email: $(`#email`).val(),
            password: $(`#password`).val(),
        }
    }).done(function () {
        alert("Usuario cadastrado com sucesso")
    }).fail(function (erro) {
        console.log(erro);
        alert("Erro ao cadastro usuario!")
    });
}

