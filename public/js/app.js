var app = {};

app.handleStar = function () {
	var star = $(".star-clickable"),
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

app.renderWeightChart = function () {
    var weightChart = document.getElementById("weightChart");
    if (!weightChart) {
        return;
    }

	var ctx = weightChart.getContext('2d');

	var config = {
            type: 'line',
            data: {
                labels: window.weightChartData,
                datasets: [{
                    label: "Weight entries",
                    data: window.weightChartData,
                    fill: true
                }]
            },
            options: {
                responsive: true,
                tooltips: {
                    mode: 'index',
                    intersect: false,
                },
                hover: {
                    mode: 'nearest',
                    intersect: true
                },
                scales: {
                    xAxes: [{
                        display: true,
                    }],
                    yAxes: [{
                        display: true,
                        scaleLabel: {
                            display: true,
                            labelString: 'Weight'
                        }
                    }]
                }
            }
        };

	var myChart = new Chart(ctx, config);
};

app.handleItemsFilterSelection = function () {
    $('.select-category').change(app.itemsCategorySelected);
    $('.select-brand').change(app.itemsBrandSelected);
    $('.select-color').change(app.itemsColorSelected);
};

app.itemsCategorySelected = function () {
    window.location.href = '/items?category=' + $('.select-category').val();
};

app.itemsBrandSelected = function () {
    window.location.href = '/items?brand=' + $('.select-brand').val();
};

app.itemsColorSelected = function () {
    window.location.href = '/items?color=' + $('.select-color').val();
};

// Page load
$(function () {
	app.handleStar();
	app.renderWeightChart();
    app.handleItemsFilterSelection();
});
