$(`#login-form`).on(`submit`, login)

function login(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $(`#email`).val(),
            password: $(`#password`).val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function (erro) {
        console.log(erro);
        alert("Email ou senha invalidos")
    });
}
