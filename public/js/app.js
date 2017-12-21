var app = {};

app.handleStar = function () {
	var star = $(".star"),
		input = $("#star");
	star.click(function (evt) {
		if (star.hasClass('fa-star-o')) {
			star.removeClass('fa-star-o');
			star.addClass('fa-star');
			input.val('true');
		} else {
			star.removeClass('fa-star');
			star.addClass('fa-star-o');
			input.val('false');
		}
	});
};

$(function () {
	app.handleStar();
});
