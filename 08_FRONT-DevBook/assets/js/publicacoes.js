$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

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

function curtirPublicacao(e) {
  e.preventDefault();

  const elementoClicado = $(e.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

  elementoClicado.prop("disabled", true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST"
  }).done(() => {
    const contadorDeCurtidas = elementoClicado.next("span");
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas + 1);
  }).fail(() => {
    alert("erro ao curtir a publicacao!")
  }).always(() => {
    elementoClicado.prop("disabled", false);
  })
}
