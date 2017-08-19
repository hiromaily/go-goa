(function(){
    var rg = function(){};
    //public
    rg.debug = 1;

//------------------------------------------------------------------------------
//public
//------------------------------------------------------------------------------
    //fetch API
    rg.callAPI = function(obj){
        let key = sessionStorage.getItem('jwt');
        //fetch
        fetch(obj.url, {
            headers: {
                'Authorization': 'Bearer '+key,
                "Content-Type": "application/json"
            }
        }).then(function(response) {
            return response.json();
        }).then(function(json) {
            if(json.status && json.status != 200){
                sessionStorage.removeItem('jwt');
                sessionStorage.removeItem('id');
                location.href = '/resume.html';
                return;
            }
            //Success
            riot.mount(obj.element, {
                items : json
            } );
        });
    };

    window.rg = rg;
})();