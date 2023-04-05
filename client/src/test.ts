/// <reference path="../lib/typings/jquery/jquery.d.ts"/>

window.onload = () => {

    let tiles = [2, 3, 1, 4, 5, 6, 7, 8, 0];

    const $target = undefined;

    var renderTiles = function ($newTarget?) {
        $target = $newTarget || $target;

        var $ul = $("<ul>", {
            "class": "n-puzzle"
        });

        $(tiles).each(function (index) {
            var correct = index + 1 == this;
            var cssClass = this == 0 ? "empty" : (correct ? "correct" : "incorrect");

            var $li = $("<li>", {
                "class": cssClass,
                "data-tile": this,
            });
            $li.text(this);
            $ul.append($li);
        })


        $target.html($ul);
    };


    renderTiles($('.eight-puzzle'));
};