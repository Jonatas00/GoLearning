$("#nova-publicacao").on("submit", criarPublicacao);
$(document).on("click", ".curtir-publicacao", curtirPublicacao);
$(document).on("click", ".descurtir-publicacao", descurtirPublicacao);

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
  const publicacaoId = elementoClicado.closest("div").data("publicacao-id");

  elementoClicado.prop("disabled", true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST"
  }).done(() => {
    const contadorDeCurtidas = elementoClicado.next("span");
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

    elementoClicado.addClass("descurtir-publicacao");
    elementoClicado.addClass("text-danger");
    elementoClicado.removeClass("curtir-publicacao");
  }).fail(() => {
    alert("erro ao curtir a publicacao!")
  }).always(() => {
    elementoClicado.prop("disabled", false);
  })
}

function descurtirPublicacao(e) {
  e.preventDefault();

  const elementoClicado = $(e.target);
  const publicacaoId = elementoClicado.closest("div").data("publicacao-id");

  elementoClicado.prop("disabled", true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/descurtir`,
    method: "POST"
  }).done(() => {
    const contadorDeCurtidas = elementoClicado.next("span");
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

    elementoClicado.removeClass("descurtir-publicacao");
    elementoClicado.removeClass("text-danger");
    elementoClicado.addClass("curtir-publicacao");
  }).fail(() => {
    alert("erro ao descurtir a publicacao!")
  }).always(() => {
    elementoClicado.prop("disabled", false);
  })
}