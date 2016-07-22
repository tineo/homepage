(function(){
  
// use strict mode
"use strict";

// Variables

var userWrapHeight;

//*****************************//


// Document Ready

$(document).ready(function() {
  
  // scrollspy init
  
  $('.scrollspy').scrollSpy({offsetTop: 0});
  
  //*****************************//
  
  
  // Main Revolution Slider Settings

  if(jQuery('.tp-banner').length > 0) {
    jQuery('.tp-banner').show().revolution({
      dottedOverlay:"none",
      delay:16000,
      startwidth:1170,
      startheight:700,
      hideThumbs:200,
      thumbWidth:100,
      thumbHeight:50,
      thumbAmount:5,
      navigationType:"none",
      navigationArrows:"solo",
      navigationStyle:"round",
      touchenabled:"on",
      onHoverStop:"off",
      swipe_velocity: 0.7,
      swipe_min_touches: 1,
      swipe_max_touches: 1,
      drag_block_vertical: false,						
      parallax:"mouse",
      parallaxBgFreeze:"on",
      parallaxLevels:[7,4,3,2,5,4,3,2,1,0],
      keyboardNavigation:"off",
      navigationHAlign:"center",
      navigationVAlign:"bottom",
      navigationHOffset:0,
      navigationVOffset:20,
      soloArrowLeftHalign:"left",
      soloArrowLeftValign:"center",
      soloArrowLeftHOffset:20,
      soloArrowLeftVOffset:0,
      soloArrowRightHalign:"right",
      soloArrowRightValign:"center",
      soloArrowRightHOffset:20,
      soloArrowRightVOffset:0,
      shadow:0,
      fullWidth:"on",
      fullScreen:"on",
      spinner:"spinner3",
      stopLoop:"off",
      stopAfterLoops:-1,
      stopAtSlide:-1,
      shuffle:"off",
      autoHeight:"off",
      forceFullWidth:"off",
      hideThumbsOnMobile:"off",
      hideNavDelayOnMobile:1500,						
      hideBulletsOnMobile:"off",
      hideArrowsOnMobile:"off",
      hideThumbsUnderResolution:0,
      hideSliderAtLimit:0,
      hideCaptionAtLimit:0,
      hideAllCaptionAtLilmit:0,
      startWithSlide:0,
      videoJsPath:"rs-plugin/videojs/",
      fullScreenOffsetContainer: ""	
    });
  }
  
  //*****************************//
  
  
  // Modal init
  
  //$('.modal-trigger').leanModal();
  
  //*****************************//
  
  
  // Mobile menu sub elements click hide all
  
  $('.mobile-menu ul > li > a').on('click',function(){
    $('.mobile-menu > a').trigger('click');
  });
  
  //*****************************//
  
  
  // Service Carousel
  
  $("#Services_Carousel").owlCarousel({
    navigation : false,
    pagination : true,
    slideSpeed : 1000,
    paginationSpeed : 2000,
    items : 3,
    itemsDesktop : [1200,3],
    itemsDesktopSmall : [992,3],
    itemsTablet : [768,2],
    itemsMobile : [400,1]
  });
  
  //*****************************//
  
  // Portfolio Carousel
  
  $("#Portfolio_Carousel").owlCarousel({
    navigation : false,
    pagination : true,
    slideSpeed : 1000,
    paginationSpeed : 2000,
    items : 3,
    itemsDesktop : [1200,3],
    itemsDesktopSmall : [992,3],
    itemsTablet : [768,2],
    itemsMobile : [400,1]
  });
  
  //*****************************//
  

  // References Carousel
  
  $("#References_Carousel").owlCarousel({
    navigation : false,
    pagination : true,
    slideSpeed : 1000,
    paginationSpeed : 2000,
    items : 5,
    itemsDesktop : [1200,5],
    itemsDesktopSmall : [992,4],
    itemsTablet : [768,3],
    itemsMobile : [400,2]
  });
  
  //*****************************//
  
  
  // Full Page Video
  
  if($('.video-container').length > 0) { fullVideo(); }
  
  //*****************************//
  
});

//*****************************//


// Window Load Function

$(window).load(function(){
  
  userContentScroll();
  
  // Loader
  
  setInterval(function(){
    $('.page-loading').fadeOut();
  }, 1200);
  
  //*****************************//
  
  $("html,body").scrollTop(0);
  
  // Full Page Image init
  
  fullPageImage();
  
  //*****************************//
});

//*****************************//


// Window Resize Function

$(window).resize(function(){
  userContentScroll();
  fullPageImage();
  if($('.video-container').length > 0) { fullVideo(); }
});

//*****************************//


// User Content Auto Scroll Function

function userContentScroll(){
  if($(window).width() > 992) {
    userWrapHeight = $(window).height()-$('.right-column header').height();
    $(".user-wrapper").css('max-height',userWrapHeight).mCustomScrollbar({
      autoHideScrollbar:true,
      theme:"minimal-dark"
    });
  }
}

//*****************************//


// Video //

var myVideo = document.getElementById("myVideo"); 

function fullVideo(){
  $('.video-container').width($('.left-column').width()).height($(window).height()-65);
}

function playPause(){ 
  if (myVideo.paused) {
    setTimeout(function(){
 		$(".video-container .play").fadeIn(500, function(){
 	      	$(".video-container .play").fadeOut(500);
 	   	});
 	  }, 100);
    myVideo.play();
  }
  else {
    setTimeout(function(){
 		$(".video-container .pause").fadeIn(500, function(){
 	      	$(".video-container .pause").fadeOut(500);
 	   	});
 	  }, 100);
    myVideo.pause(); 
  }
}

function audioControl(){
  if (myVideo.muted) {
    $('.video-container .mute i').text("volume_up");
    myVideo.muted = false;
  } else {
    $('.video-container .mute i').text("volume_off");
    myVideo.muted = true;
  }
}

//*****************************//

// Full Page Image

function fullPageImage(){
  if($(window).width()>992) {
    $('.full-image-container').height($(window).height()-65);
  } else {
    $('.full-image-container').height($(window).height());
  }
}

//*****************************//

})();