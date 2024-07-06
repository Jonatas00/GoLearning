$(`#nova-publicacao`).on(`submit`, criarPublicacao)

function criarPublicacao(e) {
  e.preventDefault();

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $(`#titulo`).val(),
      conteudo: $(`#conteudo`).val(),
    }
  }).done(() => {
    window.location = "/home";
  }).fail(() => {
    alert("Erro ao criar publicação!");
  })
}