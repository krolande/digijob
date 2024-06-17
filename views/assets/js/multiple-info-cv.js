//ajouter plusieurs informations en multiple select : formation, compétence, expériences pro, certification
$(document).ready(function() {
    //nouvel ajout
    $(".add-elements").click(function() {
        var markup =
        "<div class='bg-light shadow-sm p-3 mt-4' style='border-radius: 15px;'><form class='row'><div class='form-group mb-3 col-md-6'><label class='form-label'>Titre de la formation</label><input type='text' class='form-control' value=''></div><div class='form-group mb-3 col-md-6'><label class='form-label'>Nom de l'Etablissement ou école</label><input type='text' class='form-control' value=''></div> <div class='form-group mb-3 col-md-4 select-border'> <label class='form-label'>Date de début</label> <div class='col-md-8 d-flex'> <select class='form-select basic-select' data-select2-id='1' tabindex='-1' aria-hidden='true'> <option value='' data-select2-id='3' aria-placeholder='Mois'>Février</option> <option value='janvier'>Janvier</option> <option value='fevrier'>Février</option> <option value='mars'>Mars</option> <option value='avril'>April</option> <option value='mai'>Mai</option> <option value='juin'>Juin</option> <option value='juillet'>Juillet</option> <option value='aout'>Aout</option> <option value='septembre'>Septembre</option> <option value='octobre'>Octobre</option> <option value='nevembre'>Novembre</option> <option value='decembre'>Décembre</option> </select> <select class='form-select mx-2 basic-select' data-select2-id='1' tabindex='-1' aria-hidden='true'> <option value='value 01' selected='selected' data-select2-id='3'>2020</option> <option value='value 02'>2008</option> <option value='value 03'>2009</option> <option value='value 04'>2010</option> <option value='value 05'>2011</option> <option value='value 06'>2012</option> <option value='value 07'>2013</option> <option value='value 08'>2014</option> <option value='value 09'>2015</option> <option value='value 10'>2016</option> <option value='value 11'>2017</option> <option value='value 12'>2018</option> <option value='value 13'>2019</option> <option value='value 14'>2020</option> <option value='value 15'>2021</option> <option value='value 16'>2022</option> <option value='value 17'>2023</option> <option value='value 18'>2024</option> <option value='value 19'>2025</option> <option value='value 20'>2026</option> <option value='value 21'>2027</option> <option value='value 22'>2028</option> <option value='value 23'>2029</option> <option value='value 14'>2030</option> </select> </div> </div> <div class='form-group mb-3 col-md-4 select-border'> <label class='form-label'>Date de fin</label> <div class='col-md-8 d-flex'> <select class='form-select basic-select' data-select2-id='1' tabindex='-1' aria-hidden='true'> <option value='' data-select2-id='3' aria-placeholder='Mois'>Février</option> <option value='janvier'>Janvier</option> <option value='fevrier'>Février</option> <option value='mars'>Mars</option> <option value='avril'>April</option> <option value='mai'>Mai</option> <option value='juin'>Juin</option> <option value='juillet'>Juillet</option> <option value='aout'>Aout</option> <option value='septembre'>Septembre</option> <option value='octobre'>Octobre</option> <option value='nevembre'>Novembre</option> <option value='decembre'>Décembre</option> </select> <select class='form-select mx-2 basic-select' data-select2-id='1' tabindex='-1' aria-hidden='true'> <option value='value 01' selected='selected' data-select2-id='3'>2020</option> <option value='value 02'>2008</option> <option value='value 03'>2009</option> <option value='value 04'>2010</option> <option value='value 05'>2011</option> <option value='value 06'>2012</option> <option value='value 07'>2013</option> <option value='value 08'>2014</option> <option value='value 09'>2015</option> <option value='value 10'>2016</option> <option value='value 11'>2017</option> <option value='value 12'>2018</option> <option value='value 13'>2019</option> <option value='value 14'>2020</option> <option value='value 15'>2021</option> <option value='value 16'>2022</option> <option value='value 17'>2023</option> <option value='value 18'>2024</option> <option value='value 19'>2025</option> <option value='value 20'>2026</option> <option value='value 21'>2027</option> <option value='value 22'>2028</option> <option value='value 23'>2029</option> <option value='value 14'>2030</option> </select> </div> </div> <div class='form-group mb-3 col-md-4'> <input class='form-check-input' type='checkbox' value='' id='defaultCheck1'> <label class='form-check-label' for='defaultCheck1'> à aujourd'hui </label> </div> <div class='form-group mb-3 col-md-12'> <label class='form-label'>Description</label> <textarea class='form-control' rows='4' placeholder='Saisissez une description détaillée de la formation'></textarea> </div> <div class='form-group col-md-12 mb-0'> <a class='btn btn-md btn-primary' href='#'>Ajouter</a> </div> </form> </div>"
        $("#elements").append(markup);
    });

    //supprimer l'élément ajouté après envoie du formulaire
    $(".delete-elements").click(function(){
    $(".info").remove();
    });
    //annuler le nouveau formulaire ajouté
    $(".delete-form").click(function(){
    $(".form").remove();
    });
});


// Cacher et affichrer les formulaires d'ajout
$(document).ready(function() {
    $('#cacher').hide();
    $("#afficher").click(function(){
    $(".afficher").show();
  });

  $("#valider-cacher").click(function(){
    $(".afficher").hide();
  });
});