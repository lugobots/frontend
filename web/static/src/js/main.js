// Internal size
var screenWidth = window.innerWidth;
var screenHeight = window.innerHeight;

// Elements
var page = $("#lugobot-page");
var header = $('#lugobot-header');
var panel = $('#lugobot-admin-panel');
var login = $('#login-area');
var stadium = $('#lugobot-stadium');
var field = $('#field');
var eventView = $("#event-view");


$(window).ready(function(){
  loading('close');
  size_things();
  eventView.hide();

  $('.close-modal').on('click', function(){
              
      $("#event-view .modal").addClass("zoom-Out");
      page.removeClass('active-modal');

      setTimeout(function(){
        $("#event-view, #event-goal").removeClass('active-modal');
      },1000);

      setTimeout(function(){
        $('#event-view .modal').removeClass('active-modal');
        
        $("#event-view .modal").removeClass(function (index, css) {
          return (css.match (/\bzoom\S+/g) || []).join(' ');
        });

        $("#event-view, #score-info .score-team").removeClass(function (index, css) {
          return (css.match (/\bgoal\S+/g) || []).join(' ');
        });
  
        $("#event-view, #score-info .score-team").removeClass('goal');
        
          
        $(".score-team").removeClass("score-winner");
        eventView.hide();
      },2000);
  
    });

});

$(window).resize(function(){
  screenWidth = window.innerWidth;
  screenHeight = window.innerHeight;

  size_things();
  admin_scale();
});

  // Load functions
  function loading(action){
    if( action == 'close'){

      setTimeout(function(){
        $("#loading").fadeOut('slow',0);
        
        setTimeout(function(){
  
          // Header
          if(screenWidth >= 1200){
            header.css('transform','translateY(0%)');
            if(panel.hasClass('debug-mode')){
              panel.css('transform','translateY(0%)');
            }else{
              login.css('transform','translateY(0%)');
            }
          } 
          
          else if(screenWidth < 1200 && screenWidth >= 992 ){
            header.css('transform','translateY(0%) scale(.95)');
            if(panel.hasClass('debug-mode')){
              panel.css('transform','translateY(0%) scale(.85)');
            }else{
              login.css('transform','translateY(0%)');
            }
          } 

          else if(screenWidth < 992 && screenWidth >= 768){
            header.css('transform','translateY(0%) scale(.8)');

            if(panel.hasClass('debug-mode')){
              panel.css('transform','translateY(0%) scale(1)');
            }else{
              login.css('transform','translateY(0%)');
            }
            
          }

          else if(screenWidth < 768 && screenWidth >= 576){
            header.css('transform','translateY(0%) scale(.7)');
            
            if(panel.hasClass('debug-mode')){
              panel.css('transform','translateY(0%) scale(1)');
            }else{
              login.css('transform','translateY(0%)');
            }
          }
          
          else if(screenWidth < 576 && screenWidth >= 470){
            header.css('transform','translateY(0%) scale(.65)');
            
            if(panel.hasClass('debug-mode')){
              panel.css('transform','translateY(0%) scale(1)');
            }else{
              login.css('transform','translateY(0%)');
            }
          }
          
          else if(screenWidth < 470){
            header.css('transform','translateY(0%) scale(.65)');
            
            if(panel.hasClass('debug-mode')){
              panel.css('transform','translateY(0%) scale(1)');
            }else{
              login.css('transform','translateY(0%)');
            }
          }
  
          $("#loading").hide();
  
        },200);
  
      },2000);
    }

    else if(action == 'open'){

      if(screenWidth >= 1200){
        header.css('transform','translateY(-100%)');
        
        if(panel.hasClass('debug-mode')){
          panel.css('transform','translateY(100%)');
        }else{
          login.css('transform','translateY(120%)');
        }
      }
      
      else if(screenWidth < 1200 && screenWidth >= 992 ){
        header.css('transform','translateY(-100%) scale(.95)');

        if(panel.hasClass('debug-mode')){
          panel.css('transform','translateY(100%) scale(.85)');
        }else{
          login.css('transform','translateY(120%)');
        }
      }
      
      else if(screenWidth < 992 && screenWidth >= 768){
        header.css('transform','translateY(-100%) scale(.8)');

        if(panel.hasClass('debug-mode')){
          panel.css('transform','translateY(100%) scale(1)');
        }else{
          login.css('transform','translateY(120%)');
        }
      }
      
      else if(screenWidth < 768 && screenWidth >= 470){
        header.css('transform','translateY(-100%) scale(.7)');
        if(panel.hasClass('debug-mode')){
          panel.css('transform','translateY(100%) scale(1)');
        }else{
          login.css('transform','translateY(120%)');
        }

      }
      
      else if(screenWidth < 470){
        header.css('transform','translateY(-100%) scale(.65)');
                if(panel.hasClass('debug-mode')){
          panel.css('transform','translateY(100%) scale(1)');
        }else{
          login.css('transform','translateY(120%)');
        }
      }
      
      setTimeout(function(){
        $("#loading").fadeIn();
        loading('close');

      },2300);
    
    }
  }

  // Login functions
  
  function admin_panel(action){
    stadium.addClass('admin-mode');

    if( action == 'login'){
      login.css('transform','translateY(120%)');
      $('#lugobot-view, #lugobot-header, #lugobot-stadium, #event-view, #lugobot-admin-panel').addClass('debug-mode');
      size_things();

      setTimeout(function(){ 
        login.removeClass('user-mode');
        panel.css('transform','translateY(0%)');
      },2300);
    }
    else if ( action == 'logoff'){
      login.addClass('user-mode');
      panel.css('transform','translateY(100%)');

      setTimeout(function(){
        login.css('transform','translateY(0%)');
      },100);

      setTimeout(function(){
        $('#lugobot-view, #lugobot-header, #lugobot-stadium, #event-view, #lugobot-admin-panel').removeClass('debug-mode');
        size_things();
        
      },2300);
    }
  }

  // Resize functions
  function size_things(){
    
    // Proportion field calc
    var widthNew = field.css('width').replace(/px/,'') * 1;
    var heightNew = widthNew / 2;

    field.css('height', heightNew);

    // Position Field
    var scoreboardHeight = header.css('height').replace(/px/,'') * 1;
    var panelHeight = panel.css('height').replace(/px/,'') * 1;
    var loginHeight = login.css('height').replace(/px/,'') * 1;
    var positionField ;

    if(panel.hasClass('debug-mode')){
      positionField = (screenHeight - (heightNew + 40) - scoreboardHeight - panelHeight ) / 2;
    }else{
      positionField = (screenHeight - (heightNew + 40) - scoreboardHeight - loginHeight) / 2;
    }

    stadium.css('margin', positionField + 'px auto'); 

    
    // Position Debug bar 
    var sizeBar = panel.css('width').replace(/px/,'') * 1; 
    var positionBar = (screenWidth - sizeBar) / 2 ;
    
    if(screenWidth > 1250){
      $('#lugobot-admin-panel, #login-area').css('left', positionBar + 'px'); 
    }else{
      $('#lugobot-admin-panel, #login-area').css('left','0px'); 
    }

  }

  function admin_scale(){
      // Admin panel
       if(screenWidth < 1200 && screenWidth >= 992 ){
        panel.css('transform','translateY(0%) scale(.85)');
      } else if(screenWidth < 992){
        panel.css('transform','translateY(0%) scale(1)');
      }
  }

// Navigation between admin tabs
  function open_tab_panel(evt, tab_name) {
    var i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("tab-content");

    for (i = 0; i < tabcontent.length; i++) {
      tabcontent[i].style.display = "none";
      tabcontent[i].className = tabcontent[i].className.replace(" active-tab-content", "");
      
    }
    
    tablinks = document.getElementsByClassName("tab-link");
    for (i = 0; i < tablinks.length; i++) {
      tablinks[i].className = tablinks[i].className.replace(" active-tab-link", "");
    }

    $('#' + tab_name).css('display','flex').addClass('active-tab-content');
    evt.currentTarget.className += " active-tab-link ";
  }

  // Modals Functions
  function modal_game_over(){
    
    setTimeout(function(){ $("#event-view #modal-game-over").addClass("zoom-In");  },100);

    page.toggleClass("active-modal");
    eventView.toggleClass("active-modal");
    $("#modal-game-over").toggleClass("active-modal");
  
  }

  function modal_winner(team){
    
    var winner = team;

    $('#modal-winner #score-'+ winner +'-team ').addClass('score-winner');
 
    setTimeout(function(){ $("#event-view #modal-winner").addClass("zoom-In");  },100);

    eventView.toggleClass("active-modal").css('display','flex');
    $("#modal-winner").toggleClass("active-modal");
    page.toggleClass("active-modal");

  }

  function event_goal(team){

    var home_bg = "rgb(" + getComputedStyle(document.documentElement).getPropertyValue('--team-home-color-primary') + ")";
    var away_bg = "rgb(" + getComputedStyle(document.documentElement).getPropertyValue('--team-away-color-primary') + ")";

    var autor = team;
    
    setTimeout(function(){eventView.addClass("zoom-In"); },1000);

    eventView.addClass("active-modal");
    header.addClass("active-modal");

    $("#event-view, #event-goal, #score-info .score-team").removeClass(function (index, css) { 
      return (css.match (/\bgoal\S+/g) || []).join(' ');
    });
    
    $("#event-view, #event-goal").addClass("active-modal goal goal-"+ autor);
    $("#score-" + autor +"-team").toggleClass("goal"); 
    
    setTimeout(function(){
      eventView.addClass("zoom-Out");
      header.removeClass('active-modal');
      $(".score-team").removeClass("goal");

      setTimeout(function(){
       eventView.removeClass(function (index, css) {
          return (css.match (/\bzoom\S+/g) || []).join(' ');
        });
        $("#event-view, #event-goal, #score-info .score-team").removeClass(function (index, css) {
          return (css.match (/\bgoal\S+/g) || []).join(' ');
        });

        $("#event-view, #event-goal, #score-info .score-team").removeClass('goal');

        $("#event-view, #event-goal").removeClass('active-modal');
        eventView.hide();

      },2000);

    },6000);

  }
