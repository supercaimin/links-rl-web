function switchSliderView(selector) {
    var that = $(selector);

    that.find('.switch-preview .swiper-slide:first-child').addClass('active-nav');
    var viewSwiper = new Swiper(that.find('.switch-slider-view .swiper-container'), {
        on: {
            slideChangeTransitionStart: function () {
                updateNavPosition();
            },
        }
    });
    that.find('.switch-slider-view .arrow-left,.switch-preview .arrow-left').on('click', function (e) {
        e.preventDefault();
        if (viewSwiper.activeIndex == 0) {
            viewSwiper.slideTo(viewSwiper.slides.length - 1, 1000);
            return;
        }
        viewSwiper.slidePrev();
    });

    that.find('.switch-slider-view .arrow-right,.switch-preview .arrow-right').on('click', function (e) {
        e.preventDefault();
        if (viewSwiper.activeIndex == viewSwiper.slides.length - 1) {
            viewSwiper.slideTo(0, 1000);
            return
        }
        viewSwiper.slideNext()
    });

    var previewSwiper = new Swiper(that.find('.switch-preview .swiper-container'), {
        //visibilityFullFit: true,
        slidesPerView: 'auto',
        allowTouchMove: false,
        on: {
            tap: function () {

                viewSwiper.slideTo(previewSwiper.clickedIndex);
            }
        }
    });

    that.find('.switch-preview .swiper-slide').on('click', function () {
        var th = $(this);
        if (th.hasClass('last-clicked')) return;
        that.find('.switch-preview .last-clicked').removeClass('last-clicked');
        var ph = th.parent().offset().left + th.parent().width() / 2;
        if ((ph - th.offset().left) > 0) {
            previewSwiper.slidePrev();
        } else {
            previewSwiper.slideNext();
        }
        $(this).addClass('last-clicked');

    });
    var updateNavPosition = function () {
        that.find('.switch-preview .active-nav').removeClass('active-nav');
        var activeNav = that.find('.switch-preview .swiper-slide').eq(viewSwiper.activeIndex).addClass('active-nav');
        if (!activeNav.hasClass('swiper-slide-visible')) {
            if (activeNav.index() > previewSwiper.activeIndex) {
                var thumbsPerNav = Math.floor(previewSwiper.width / activeNav.width()) - 1;
                previewSwiper.slideTo(activeNav.index() - thumbsPerNav);
            } else {
                previewSwiper.slideTo(activeNav.index());
            }
        }
    };
}

$.fn.extend({
    animateCss: function (animationName, callback) {
        var animationEnd = (function (el) {
            var animations = {
                animation: 'animationend',
                OAnimation: 'oAnimationEnd',
                MozAnimation: 'mozAnimationEnd',
                WebkitAnimation: 'webkitAnimationEnd',
            };

            for (var t in animations) {
                if (el.style[t] !== undefined) {
                    return animations[t];
                }
            }
        })(document.createElement('div'));

        this.addClass('animated ' + animationName).one(animationEnd, function () {
            $(this).removeClass('animated ' + animationName);

            if (typeof callback === 'function') callback($(this));
        });

        return this;
    },
});

$(function () {
    if ($('.switch-slider').length) switchSliderView('.switch-slider');

    if ($('.footer-hot-cards').length) {
        var mySwiper = new Swiper('.footer-hot-cards .swiper-container', {
            loop: true,
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev'
            }
        })
    }

    $('body').on('click', '[data-target=link]', function () {
        var link = $(this).data('href');
        if (link) window.location.href = link;
    }).on('click', '.live-in-hk', function () {
        var el = $(this);
        if(el.next().hasClass('collapsed'))
        {
            el.next().animateCss('fadeOut',function (el) {
                el.addClass('collapse').removeClass('collapsed');
            });
        }
        else{
            el.next().addClass('collapsed').removeClass('collapse').animateCss('fadeIn');
        }

    });

});