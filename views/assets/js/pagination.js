/*--------Pagination job list--------*/
function getPageList(totalPages, page, maxLength){
    function range(start, end){
        return Array.from(Array(end - start + 1), (_, i) => i + start);
    }

    var sideWidth = maxLength < 9 ? 1 : 2;
    var leftWidth = (maxLength - sideWidth * 2 - 3) >> 1;
    var rightWidth = (maxLength - sideWidth * 2 - 3) >> 1;
    if(totalPages <= maxLength){
        return range(1, totalPages);
    }

    if(page <= maxLength - sideWidth - 1 - rightWidth){
        return range(1, maxLength - sideWidth - 1).concat(0, range(totalPages - sideWidth + 1, totalPages));
    }

    if(page >= totalPages - sideWidth - 1 - rightWidth){
        return range(1, sideWidth).concat(0, range(totalPages - sideWidth - 1 - rightWidth - leftWidth, totalPages));
    }

    return range(1, sideWidth).concat(0, range(page - leftWidth, page + rightWidth), 0, range(totalPages - sideWidth + 1, totalPages));
}
$(function(){
    var numberOfItems = $(".job-list-area .single-job-list-box").length;
    var limitPerPage = 6;
    var totalPages = Math.ceil(numberOfItems/limitPerPage);
    var paginationSize = 7;
    var currentPage;

    function showPage(whichPage){
        if(whichPage < 1 || whichPage > totalPages){
            return false;
        }
        currentPage = whichPage;

        $(".job-list-area .single-job-list-box").hide().slice((currentPage - 1) * limitPerPage, currentPage * limitPerPage).show();
        $(".pagination-area li").slice(1, -1).remove();

        getPageList(totalPages, currentPage, paginationSize).forEach(item =>{
            $("<li>").addClass("page-numbers").addClass(item ? "currentPage": "dots").toggleClass("current", item === currentPage)
            .append($("<a>").addClass("link").attr({href: "javascript:void(0)"}).text(item || "..."))
            .insertBefore(".next");
        });
        
        $(".prev").toggleClass("disable", currentPage === 1);
        $(".next").toggleClass("disable", currentPage === totalPages);
        return true
    }

    $(".pagination-area").append($("<li>").addClass("page-numbers").addClass("prev").append($("<a>")
    .addClass("link").attr({href: "javascript:void(0)"}).text("Prev")));
    $(".pagination-area").append($("<li>").addClass("page-numbers").addClass("next").append($("<a>")
    .addClass("link").attr({href: "javascript:void(0)"}).text("Next")));

	$(".job-list-area").show();
	showPage(1);

	$(document).on("click", ".pagination-area li.currentPage:not(.current)", function(){
		return showPage(+$(this).text());
	});

	$(".next").on("click", function(){
		return showPage(currentPage + 1);
	});
	$(".prev").on("click", function(){
		return showPage(currentPage - 1);
	});
});
/*--------End Pagination job list--------*/

/*--------Pagination digicodeuse list--------*/
function getPageDigi(totalPages, page, maxLength){
    function range(start, end){
        return Array.from(Array(end - start + 1), (_, i) => i + start);
    }

    var sideWidth = maxLength < 9 ? 1 : 2;
    var leftWidth = (maxLength - sideWidth * 2 - 3) >> 1;
    var rightWidth = (maxLength - sideWidth * 2 - 3) >> 1;
    if(totalPages <= maxLength){
        return range(1, totalPages);
    }

    if(page <= maxLength - sideWidth - 1 - rightWidth){
        return range(1, maxLength - sideWidth - 1).concat(0, range(totalPages - sideWidth + 1, totalPages));
    }

    if(page >= totalPages - sideWidth - 1 - rightWidth){
        return range(1, sideWidth).concat(0, range(totalPages - sideWidth - 1 - rightWidth - leftWidth, totalPages));
    }

    return range(1, sideWidth).concat(0, range(page - leftWidth, page + rightWidth), 0, range(totalPages - sideWidth + 1, totalPages));
}
$(function(){
    var numberOfItems = $(".digicodeuse .digi").length;
    var limitPerPage = 6;
    var totalPages = Math.ceil(numberOfItems/limitPerPage);
    var paginationSize = 7;
    var currentPage;

    function showPage(whichPage){
        if(whichPage < 1 || whichPage > totalPages){
            return false;
        }
        currentPage = whichPage;

        $(".digicodeuse .digi").hide().slice((currentPage - 1) * limitPerPage, currentPage * limitPerPage).show();
        $(".pagination-area li").slice(1, -1).remove();

        getPageList(totalPages, currentPage, paginationSize).forEach(item =>{
            $("<li>").addClass("page-numbers").addClass(item ? "currentPage": "dots").toggleClass("current", item === currentPage)
            .append($("<a>").addClass("link").attr({href: "javascript:void(0)"}).text(item || "..."))
            .insertBefore(".next");
        });
        
        $(".prev").toggleClass("disable", currentPage === 1);
        $(".next").toggleClass("disable", currentPage === totalPages);
        return true
    }

    $(".pagination-area").append($("<li>").addClass("page-numbers").addClass("prev").append($("<a>")
    .addClass("link").attr({href: "javascript:void(0)"}).text("Prev")));
    $(".pagination-area").append($("<li>").addClass("page-numbers").addClass("next").append($("<a>")
    .addClass("link").attr({href: "javascript:void(0)"}).text("Next")));

	$(".job-list-area").show();
	showPage(1);

	$(document).on("click", ".pagination-area li.currentPage:not(.current)", function(){
		return showPage(+$(this).text());
	});

	$(".next").on("click", function(){
		return showPage(currentPage + 1);
	});
	$(".prev").on("click", function(){
		return showPage(currentPage - 1);
	});
});
/*--------End Pagination digicodeuse list--------*/

/*------Pagination entreprise & evenement------*/ 
function getPageListCe(totalPages, page, maxLength){
    function range(start, end){
        return Array.from(Array(end - start + 1), (_, i) => i + start);
    }

    var sideWidth = maxLength < 9 ? 1 : 2;
    var leftWidth = (maxLength - sideWidth * 2 - 3) >> 1;
    var rightWidth = (maxLength - sideWidth * 2 - 3) >> 1;
    if(totalPages <= maxLength){
        return range(1, totalPages);
    }

    if(page <= maxLength - sideWidth - 1 - rightWidth){
        return range(1, maxLength - sideWidth - 1).concat(0, range(totalPages - sideWidth + 1, totalPages));
    }

    if(page >= totalPages - sideWidth - 1 - rightWidth){
        return range(1, sideWidth).concat(0, range(totalPages - sideWidth - 1 - rightWidth - leftWidth, totalPages));
    }

    return range(1, sideWidth).concat(0, range(page - leftWidth, page + rightWidth), 0, range(totalPages - sideWidth + 1, totalPages));
}
$(function(){
    var numberOfItems = $(".entre-event-list-area .card").length;
    var limitPerPage = 6;
    var totalPages = Math.ceil(numberOfItems/limitPerPage);
    var paginationSize = 7;
    var currentPage;

    function showPageCe(whichPage){
        if(whichPage < 1 || whichPage > totalPages){
            return false;
        }
        currentPage = whichPage;

        $(".entre-event-list-area .card").hide().slice((currentPage - 1) * limitPerPage, currentPage * limitPerPage).show();
        $(".pagination-entre-event li").slice(1, -1).remove();

        getPageListCe(totalPages, currentPage, paginationSize).forEach(item =>{
            $("<li>").addClass("page-numbers").addClass(item ? "currentPage": "dots").toggleClass("current", item === currentPage)
            .append($("<a>").addClass("link").attr({href: "javascript:void(0)"}).text(item || "..."))
            .insertBefore(".next");
        });
        
        $(".prev").toggleClass("disable", currentPage === 1);
        $(".next").toggleClass("disable", currentPage === totalPages);
        return true
    }

    $(".pagination-entre-event").append($("<li>").addClass("page-numbers").addClass("prev").append($("<a>")
    .addClass("link").attr({href: "javascript:void(0)"}).text("Prev")));
    $(".pagination-entre-event").append($("<li>").addClass("page-numbers").addClass("next").append($("<a>")
    .addClass("link").attr({href: "javascript:void(0)"}).text("Next")));

	$(".entre-event-list-area").show();
	showPageCe(1);

	$(document).on("click", ".pagination-entre-event li.currentPage:not(.current)", function(){
		return showPageCe(+$(this).text());
	});

	$(".next").on("click", function(){
		return showPageCe(currentPage + 1);
	});
	$(".prev").on("click", function(){
		return showPageCe(currentPage - 1);
	});
});
/*------End-Pagination entreprise & evenement------*/ 