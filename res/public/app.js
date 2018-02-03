window.onload = function () {
    Split(['#pillarid', '#contentid'], {
        sizes: [25, 75],
        elementStyle: function (dimension, size, gutterSize) {
            return {
                'flex-basis': 'calc(' + size + '% - ' + gutterSize + 'px)'
            }
        },
        gutterStyle: function (dimension, gutterSize) {
            return {
                'flex-basis': gutterSize + 'px'
            }
        }
    });
}